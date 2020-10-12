package rest

import (
	_ "rpcf/accounts/managers"
	_ "rpcf/app/accounts/adapters/repositories"
	_ "rpcf/app/dataproviders/sql"
)

func setupLoginRoutes(s *server) {
	login, err := loadLoginHandler()
	checkError(err)

	register, err := loadRegisterHandler()
	checkError(err)

	s.router.GET("/v1/login", login.Login)
	s.router.POST("/v1/register", register.Create)
}
