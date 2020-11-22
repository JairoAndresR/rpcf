package builders

import (
	"errors"
	"rpcf/app/products/adapters/repositories"
)

type unIdentifiedBuilder struct {
}

func NewUnIdentifiedBuilder() repositories.ProductSelector {
	return &unIdentifiedBuilder{}
}

func (u unIdentifiedBuilder) SetNext(s repositories.ProductSelector) {}

func (u unIdentifiedBuilder) Build(product map[string]string, name string) (interface{}, error) {
	return nil, errors.New(repositories.UnidentifiedProductError)
}
