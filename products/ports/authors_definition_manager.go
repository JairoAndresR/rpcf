package ports

import "rpcf/products"

type AuthorsDefinitionManager interface {

	Create(ad *products.AuthorDefinition) (*products.AuthorDefinition, error)
}
