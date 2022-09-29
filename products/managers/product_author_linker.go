package managers

import (
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"
	"strings"
	"unicode"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"fmt"
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
		names = removeAccents(names)
		fmt.Println(names)
		authorsIndex[names] = a
	}
	fmt.Println("bd -------------------------------------------------------------------------- pr")
	for _, p := range ps {
		fieldAuthors := p.Fields["autores"]
		authorsNames := strings.Split(fieldAuthors, ",")
		authorsResult := make([]*products.Author, 0)
		if len(authorsNames) == 0 {
			fmt.Println("no authors", p.Fields)
		}
		for _, name := range authorsNames {
			name = strings.ReplaceAll(strings.ToLower(name), ",", "")
			name = strings.TrimSpace(name)
			name = removeAccents(name)
			fmt.Println("product with autores:", p.Fields)
			fmt.Println(name)
			foundAuthor, exist := authorsIndex[name]

			if exist {
				authorsResult = append(authorsResult, foundAuthor)
			}
		}

		p.Authors = authorsResult
	}

	return ps
}

func removeAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		panic(e)
	}
	return output
}
