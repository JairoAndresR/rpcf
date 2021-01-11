package rest

func setupGruplacDefinitionRoutes(s *server) {
	handler, err := loadGruplacDefinitionsHandler()
	checkError(err)

	auth := s.authorizer.AuthorizeAdmin()
	gruplacDefinitionGroup := s.router.Group("/v1/gruplac-definitions")
	gruplacDefinitionGroup.Use(auth)

	gruplacDefinitionGroup.POST("", handler.Create)
	gruplacDefinitionGroup.GET("", handler.GetAll)
	gruplacDefinitionGroup.GET("/:id", handler.GetById)
	gruplacDefinitionGroup.DELETE("/:id", handler.Delete)
	gruplacDefinitionGroup.PUT("/:id", handler.Update)
}
