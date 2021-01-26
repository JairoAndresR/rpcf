package ports

import "rpcf/products"

type AuthorDefinitionReader interface {
	// GetAuthorDefinition retrieve the only author definition
	GetAuthorDefinition() (*products.AuthorDefinition, error)
}
