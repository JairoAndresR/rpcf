package rest

func setupAuthorsRoutes(s *server) {
	handler, err := loadAuthorsHandler()
	checkError(err)
	s.router.GET("/v1/authors", handler.GetAll)
}
