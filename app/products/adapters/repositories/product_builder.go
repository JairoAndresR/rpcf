package repositories

import (
	"rpcf/app/products/adapters/repositories/builders"
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

	articlesBuilder := builders.NewArticlesBuilder()

	// The unIdentifiedBuilder must be the last builder always. 
	unIdentifiedBuilder := builders.NewUnIdentifiedBuilder()
	articlesBuilder.SetNext(unIdentifiedBuilder)

	return unIdentifiedBuilder.Build(product, name)
}
