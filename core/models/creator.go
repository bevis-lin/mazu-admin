package models

import (
	"github.com/onflow/flow-go-sdk"
)

type Creator struct {
	Address    flow.Address                `json:"address"`
	Email      string                      `json:"email"`
	Name       string                      `json:"name"`
	ProfileURL string                      `json:"profileUrl"`
	Data       map[interface{}]interface{} `json:"data"`
}
