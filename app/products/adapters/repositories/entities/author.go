package entities

import (
	"gorm.io/gorm"
	"rpcf/core/entities"
	"rpcf/products"
	"strings"
)

type Author struct {
	*entities.Base
	ID       string `gorm:"primaryKey"`
	Names    string
	GrupLACS []GrupLAC  `gorm:"many2many:authors_gruplacs;"`
	Products []*Product `gorm:"many2many:authors_products;"`
}

func NewAuthor(a *products.Author) *Author {
	return &Author{
		ID:    a.ID,
		Names: a.Names,
	}
}
func (a *Author) BeforeCreate(tx *gorm.DB) error {
	id := a.GenerateID()
	a.ID = id
	a.Names = strings.ToUpper(a.Names)
	return nil
}

func (a *Author) ToDomain() *products.Author {
	return &products.Author{
		ID:    a.ID,
		Names: a.Names,
	}
}

func MapAuthorsToDomain(list []*Author) []*products.Author {

	authors := make([]*products.Author, 0)

	for _, a := range list {
		authors = append(authors, a.ToDomain())
	}
	return authors
}

func NewAuthors(authors []*products.Author) []*Author {

	list := make([]*Author, 0)
	for _, a := range authors {
		list = append(list, NewAuthor(a))
	}
	return list
}
