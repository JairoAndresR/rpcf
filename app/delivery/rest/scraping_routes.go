package rest

import (
	_ "rpcf/app/dataproviders/queue"
	_ "rpcf/app/gruplacs/adapters/repositories"
	_ "rpcf/gruplacs/managers"
)

func setupScrapingRoutes(s *server) {
	handler, err := loadScrapingHandler()
	checkError(err)

	auth := s.authorizer.AuthorizeAdmin()

	scrapingGroup := s.router.Group("/v1/gruplacs/scraping")
	scrapingGroup.Use(auth)
	scrapingGroup.POST("", handler.Scrap)
}
