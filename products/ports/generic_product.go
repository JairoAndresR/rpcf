package ports

import (
	"encoding/json"
	"rpcf/products"
)

// Represents the generic information of any kind of product in the system
type GenericProduct interface {
	GetID() string
}

func NewGenericProduct(generic interface{}, gruplacCode, typeName string) *products.Product {
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
	product.TypeName = typeName
	product.GrouplacCode = gruplacCode
	return &product
}
