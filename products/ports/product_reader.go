package ports

import "rpcf/products"

type ProductReader interface {

	// GetAll retrieves all products in the database.
	GetAll() ([]*products.Product, error)
}
