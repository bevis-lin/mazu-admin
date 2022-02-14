package controllers

import (
	"fmt"
	"io/ioutil"
	"mazu/admin/core/models"
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

func getAllTemplate() []models.Template {
	script, err := ioutil.ReadFile("flow/get-all-template.cdc")
	if err != nil {
		panic("failed to load Candence script")
	}

	result := ExecuteScript(Node, []byte(script))

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

	result := ExecuteScript(Node, []byte(script))

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

	c.JSON(200, gin.H{"status": "ok", "txId": tx.ID().String()})

}
