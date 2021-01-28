package managers

import (
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"
	"strings"
)

func init() {
	err := core.Injector.Provide(newProductAuthorLinker)
	core.CheckInjection(err, "newProductAuthorLinker")
}

type productAuthorLinker struct {
}

func newProductAuthorLinker() ports.ProductAuthorLinker {
	return &productAuthorLinker{}
}

func (l *productAuthorLinker) Link(authors []*products.Author, ps []*products.ProductResult) []*products.ProductResult {

	authorsIndex := make(map[string]*products.Author, 0)

	for _, a := range authors {
		authorsIndex[a.Names] = a
	}

	for _, p := range ps {
		fieldAuthors := p.Fields["authors"]
		authorsNames := strings.Split(fieldAuthors, ",")
		authorsResult := make([]*products.Author, 0)

		for _, name := range authorsNames {
			name = strings.ToLower(name)
			foundAuthor, exist := authorsIndex[name]

			if exist {
				authorsResult = append(authorsResult, foundAuthor)
			}
		}

		p.Authors = authorsResult
	}

	return ps
}
