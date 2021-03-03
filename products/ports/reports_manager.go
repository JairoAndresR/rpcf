package ports

import "rpcf/products"

type ReportsManager interface {
	CountAll(filters map[string]string, group string) ([]*products.Report, error)
	CountProductsByCategory() []*products.Report
	WordsFrequency(filters map[string]string) ([]*products.Report, error)
}
