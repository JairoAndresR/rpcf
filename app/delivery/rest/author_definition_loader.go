package rest

import (
	"rpcf/app/products/adapters/handlers"
	_ "rpcf/app/products/adapters/repositories"
	"rpcf/core"
	_ "rpcf/products/managers"
)

func loadAuthorDefinitionsHandler() (*handlers.AuthorsDefinitionHandler, error) {
	var handler *handlers.AuthorsDefinitionHandler
	invokeFun := func(h *handlers.AuthorsDefinitionHandler) {
		handler = h
	}
	err := core.Injector.Invoke(invokeFun)
	return handler, err
}
