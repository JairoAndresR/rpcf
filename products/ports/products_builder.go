package ports

type ProductsBuilder interface {

	// Build creates a instance of a specific product according with the name
	Build(name string) (interface{}, error)
}
