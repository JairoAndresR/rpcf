package ports

import "rpcf/products"

type AuthorsDefinitionWriter interface {

	// Create creates a new author definition
	Create(d *products.AuthorDefinition) (*products.AuthorDefinition, error)
}
