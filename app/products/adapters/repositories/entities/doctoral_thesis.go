package entities

import (
	"encoding/json"
	"rpcf/core/entities"
)

const (
	DoctoralThesisTable = "doctoral_theses"
)

type DoctoralThesis struct {
	*entities.Base
	ID              string `gorm:"primaryKey"`
	Title           string
	Type            string
	Orientation     string
	Institution     string
	StudentName     string `json:"student_name"`
	AcademicProgram string `json:"academic_program"`
	Pages           string
	StartDate       string `json:"start_date"`
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
