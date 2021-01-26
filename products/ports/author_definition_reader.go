package ports

import "rpcf/products"

type AuthorDefinitionReader interface {
	// GetAuthorDefinition retrieve the author definition
	GetAuthorDefinition() (*products.AuthorDefinition, error)
}
