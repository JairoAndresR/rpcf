package ports

import (
	"rpcf/products"
)

type ProductCollector interface {
	Process(content string) []error

	// It retrieves the products parsed of the raw content
	Parse(content string) ([]*products.ProductResult, []error)
}
