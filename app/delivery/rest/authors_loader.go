package rest

import (
	"rpcf/app/products/adapters/handlers"
	"rpcf/core"
)

func loadAuthorsHandler() (*handlers.AuthorsHandler, error) {
	var handler *handlers.AuthorsHandler
	invokeFun := func(h *handlers.AuthorsHandler) {
		handler = h
	}
	err := core.Injector.Invoke(invokeFun)
	return handler, err
}
