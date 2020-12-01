package ports

import "rpcf/products"

type ProductWriter interface {

	//Create is in charge to save for first time a product
	Create(p *products.Product) (*products.Product, error)
}
