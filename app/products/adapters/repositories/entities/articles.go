package entities

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"rpcf/core/entities"
	"time"
)

const (
	ArticlesTableName = "articles"
)

type Articles struct {
	*entities.Base
	ID               string `gorm:"PRIMARY_KEY"`
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
func (p *Articles) BeforeCreate(scope *gorm.Scope) error {
	id := p.GenerateID()
	err := scope.SetColumn("ID", id)
	if err != nil {
		return err
	}
	return nil
}

func (a *Articles) TableName() string {
	return ArticlesTableName
}
