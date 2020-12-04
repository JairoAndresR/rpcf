package ports

import "rpcf/products"

type ProductsManager interface {
	GetAll(filters map[string]string) ([]*products.Product, error)
}
