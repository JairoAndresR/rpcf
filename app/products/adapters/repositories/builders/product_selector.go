package builders

type ProductSelector interface {
	SetNext(s ProductSelector)

	Build(product map[string]string, name string) (interface{}, error)
}
