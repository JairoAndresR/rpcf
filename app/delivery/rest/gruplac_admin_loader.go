package rest

import (
	"rpcf/app/gruplacs/adapters/handlers"
	_ "rpcf/app/gruplacs/adapters/repositories"
	"rpcf/core"
	_ "rpcf/gruplacs/managers"
)

func loadGruplacDefinitionsHandler() (*handlers.GruplacDefinitionsHandler, error) {
	var handler *handlers.GruplacDefinitionsHandler
	invokeFun := func(h *handlers.GruplacDefinitionsHandler) {
		handler = h
	}
	err := core.Injector.Invoke(invokeFun)
	return handler, err
}
