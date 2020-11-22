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
	booksBuilder := builders.NewBooksBuilder()
	booksBuilder.SetNext(articlesBuilder)

	unIdentifiedBuilder := builders.NewUnIdentifiedBuilder()
	booksBuilder.SetNext(unIdentifiedBuilder)

	return booksBuilder.Build(product, name)
}
