package ports

import (
	"rpcf/products"
)

type GrupLACWriter interface {
	Create(g *products.GrupLAC) (*products.GrupLAC, error)

	Update(g *products.GrupLAC) (*products.GrupLAC, error)
}
