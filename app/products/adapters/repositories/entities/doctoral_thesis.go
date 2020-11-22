package entities

import (
	"encoding/json"
	"rpcf/core/entities"
)

const (
	DoctoralThesisTable = "doctoral_thesis"
)

type DoctoralThesis struct {
	*entities.Base
	ID          string `gorm:"PRIMARY_KEY"`
	Type        string
	Orientation string
	Institution string
}

func NewDoctoralThesis(product map[string]string) (*DoctoralThesis, error) {
	var p DoctoralThesis

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
