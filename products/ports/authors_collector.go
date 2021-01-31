package ports

import (
	"rpcf/products"
)

type AuthorsCollector interface {
	Process(content string) ([]*products.Author, []error)
}
