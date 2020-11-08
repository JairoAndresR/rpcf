package rest

func setupProductDefinitionRoutes(s *server) {
	handler, err := loadProductDefinitionsHandler()
	checkError(err)
	s.router.POST("/v1/product-definitions", handler.Create)
	s.router.GET("/v1/product-definitions/:name", handler.GetByName)
	s.router.GET("/v1/product-definitions", handler.GetAll)
	s.router.DELETE("/v1/product-definitions/:id", handler.Delete)
	s.router.PUT("/v1/product-definitions/:id", handler.Update)
}
