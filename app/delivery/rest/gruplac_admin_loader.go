package rest

import (
	"rpcf/app/gruplacs/adapters/handlers"
	"rpcf/core"
)

func loadGruplacDefinitionsHandler() (*handlers.GruplacDefinitionsHandler, error) {
	var handler *handlers.GruplacDefinitionsHandler
	invokeFun := func(h *handlers.GruplacDefinitionsHandler) {
		handler = h
	}
	err := core.Injector.Invoke(invokeFun)
	return handler, err
}
