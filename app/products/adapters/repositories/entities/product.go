package entities

import (
	"rpcf/products"
	"time"
)

type Product struct {
	ID        string `gorm:"primaryKey"`
	GroupCode string `gorm:"type:varchar(36)"`
	GroupName string
	TypeId    string `gorm:"unique"`
	TypeName  string
	Title     string
	StartYear string
	EndYear   string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Authors   []*Author `gorm:"many2many:authors_products;"`
}

func NewProduct(p *products.Product) *Product {

	authors := NewAuthors(p.Authors)
	return &Product{
		ID:        p.ID,
		GroupCode: p.GrouplacCode,
		GroupName: p.GroupName,
		TypeId:    p.TypeId,
		TypeName:  p.TypeName,
		Title:     p.Title,
		StartYear: p.StartYear,
		EndYear:   p.EndYear,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Authors:   authors,
	}
}

func (p *Product) ToDomain() *products.Product {

	authors := MapAuthorsToDomain(p.Authors)
	return &products.Product{
		ID:           p.ID,
		GrouplacCode: p.GroupCode,
		GroupName:    p.GroupName,
		TypeId:       p.TypeId,
		TypeName:     p.TypeName,
		Title:        p.Title,
		StartYear:    p.StartYear,
		EndYear:      p.EndYear,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
		Authors:      authors,
	}
}

func MapListToDomain(list []*Product) []*products.Product {
	var results []*products.Product

	for _, p := range list {
		product := p.ToDomain()
		results = append(results, product)
	}
	return results
}
