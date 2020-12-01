package ports

// ProductGenericWriter is in charge to save any product
type ProductGenericWriter interface {

	// Write is in charge to save a product.
	// It recognizes where put the product according to the name.
	//
	// If all is ok it returns a id of the resource saved and a nil error,
	// in other case it should returns an specific error.
	WriteMap(product map[string]string, name string) error

	// It writes multiples products.
	WriteMaps(products []map[string]string, name string) []error
}
