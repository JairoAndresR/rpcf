package ports

import "rpcf/products"

type ProductDefinitionReader interface {
	// GetAll retrieves all products
	GetAll() ([]products.ProductDefinition, error)
}
