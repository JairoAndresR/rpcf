package ports

import "rpcf/products"

type ProductGenericManager interface {
	Write(product *products.ProductResult) (GenericProduct, error)
}
