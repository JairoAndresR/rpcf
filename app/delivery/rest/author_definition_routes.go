package rest

func setupAuthorDefinitionRoutes(s *server) {
	handler, err := loadAuthorDefinitionsHandler()
	checkError(err)

	auth := s.authorizer.AuthorizeAdmin()

	authorDefinitionGroup := s.router.Group("/v1/author-definitions")
	authorDefinitionGroup.Use(auth)
	authorDefinitionGroup.POST("", handler.Create)
	authorDefinitionGroup.GET("", handler.GetAll)
}
