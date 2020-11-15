package ports

import (
	"rpcf/products"
)

type ProductCollector interface {

	// It retrieves the products parsed of the raw content
	Parse(content string) ([]*products.ParsedProducts, []error)
}
