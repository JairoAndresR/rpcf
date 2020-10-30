package ports

type ProductManager interface {
	Write(product map[string]string, name string) error
}
