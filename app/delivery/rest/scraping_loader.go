package rest

import (
	"rpcf/app/collector/adapters/handlers"
	_ "rpcf/collector/managers"
	"rpcf/core"
	_ "rpcf/gruplacs/managers"
)

func loadProductScrapingHandler() (*handlers.ProductsScrapingHandler, error) {
	var handler *handlers.ProductsScrapingHandler
	invokeFun := func(h *handlers.ProductsScrapingHandler) {
		handler = h
	}
	err := core.Injector.Invoke(invokeFun)
	return handler, err
}

func loadAuthorsScrapingHandler() (*handlers.AuthorsScrapingHandler, error) {
	var handler *handlers.AuthorsScrapingHandler
	invokeFun := func(h *handlers.AuthorsScrapingHandler) {
		handler = h
	}
	err := core.Injector.Invoke(invokeFun)
	return handler, err
}
