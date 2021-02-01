package rest

import (
	_ "rpcf/app/collector/adapters/repositories"
	_ "rpcf/app/dataproviders/queue"
	_ "rpcf/app/gruplacs/adapters/repositories"
	_ "rpcf/gruplacs/managers"
)

func setupScrapingRoutes(s *server) {
	setupProductScrapingRoutes(s)
	setupAuthorsScrapingRoutes(s)

}
func setupProductScrapingRoutes(s *server) {
	handler, err := loadProductScrapingHandler()
	checkError(err)

	auth := s.authorizer.AuthorizeAdmin()

	scrapingGroup := s.router.Group("/v1/collectors/product/scraping")
	scrapingGroup.Use(auth)
	scrapingGroup.POST("", handler.Scrap)
}

func setupAuthorsScrapingRoutes(s *server) {

	handler, err := loadAuthorsScrapingHandler()
	checkError(err)

	auth := s.authorizer.AuthorizeAdmin()
	scrapingGroup := s.router.Group("/v1/collectors/authors/scraping")
	scrapingGroup.Use(auth)
	scrapingGroup.POST("", handler.Scrap)
}
