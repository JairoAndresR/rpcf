package ports

import "rpcf/products"

type AuthorsReader interface {
	GetAll() ([]*products.Author, error)

	GetAllByGroup(groupCode string) ([]*products.Author, error)
}
