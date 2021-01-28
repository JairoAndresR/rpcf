package products

type ProductResult struct {
	Name        string            // Name of the product result
	GrupLACCode string            // GrupLACCode code of the research group
	GrupLACName string            // GrupLACName name of the research group
	Authors     []*Author         // Authors is the authors associated to this product
	Fields      map[string]string // Fields is a map[product_field_name]=product_field_value
}

func newProductResult(fields map[string]string, name, grupLACCode, grupLACName string) *ProductResult {
	return &ProductResult{
		Name:        name,
		GrupLACCode: grupLACCode,
		GrupLACName: grupLACName,
		Fields:      fields,
	}

}

func NewProductResults(products []map[string]string, name, grupLACCode, grupLACName string) []*ProductResult {

	results := make([]*ProductResult, 0)

	for _, field := range products {
		r := newProductResult(field, name, grupLACCode, grupLACName)
		results = append(results, r)
	}
	return results
}
