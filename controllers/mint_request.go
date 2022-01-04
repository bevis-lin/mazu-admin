package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
)

var AdminAddress = os.Getenv("ADMIN_ADDRESS")
var AdminPrivateKey = os.Getenv("ADMIN_PRIVATE_KEY")
var Node = os.Getenv("ACCESS_NODE")

func GetAllRequests(c *gin.Context) {
	script, err := ioutil.ReadFile("flow/get-all-mint-request.cdc")
	if err != nil {
		panic("failed to load Candence script")
	}

	result := ExecuteScript(Node, []byte(script))

	s := result.(cadence.Dictionary)

	var mintRequests []MintRequest

	for _, mintRequest := range s.Pairs {
		mintT := mintRequest.Value.(cadence.Struct)
		mintRequestStruct := MintRequest{
			RequestId:  mintT.Fields[0].ToGoValue().(uint64),
			Creator:    mintT.Fields[1].ToGoValue().([flow.AddressLength]byte),
			TemplateId: mintT.Fields[2].ToGoValue().(uint64),
			Price:      mintT.Fields[3].String(),
			Completed:  mintT.Fields[4].ToGoValue().(bool)}

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
