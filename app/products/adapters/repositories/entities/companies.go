package entities

import (
	"encoding/json"
	"rpcf/core/entities"
)

const (
	CompaniesTable = "companies"
)

type Companies struct {
	*entities.Base
	ID          string `gorm:"primaryKey" json:"id"`
	Type        string
	Title       string
	NIT         string
	RegistredAt string `json:"registred_at"`
}

func NewCompanies(product map[string]string) (*Companies, error) {
	var p Companies
	jsonBody, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBody, &p)
	if err != nil {
		return nil, err
	}
	p.ID = p.GenerateID()
	return &p, nil
}
