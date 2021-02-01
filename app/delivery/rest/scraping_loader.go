package rest

import (
	"rpcf/app/collector/adapters/handlers"
	"rpcf/core"
)

func loadProductScrapingHandler() (*handlers.ProductsScrapingHandler, error) {
	var handler *handlers.ProductsScrapingHandler
	invokeFun := func(h *handlers.ProductsScrapingHandler) {
		handler = h
	}
	err := core.Injector.Invoke(invokeFun)
	return handler, err
}
