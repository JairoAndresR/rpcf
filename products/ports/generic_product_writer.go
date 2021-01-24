package ports

import "rpcf/products"

// GenericProductWriter is in charge to save any product
type GenericProductWriter interface {
	Write(product *products.Product) (*products.Product, error)

	// Write is in charge to save a product.
	// It recognizes where put the product according to the name.
	//
	// If all is ok it returns a id of the resource saved and a nil error,
	// in other case it should returns an specific error.
	WriteMap(product map[string]string, gruplacCode,groupName, productName string) (*products.Product, error)

	WriteGeneric(product map[string]string, gruplacCode,groupName, productName string) (*products.Product, error)

	// It writes multiples products.
	WriteGenerics(parsed *products.ParsedProducts) ([]*products.Product, []error)
}
