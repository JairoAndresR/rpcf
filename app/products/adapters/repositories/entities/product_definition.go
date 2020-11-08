package entities

import (
	"github.com/jinzhu/gorm"
	"rpcf/core/entities"
	"rpcf/products"
	"time"
)

type ProductDefinition struct {
	*entities.Base
	ID         string
	Name       string
	Definition string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

func (p *ProductDefinition) BeforeCreate(scope *gorm.Scope) error {
	id := p.GenerateID()
	err := scope.SetColumn("ID", id)
	if err != nil {
		return err
	}
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
