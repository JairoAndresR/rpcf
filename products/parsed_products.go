package products

type ParsedProducts struct {
	Name        string
	GrupLACCode string
	GroupName   string
	Results     []map[string]string
}

func NewParsedProducts(results []map[string]string, name, gruplacCode, gruplacName string) *ParsedProducts {
	return &ParsedProducts{
		Results:     results,
		Name:        name,
		GrupLACCode: gruplacCode,
		GroupName:   gruplacName,
	}
}
