package ports

import "rpcf/products"

type AuthorsLinker interface {
	Link(authors []*products.Author, grupLACCode string) []*products.Author
}
