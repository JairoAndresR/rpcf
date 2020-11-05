package entities

import (
	"github.com/jinzhu/gorm"
	"rpcf/core/entities"
	"rpcf/products"
)

type ProductDefinition struct {
	*entities.Base
	ID string
	Name string
	Definition string
}

func (p *ProductDefinition) BeforeCreate(scope *gorm.Scope) error{
	id := p.GenerateID()
	err := scope.SetColumn("ID", id)
	if err != nil {
		return err
	}
	return nil
}

func newProductDefinition(p products.ProductDefinition) *ProductDefinition {
	return &ProductDefinition{
		ID: p.ID,
		Name: p.Name,
		Definition: p.Definition,
	}
}

func (p *ProductDefinition) ToDomain() products.ProductDefinition {
	return products.ProductDefinition{
		ID: p.ID,
		Name: p.Name,
		Definition: p.Definition,
	}
}

func MapListToDomain(list []*ProductDefinition) []products.ProductDefinition {
	result := make([]products.ProductDefinition, len(list)-1)
	for _, p := range list{
		result = append(result, p.ToDomain())
	}
	return result
}