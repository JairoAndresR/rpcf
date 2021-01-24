package rest

import (
	"rpcf/app/products/adapters/handlers"
	"rpcf/core"
)

func loadReportsHandler() (*handlers.ReportsHandler, error) {
	var handler *handlers.ReportsHandler
	invokeFun := func(h *handlers.ReportsHandler) {
		handler = h
	}
	err := core.Injector.Invoke(invokeFun)
	return handler, err
}
