package rest

func setupScrapingRoutes(s *server) {
	handler, err := loadScrapingHandler()
	checkError(err)
	s.router.GET("/v1/scraping", handler.Scrap)
}
