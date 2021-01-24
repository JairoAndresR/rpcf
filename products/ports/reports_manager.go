package ports

import "rpcf/products"

type ReportsManager interface {
	CountAll(filters map[string]string, group string) ([]*products.Report, error)
}
