package ports

type ProductGenericManager interface {
	Write(product map[string]string, grouplacCode, groupName, productName string) (GenericProduct, error)
}
