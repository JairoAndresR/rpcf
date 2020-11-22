package builders

type ProductSelector interface {
	Build(product map[string]string, name string) (interface{}, error)
	SetNext(s ProductSelector)
	IsProduct(name string) bool
}
