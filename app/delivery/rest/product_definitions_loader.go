package rest

import (
	"rpcf/app/products/adapters/handlers"
	_ "rpcf/app/products/adapters/repositories"
	"rpcf/core"
	_ "rpcf/products/managers"
)

func loadProductDefinitionsHandler() (*handlers.ProductDefinitionsHandler, error) {
	var handler *handlers.ProductDefinitionsHandler
	invokeFun := func(h *handlers.ProductDefinitionsHandler) {
		handler = h
	}
	err := core.Injector.Invoke(invokeFun)
	return handler, err
}
