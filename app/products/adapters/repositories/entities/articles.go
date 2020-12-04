package entities

import (
	"encoding/json"
	"gorm.io/gorm"
	"rpcf/core/entities"
	"time"
)

const (
	ArticlesTableName = "articles"
)

type Articles struct {
	*entities.Base
	ID               string `gorm:"primaryKey"`
	Title            string
	PublishedCountry string `json:"published_country"`
	MagazineName     string `json:"maganize_name"`
	MagazineISSN     string `json:"maganize_issn"`
	Vol              string
	Fasc             string
	Pags             string
	Doi              string
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
}

func NewArticle(product map[string]string) (*Articles, error) {
	var a Articles
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
func (p *Articles) BeforeCreate(tx *gorm.DB) error {
	id := p.GenerateID()
	p.ID = id
	return nil
}

func (a *Articles) TableName() string {
	return ArticlesTableName
}
