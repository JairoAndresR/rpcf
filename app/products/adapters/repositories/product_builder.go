package repositories

import (
	"errors"
	"rpcf/products/ports"
)

type productBuilder struct {
}

func newProductBuilder() ports.ProductsBuilder {
	return &productBuilder{}
}

func (p *productBuilder) Build(name string) (interface{}, error) {
	product, ok := productEntities[name]

	if ok {
		return product, nil
	}
	return nil, errors.New(UnidentifiedProductError)
}
