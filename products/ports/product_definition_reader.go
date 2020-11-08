package ports

import "rpcf/products"

type ProductDefinitionReader interface {
	// GetAll retrieves all products
	GetAll() ([]*products.ProductDefinition, error)

	// GetById retrieves a product if exist.
	GetById(id string) (*products.ProductDefinition, error)

	// GetByName retrieves a product if exist according with the name
	GetByName(name string) (*products.ProductDefinition, error)
}
