package rest

func setupProductDefinitionRoutes(s *server) {
	handler, err := loadProductDefinitionsHandler()
	checkError(err)

	auth := s.authorizer.AuthorizeAdmin()

	productDefinitionGroup := s.router.Group("/v1/product-definitions")
	productDefinitionGroup.Use(auth)

	productDefinitionGroup.POST("", handler.Create)
	productDefinitionGroup.GET("/:id", handler.GetById)
	productDefinitionGroup.GET("", handler.GetAll)
	productDefinitionGroup.DELETE("/:id", handler.Delete)
	productDefinitionGroup.PUT("/:id", handler.Update)
}
