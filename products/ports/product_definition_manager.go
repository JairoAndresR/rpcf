package ports

import "rpcf/products"

type ProductDefinitionManager interface {
	Create(p *products.ProductDefinition) (*products.ProductDefinition, error)

	Update(p *products.ProductDefinition) (*products.ProductDefinition, error)

	Delete(id string) error

	GetAll() ([]*products.ProductDefinition, error)

	GetByName(name string) (*products.ProductDefinition, error)

	GetById(id string) (*products.ProductDefinition, error)
}
