package builders

import (
	"errors"
)

const (
	UnidentifiedProductError = "unidentified product"
)

type unIdentifiedBuilder struct {
}

func NewUnIdentifiedBuilder() ProductSelector {
	return &unIdentifiedBuilder{}
}

func (u unIdentifiedBuilder) Build(product map[string]string, name string) (interface{}, error) {
	return nil, errors.New(UnidentifiedProductError)
}

func (u unIdentifiedBuilder) SetNext(s ProductSelector) {

}
func (u unIdentifiedBuilder) IsProduct(name string) bool {
	panic("implement me")
}
