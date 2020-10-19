package rest

import (
	"rpcf/app/gruplacs/adapters/handlers"
	"rpcf/core"
)

func loadScrapingHandler() (*handlers.ScrapingHandler, error) {
	var handler *handlers.ScrapingHandler
	invokeFun := func(h *handlers.ScrapingHandler) {
		handler = h
	}
	err := core.Injector.Invoke(invokeFun)
	return handler, err
}
