package ports

import "rpcf/products"

type AuthorsDefinitionReader interface {
	// GetAuthorDefinition retrieve the author definition
	GetAuthorDefinition() (*products.AuthorDefinition, error)
}
