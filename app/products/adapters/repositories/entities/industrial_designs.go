package entities

import (
	"encoding/json"
	"rpcf/core/entities"
)

const (
	IndustrialDesignsTable = "industrial_designs"
)

type IndustrialDesigns struct {
	*entities.Base
	ID                 string `gorm:"primaryKey"`
	Title              string
	Availability       string
	FundingInstitution string `json:"funding_institution"`
	year               string
}

func (IndustrialDesigns) TableName() string {
	return IndustrialDesignsTable
}

func NewIndustrialDesigns(product map[string]string) (*IndustrialDesigns, error) {
	var p IndustrialDesigns

	jsonBody, err := json.Marshal(product)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBody, &p)

	if err != nil {
		return nil, err
	}

	return &p, nil
}
