package ports

import "rpcf/products"

type GrupLACReader interface {
	GetAll() ([]*products.GrupLAC, error)
}
