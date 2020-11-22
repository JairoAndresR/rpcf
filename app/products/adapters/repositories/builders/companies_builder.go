package builders

import "rpcf/app/products/adapters/repositories/entities"

type companiesBuilder struct {
	next ProductSelector
}

func NewCompaniesBuilder() ProductSelector {
	return &companiesBuilder{}
}
func (c *companiesBuilder) Build(product map[string]string, name string) (interface{}, error) {
	if !c.IsProduct(name) {
		return c.next.Build(product, name)
	}

	p, err := entities.NewCompanies(product)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (c *companiesBuilder) IsProduct(name string) bool {
	return entities.CompaniesTable == name
}

func (c *companiesBuilder) SetNext(s ProductSelector) {
	c.next = s
}
