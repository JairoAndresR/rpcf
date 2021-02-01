package rest

import (
	_ "rpcf/app/dataproviders/queue"
	_ "rpcf/app/gruplacs/adapters/repositories"
	_ "rpcf/gruplacs/managers"
)

func setupScrapingRoutes(s *server) {
	setupProductScrapingRoutes(s)
}
func setupProductScrapingRoutes(s *server) {
	handler, err := loadProductScrapingHandler()
	checkError(err)

	auth := s.authorizer.AuthorizeAdmin()

	scrapingGroup := s.router.Group("/v1/collectors/product/scraping")
	scrapingGroup.Use(auth)
	scrapingGroup.POST("", handler.Scrap)
}
