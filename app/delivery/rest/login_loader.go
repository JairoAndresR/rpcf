package rest

import (
	"rpcf/app/accounts/adapters/handlers"
	"rpcf/core"
)

func loadLoginHandler() (*handlers.LoginHandler, error) {
	var handler *handlers.LoginHandler
	invokeFun := func(h *handlers.LoginHandler) {
		handler = h
	}
	err := core.Injector.Invoke(invokeFun)
	return handler, err
}

func loadRegisterHandler() (*handlers.RegisterUserHandler, error) {
	var handler *handlers.RegisterUserHandler
	invokeFun := func(h *handlers.RegisterUserHandler) {
		handler = h
	}
	err := core.Injector.Invoke(invokeFun)
	return handler, err
}
