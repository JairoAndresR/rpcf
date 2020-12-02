package rest

import (
	"rpcf/app/products/adapters/handlers"
	"rpcf/core"
)

func loadProductsHandler() (*handlers.ProductsHandler, error) {
	var handler *handlers.ProductsHandler
	invokeFun := func(h *handlers.ProductsHandler) {
		handler = h
	}
	err := core.Injector.Invoke(invokeFun)
	return handler, err
}
