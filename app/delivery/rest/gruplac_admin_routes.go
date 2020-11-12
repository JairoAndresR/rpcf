package rest

func setupGruplacDefinitionRoutes(s *server) {
	handler, err := loadGruplacDefinitionsHandler()
	checkError(err)
	s.router.POST("/v1/gruplac-definitions", handler.Create)
	s.router.GET("/v1/gruplac-definitions", handler.GetByName)
	s.router.GET("/v1/gruplac-definitions/:id", handler.GetAll)
	s.router.DELETE("/v1/gruplac-definitions/:id", handler.Delete)
	s.router.PUT("/v1/gruplac-definitions/:id", handler.Update)
}
