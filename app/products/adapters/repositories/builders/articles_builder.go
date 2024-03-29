package builders

import (
	"rpcf/app/products/adapters/repositories/entities"
)

type articlesBuilder struct {
	next ProductSelector
}

func NewArticlesBuilder() ProductSelector {
	return &articlesBuilder{}
}
func (b *articlesBuilder) Build(product map[string]string, name string) (interface{}, error) {

	if !b.IsProduct(name) {
		return b.next.Build(product, name)
	}

	p, err := entities.NewArticle(product)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (b *articlesBuilder) IsProduct(name string) bool {
	return entities.ArticlesTableName == name
}
func (b *articlesBuilder) SetNext(s ProductSelector) {
	b.next = s
}
