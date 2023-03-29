package ports

import (
	"rpcf/products"
)

type AuthorsWriter interface {
	Create(a *products.Author) (*products.Author, error)
	ClearDB()
}
