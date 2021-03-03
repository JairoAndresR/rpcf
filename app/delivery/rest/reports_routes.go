package rest

import (
	_ "rpcf/app/dataproviders/sql"
	_ "rpcf/app/products/adapters/repositories"
	_ "rpcf/products/managers"
)

func setupReportsRoutes(s *server) {
	h, err := loadReportsHandler()
	checkError(err)
	s.router.GET("/v1/products/reports", h.CountAll)
	s.router.GET("/v1/products/reports/categories", h.CountProductsByCategory)
}
