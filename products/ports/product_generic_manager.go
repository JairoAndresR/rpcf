package ports

type ProductGenericManager interface {
	Write(product map[string]string, grouplacCode, productName string) (GenericProduct, error)
}
