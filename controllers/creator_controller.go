package controllers

import (
	"errors"
	"io/ioutil"
	"mazu/admin/core/models"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
)

func getCreatorProfile(address flow.Address) (*models.Creator, error) {
	script, err := ioutil.ReadFile("flow/get-creator-profile.cdc")
	if err != nil {
		return nil, errors.New("failed to load Candence script...")
	}

	args := []cadence.Value{cadence.NewAddress(address)}

	result := ExecuteScript(Node, []byte(script), args)

	creatorT := result.(cadence.Optional).Value.(cadence.Struct)

	creatorStruct := models.Creator{
		Address:    creatorT.Fields[0].ToGoValue().([flow.AddressLength]byte),
		Name:       creatorT.Fields[1].ToGoValue().(string),
		Email:      creatorT.Fields[2].ToGoValue().(string),
		ProfileURL: creatorT.Fields[3].ToGoValue().(string),
		Data:       creatorT.Fields[4].ToGoValue().(map[interface{}]interface{}),
	}

	return &creatorStruct, nil

}
