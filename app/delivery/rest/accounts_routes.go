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

	auth, err := loadAuthHandler()

	s.router.POST("/v1/accounts/login", login.Login)
	s.router.POST("/v1/accounts/register", register.Create)
	s.router.POST("/v1/accounts/validate", auth.ValidateToken)
}
