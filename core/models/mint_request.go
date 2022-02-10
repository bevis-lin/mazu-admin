package models

import (
	"github.com/onflow/flow-go-sdk"
)

type MintRequest struct {
	RequestId  uint64       `json:"requestId"`
	Creator    flow.Address `json:"creator"`
	TemplateId uint64       `json:"templateId"`
	//Price      string       `json:"price"`
	Completed bool     `json:"completed"`
	Template  Template `json:"template"`
}
