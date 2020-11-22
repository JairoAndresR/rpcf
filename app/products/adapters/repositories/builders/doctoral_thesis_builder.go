package builders

import "rpcf/app/products/adapters/repositories/entities"

type DoctoralThesisBuilder struct {
	next ProductSelector
}

func NewDoctoralThesisBuilder() ProductSelector {
	return &DoctoralThesisBuilder{}
}

func (d *DoctoralThesisBuilder) Build(product map[string]string, name string) (interface{}, error) {

	if !d.IsProduct(name) {
		return d.next.Build(product, name)
	}

	p, err := entities.NewDoctoralThesis(product)
	if err != nil {
		return nil, err
	}

	p.ID = p.GenerateID()
	return p, nil
}

func (d *DoctoralThesisBuilder) SetNext(s ProductSelector) {
	d.next = s
}

func (d *DoctoralThesisBuilder) IsProduct(name string) bool {
	return entities.DoctoralThesisTable == name
}
