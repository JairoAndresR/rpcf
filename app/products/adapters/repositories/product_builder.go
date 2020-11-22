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
	companiesBuilder := builders.NewCompaniesBuilder()

	unIdentifiedBuilder := builders.NewUnIdentifiedBuilder()

	articlesBuilder.SetNext(booksBuilder)
	booksBuilder.SetNext(companiesBuilder)
	// The unIdentifiedBuilder must be the last assigned
	companiesBuilder.SetNext(unIdentifiedBuilder)

	return articlesBuilder.Build(product, name)
}
