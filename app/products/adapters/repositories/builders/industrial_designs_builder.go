package builders

import "rpcf/app/products/adapters/repositories/entities"

type IndustrialDesignsBuilder struct {
	next ProductSelector
}

func NewIndustrialDesignsBuilder() ProductSelector {
	return &IndustrialDesignsBuilder{}
}

func (i *IndustrialDesignsBuilder) Build(product map[string]string, name string) (interface{}, error) {
	if !i.IsProduct(name) {
		return i.next.Build(product, name)
	}

	p, err := entities.NewIndustrialDesigns(product)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (i *IndustrialDesignsBuilder) SetNext(s ProductSelector) {
	i.next = s
}

func (i *IndustrialDesignsBuilder) IsProduct(name string) bool {
	return entities.IndustrialDesignsTable == name
}
