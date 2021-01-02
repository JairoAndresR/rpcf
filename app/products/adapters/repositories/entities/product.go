package entities

import (
	"rpcf/products"
	"time"
)

type Product struct {
	ID        string `gorm:"primaryKey"`
	GroupId   string `gorm:"type:varchar(36)"`
	TypeId    string `gorm:"unique"`
	TypeName  string
	Title     string
	StartYear string
	EndYear   string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Authors   []Author `gorm:"many2many:authors_products;"`
}

func NewProduct(p *products.Product) *Product {
	return &Product{
		ID:        p.ID,
		GroupId:   p.GroupId,
		TypeId:    p.TypeId,
		TypeName:  p.TypeName,
		Title:     p.Title,
		StartYear: p.StartYear,
		EndYear:   p.EndYear,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (p *Product) ToDomain() *products.Product {
	return &products.Product{
		ID:        p.ID,
		GroupId:   p.GroupId,
		TypeId:    p.TypeId,
		TypeName:  p.TypeName,
		Title:     p.Title,
		StartYear: p.StartYear,
		EndYear:   p.EndYear,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
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
