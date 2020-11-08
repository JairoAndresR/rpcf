package ports

import "rpcf/products"

type ProductDefinitionWriter interface {
	Create(p *products.ProductDefinition) (*products.ProductDefinition, error)

	Update(p *products.ProductDefinition) (*products.ProductDefinition, error)

	Delete(id string) error
}
