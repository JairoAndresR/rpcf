package builders

import (
	"rpcf/app/products/adapters/repositories/entities"
)

type BooksBuilder struct {
	next ProductSelector
}

func NewBooksBuilder() ProductSelector {
	return &BooksBuilder{}
}

func (b *BooksBuilder) SetNext(s ProductSelector) {
	b.next = s
}

func (b *BooksBuilder) Build(product map[string]string, name string) (interface{}, error) {
	if !b.IsProduct(name) {
		return b.next.Build(product, name)
	}

	p, err := entities.NewBook(product)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (b *BooksBuilder) IsProduct(name string) bool {
	return entities.BooksTableName == name
}
