package ports

import "rpcf/products"

type AuthorsManager interface {
	GetAll() ([]*products.Author, error)
}
