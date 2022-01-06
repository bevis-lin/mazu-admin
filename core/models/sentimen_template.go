package models

import (
	"github.com/onflow/flow-go-sdk"
)

type Template struct {
	TemplateId  uint64       `json:"templateId"`
	SiteId      string       `json:"siteId"`
	Creator     flow.Address `json:"creator"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	ImageUrl    string       `json:"imageUrl"`
	Data        string       `json:"data"`
	TotalSupply uint64       `json:"totalSupply"`
	TotalMinted uint64       `json:"totalMinted"`
}
