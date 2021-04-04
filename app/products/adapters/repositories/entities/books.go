package entities

import (
	"encoding/json"
	"rpcf/core/entities"
)

const (
	BooksTableName = "books"
)

type Books struct {
	*entities.Base
	ID               string `gorm:"primaryKey"`
	Title            string `json:"title"`
	PublishedCountry string `json:"published_country"`
	Editorial        string
}

func NewBook(product map[string]string) (*Books, error) {
	var a Books
	jsonBody, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBody, &a)

	if err != nil {
		return nil, err
	}

	a.ID = a.GenerateID()
	return &a, nil
}
