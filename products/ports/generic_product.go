package ports

import (
	"encoding/json"
	"rpcf/products"
)

// Represents the generic information of any kind of product in the system
type GenericProduct interface {
	GetID() string
}

func NewGenericProduct(generic interface{}, p *products.ProductResult) *products.Product {
	b, err := json.Marshal(generic)

	if err != nil {
		return nil
	}

	var product products.Product
	err = json.Unmarshal(b, &product)

	if err != nil {
		return nil
	}
	product.TypeId = product.ID
	product.TypeName = p.Name
	product.GrouplacCode = p.GrupLACCode
	product.GroupName = p.GrupLACName
	product.Authors = p.Authors
	return &product
}
