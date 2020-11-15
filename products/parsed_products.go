package products

type ParsedProducts struct {
	Name        string
	GrupLACCode string
	Results     []map[string]string
}

func NewParsedProducts(results []map[string]string, name, code string) *ParsedProducts {
	return &ParsedProducts{
		Results:     results,
		Name:        name,
		GrupLACCode: code,
	}
}
