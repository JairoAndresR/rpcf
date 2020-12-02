package rest

func setupProductsRoutes(s *server) {
	handler, err := loadProductsHandler()
	checkError(err)
	s.router.GET("/v1/products", handler.GetAll)
}
