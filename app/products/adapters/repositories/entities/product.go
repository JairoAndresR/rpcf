package entities

import (
	"rpcf/products"
	"time"
)

const (
	ProductTableName = "products_fact"
)

type Product struct {
	ID              string `gorm:"PRIMARY_KEY"`
	SKResearcher    string `gorm:"type:varchar(36)"`
	SKResearchGroup string `gorm:"type:varchar(36)"`
	TypeId          string `gorm:"UNIQUE"`
	TypeName        string
	Title           string
	StartYear       string
	EndYear         string
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}

func (Product) TableName() string {
	return ProductTableName
}
func NewProduct(p *products.Product) *Product {
	return &Product{
		ID:              p.ID,
		SKResearcher:    p.SKResearcher,
		SKResearchGroup: p.SKResearchGroup,
		TypeId:          p.TypeId,
		TypeName:        p.TypeName,
		Title:           p.Title,
		StartYear:       p.StartYear,
		EndYear:         p.EndYear,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
	}
}

func (p *Product) ToDomain() *products.Product {
	return &products.Product{
		ID:              p.ID,
		SKResearcher:    p.SKResearcher,
		SKResearchGroup: p.SKResearchGroup,
		TypeId:          p.TypeId,
		TypeName:        p.TypeName,
		Title:           p.Title,
		StartYear:       p.StartYear,
		EndYear:         p.EndYear,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
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
