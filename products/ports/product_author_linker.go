package ports

import "rpcf/products"

type ProductAuthorLinker interface {
	// Link try to generate the link between products and authors.
	Link(authors []*products.Author, products []*products.ProductResult) []*products.ProductResult
}
