package entities

import "rpcf/products"

type Author struct {
	ID       string `gorm:"primaryKey"`
	Names    string
	Products []*Product `gorm:"many2many:authors_products;"`
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