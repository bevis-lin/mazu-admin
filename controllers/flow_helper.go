package controllers

import (
	"context"
	"errors"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"github.com/onflow/flow-go-sdk/crypto"
	"google.golang.org/grpc"
)

func ExecuteScript(node string, script []byte, args []cadence.Value) cadence.Value {
	ctx := context.Background()
	c, err := client.New(node, grpc.WithInsecure())
	if err != nil {
		panic("failed to connect to node")
	}

	result, err := c.ExecuteScriptAtLatestBlock(ctx, script, args)
	if err != nil {
		panic(err)
	}

	return result

}

func InitTransaction(node string, script []byte, adminAddress string) (flow.Transaction, error) {

	proposerAddress := flow.HexToAddress(adminAddress)
	proposerKeyIndex := 0

	payerAddress := flow.HexToAddress(adminAddress)
	authorizerAddress := flow.HexToAddress(adminAddress)

	flowClient, err := client.New(node, grpc.WithInsecure())
	if err != nil {
		//panic("failed to connect to node")
		return flow.Transaction{}, errors.New("failed to connect to node")
	}

	latestBlock, err := flowClient.GetLatestBlockHeader(context.Background(), true)
	if err != nil {
		//panic("failed to fetch latest block")
		return flow.Transaction{}, errors.New("failed to fetch latest block")
	}

	proposerAccount, err := flowClient.GetAccountAtLatestBlock(context.Background(), proposerAddress)
	if err != nil {
		//panic("failed to fetch proposer account")
		return flow.Transaction{}, errors.New("failed to fectch proposer account")
	}

	sequenceNumber := proposerAccount.Keys[proposerKeyIndex].SequenceNumber

	tx := flow.NewTransaction().
		SetScript(script).
		SetGasLimit(1000).
		SetReferenceBlockID(latestBlock.ID).
		SetProposalKey(proposerAddress, proposerKeyIndex, sequenceNumber).
		SetPayer(payerAddress).
		AddAuthorizer(authorizerAddress)

	return *tx, nil
}

func SignTransaction(node string, tx *flow.Transaction, adminAddress string) (*flow.Transaction, error) {

	proposerAddress := flow.HexToAddress(adminAddress)
	proposerKeyIndex := 0

	flowClient, err := client.New(node, grpc.WithInsecure())
	if err != nil {
		//panic("failed to connect to node")
		return tx, errors.New("failed to connect to node")
	}

	proposerAccount, err := flowClient.GetAccountAtLatestBlock(context.Background(), proposerAddress)
	if err != nil {
		//panic("failed to fetch proposer account")
		return tx, errors.New("failed to fectch proposer account")
	}

	key1 := proposerAccount.Keys[proposerKeyIndex]

	privateKey, err := crypto.DecodePrivateKeyHex(key1.SigAlgo, AdminPrivateKey)

	if err != nil {
		return tx, errors.New("failed to decode private key")
	}

	mySigner := crypto.NewInMemorySigner(privateKey, key1.HashAlgo)

	tx.SignEnvelope(proposerAddress, key1.Index, mySigner)

	return tx, nil

}

func SendTransaction(node string, tx *flow.Transaction) (*flow.Transaction, error) {

	flowClient, err := client.New(node, grpc.WithInsecure())
	if err != nil {
		//panic("failed to connect to node")
		return tx, errors.New("failed to connect to node")
	}

	flowClient.SendTransaction(context.Background(), *tx)

	return tx, nil

}
