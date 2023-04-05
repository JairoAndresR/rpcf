package managers

import (
	analysisPorts "rpcf/analysis/ports"
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"
	"strings"
)

func init() {
	err := core.Injector.Provide(newReportsManager)
	core.CheckInjection(err, "newReportsManager")
}

type ReportsManager struct {
	reader           ports.ReportsReader
	productReader    ports.ProductReader
	frequencyCounter analysisPorts.WordFrequencyCounter
}

func newReportsManager(r ports.ReportsReader, pr ports.ProductReader, fc analysisPorts.WordFrequencyCounter) ports.ReportsManager {
	return &ReportsManager{
		reader:           r,
		productReader:    pr,
		frequencyCounter: fc,
	}
}

func (m *ReportsManager) CountAll(filters map[string]string, group string) ([]*products.Report, error) {
	return m.reader.CountAllBy(filters, group)
}

func (m *ReportsManager) CountProductsByCategory() []*products.Report {
	return m.reader.CountProductsByCategory()
}

func (m *ReportsManager) WordsFrequency(filters map[string]string) ([]*products.Report, error) {
	ps, err := m.productReader.GetAll(filters)

	if err != nil {
		return nil, err
	}

	var textBuilder strings.Builder
	for _, p := range ps {
		textBuilder.WriteString(strings.ToLower(p.Title))
	}

	text := textBuilder.String()
	frequencies := m.frequencyCounter.Count(text)

	reports := make([]*products.Report, 0)
	for _, word := range frequencies {
		report := products.NewReport(word.Key, int64(word.Value))
		if word.Key != "" {
			reports = append(reports, report)
		}
	}

	return reports, nil
}
