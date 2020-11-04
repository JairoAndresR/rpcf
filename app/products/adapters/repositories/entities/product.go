package entities

import (
	"github.com/jinzhu/gorm"
	"rpcf/core/entities"
	"rpcf/products"
)

type Product struct {
	*entities.Base
	ID string
	Name string
	Definition string
}

func (p *Product) BeforeCreate(scope *gorm.Scope) error{
	id := p.GenerateID()
	err := scope.SetColumn("ID", id)
	if err != nil {
		return err
	}
	return nil
}

func newProduct(p products.Product) *Product{
	return &Product{
		ID: p.ID,
		Name: p.Name,
		Definition: p.Definition,
	}
}

func (p *Product) ToDomain() products.Product {
	return products.Product{
		ID: p.ID,
		Name: p.Name,
		Definition: p.Definition,
	}
}

func MapListToDomain(list []*Product) []products.Product {
	result := make([]products.Product, len(list)-1)
	for _, p := range list{
		result = append(result, p.ToDomain())
	}
	return result
}