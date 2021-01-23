package ports

import "rpcf/products"

type AuthorsManager interface {
	GetAll(groupCode string) ([]*products.Author, error)
}
