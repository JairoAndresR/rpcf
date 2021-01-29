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
		names := strings.ToLower(a.Names)
		names = strings.TrimSpace(names)
		authorsIndex[names] = a
	}

	for _, p := range ps {
		fieldAuthors := p.Fields["autores"]
		authorsNames := strings.Split(fieldAuthors, ",")
		authorsResult := make([]*products.Author, 0)

		for _, name := range authorsNames {
			name = strings.ReplaceAll(strings.ToLower(name), ",", "")
			name = strings.TrimSpace(name)
			foundAuthor, exist := authorsIndex[name]

			if exist {
				authorsResult = append(authorsResult, foundAuthor)
			}
		}

		p.Authors = authorsResult
	}

	return ps
}
