package ports

import "rpcf/products"

type ReportsReader interface {
	CountAllBy(filters map[string]string, group string) ([]*products.Report, error)
}
