package ports

type ProductGenericManager interface {
	Write(product map[string]string, name string) (GenericProduct, error)
}
