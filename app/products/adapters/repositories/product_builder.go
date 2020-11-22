package repositories

import (
	"errors"
	"rpcf/app/products/adapters/repositories/entities"
	"rpcf/core"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newProductBuilder)
	core.CheckInjection(err, "newProductBuilder")
}

type productBuilder struct {
}

func newProductBuilder() ports.ProductsBuilder {
	return &productBuilder{}
}

func (p *productBuilder) Build(product map[string]string, name string) (interface{}, error) {

	if entities.ArticlesTableName == name {
		p, err := entities.NewArticle(product)
		if err != nil {
			return nil, err
		}
		return p, nil
	}

	return nil, errors.New(UnidentifiedProductError)
}
