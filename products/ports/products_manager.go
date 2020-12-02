package ports

import "rpcf/products"

type ProductsManager interface {
	GetAll() ([]*products.Product, error)
}
