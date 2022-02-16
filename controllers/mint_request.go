package controllers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mazu/admin/core/models"
	"net/smtp"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
)

var AdminAddress = os.Getenv("ADMIN_ADDRESS")
var AdminPrivateKey = os.Getenv("ADMIN_PRIVATE_KEY")
var Node = os.Getenv("ACCESS_NODE")
var EmailPassword = os.Getenv("EMAIL_PASSWORD")

func getAllTemplate() []models.Template {
	script, err := ioutil.ReadFile("flow/get-all-template.cdc")
	if err != nil {
		panic("failed to load Candence script")
	}

	result := ExecuteScript(Node, []byte(script), nil)

	s := result.(cadence.Dictionary)

	var templates []models.Template

	//fmt.Println(s)

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	for _, template := range s.Pairs {
		templateT := template.Value.(cadence.Struct)
		data, err := json.Marshal(templateT.Fields[6].ToGoValue().(map[interface{}]interface{}))
		if err != nil {
			panic("failed to convert data to json string")
		}

		templateVar := models.Template{
			TemplateId:  templateT.Fields[0].ToGoValue().(uint64),
			SiteId:      templateT.Fields[1].ToGoValue().(string),
			Creator:     templateT.Fields[2].ToGoValue().([flow.AddressLength]byte),
			Name:        templateT.Fields[3].ToGoValue().(string),
			Description: templateT.Fields[4].ToGoValue().(string),
			ImageUrl:    templateT.Fields[5].ToGoValue().(string),
			Data:        string(data),
			TotalSupply: templateT.Fields[7].ToGoValue().(uint64),
			TotalMinted: templateT.Fields[8].ToGoValue().(uint64)}

		templates = append(templates, templateVar)

	}

	return templates
}

func GetAllRequests(c *gin.Context) {

	templates := getAllTemplate()

	//fmt.Println(templates)

	script, err := ioutil.ReadFile("flow/get-all-mint-request.cdc")
	if err != nil {
		panic("failed to load Candence script...v3")
	}

	result := ExecuteScript(Node, []byte(script), nil)

	s := result.(cadence.Dictionary)

	var mintRequests []models.MintRequest

	for _, mintRequest := range s.Pairs {
		mintT := mintRequest.Value.(cadence.Struct)
		mintRequestStruct := models.MintRequest{
			RequestId:  mintT.Fields[0].ToGoValue().(uint64),
			Creator:    mintT.Fields[1].ToGoValue().([flow.AddressLength]byte),
			TemplateId: mintT.Fields[2].ToGoValue().(uint64),
			//Price:      mintT.Fields[3].String(),
			Completed: mintT.Fields[3].ToGoValue().(bool)}

		for _, template := range templates {
			if template.TemplateId == mintRequestStruct.TemplateId {
				mintRequestStruct.Template = template
				break
			}
		}

		mintRequests = append(mintRequests, mintRequestStruct)
	}

	c.JSON(200, mintRequests)
}

func ApproveMintRequest(c *gin.Context) {

	approveScript, err := ioutil.ReadFile("flow/approve-mint-request.cdc")
	if err != nil {
		panic("faild to load Candence script")
	}

	tx, err := InitTransaction(Node, approveScript, AdminAddress)

	requestIdValue := c.PostForm("requestId")

	rqID, err := strconv.ParseUint(requestIdValue, 10, 64)

	fmt.Println("requestId:", rqID)

	if err != nil {
		panic("failed to parse Uint")
	}

	requestId := cadence.NewUInt64(rqID)

	err = tx.AddArgument(requestId)

	if err != nil {
		panic("invalid argument")
	}

	SignTransaction(Node, &tx, AdminAddress)

	SendTransaction(Node, &tx)

	creator, err := getCreatorByMintRequestId(rqID)

	if err != nil {
		panic("can't get creator profile")
	}

	fmt.Println("creator:", creator)

	from := "bevis@digi96.com"
	fmt.Println("password:", EmailPassword)
	password := EmailPassword
	to := []string{creator.Email}
	smtpHost := "smtpout.secureserver.net"
	smtpPort := "587"

	srtMessage := "Mint request(id:" + requestIdValue + ") has been approved."
	msg := "From: " + from + "\n" + "To: " + to[0] + "\n" + "Subject: Sentimen mint result\n\n" + srtMessage

	message := []byte(msg)

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{"status": "ok", "txId": tx.ID().String()})

}

func getRequestById(requestId uint64) (*models.MintRequest, error) {
	script, err := ioutil.ReadFile("flow/get-mint-request-by-id.cdc")
	if err != nil {
		return nil, errors.New("failed to load Candence script...")
	}

	args := []cadence.Value{cadence.NewUInt64(requestId)}

	result := ExecuteScript(Node, []byte(script), args)

	mintT := result.(cadence.Optional).Value.(cadence.Struct)

	mintRequestStruct := models.MintRequest{
		RequestId:  mintT.Fields[0].ToGoValue().(uint64),
		Creator:    mintT.Fields[1].ToGoValue().([flow.AddressLength]byte),
		TemplateId: mintT.Fields[2].ToGoValue().(uint64),
		//Price:      mintT.Fields[3].String(),
		Completed: mintT.Fields[3].ToGoValue().(bool)}

	return &mintRequestStruct, nil

}

func getCreatorByMintRequestId(requestId uint64) (*models.Creator, error) {
	mintRequest, err := getRequestById(requestId)
	if err != nil {
		return nil, err
	}

	creator, err := getCreatorProfile(mintRequest.Creator)

	if err != nil {
		return nil, err
	}

	return creator, nil

}
