package ports

type ProductsBuilder interface {

	// Build creates a instance of a specific product according with the name
	Build(product map[string]string, name string) (interface{}, error)
}
