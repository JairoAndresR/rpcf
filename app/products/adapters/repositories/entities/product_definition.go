package entities

import (
	"gorm.io/gorm"
	"rpcf/core/entities"
	"rpcf/products"
	"time"
)

type ProductDefinition struct {
	*entities.Base
	ID         string `gorm:primaryKey"`
	Name       string `gorm:"unique;not null"`
	Definition string `gorm:"type:text;not null"`
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

func (p *ProductDefinition) BeforeCreate(tx *gorm.DB) error {
	id := p.GenerateID()
	p.ID = id
	return nil
}

func NewProductDefinition(p *products.ProductDefinition) *ProductDefinition {
	return &ProductDefinition{
		ID:         p.ID,
		Name:       p.Name,
		Definition: p.Definition,
	}
}

func (p *ProductDefinition) ToDomain() products.ProductDefinition {
	return products.ProductDefinition{
		ID:         p.ID,
		Name:       p.Name,
		Definition: p.Definition,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func (p *ProductDefinition) GetDomainReference() *products.ProductDefinition {
	return &products.ProductDefinition{
		ID:         p.ID,
		Name:       p.Name,
		Definition: p.Definition,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func MapProductDefinitionListToDomain(list []*ProductDefinition) []*products.ProductDefinition {
	result := make([]*products.ProductDefinition, 0)
	for _, p := range list {
		result = append(result, p.GetDomainReference())
	}
	return result
}
