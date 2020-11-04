package ports

import "rpcf/products"

type ProductReader interface {
	// GetAll retrieves all products
	GetAll() ([]products.Product, error)
}
