package managers

import (
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newReportsManager)
	core.CheckInjection(err, "newReportsManager")
}

type ReportsManager struct {
	reader ports.ReportsReader
}

func newReportsManager(r ports.ReportsReader) ports.ReportsManager {
	return &ReportsManager{reader: r}
}

func (m *ReportsManager) CountAll(filters map[string]string, group string) ([]*products.Report, error) {
	return m.reader.CountAllBy(filters, group)
}
