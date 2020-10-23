package rest

import (
	_ "rpcf/app/dataproviders/queue"
	_ "rpcf/app/gruplacs/adapters/repositories"
	_ "rpcf/gruplacs/managers"
)

func setupScrapingRoutes(s *server) {
	handler, err := loadScrapingHandler()
	checkError(err)
	s.router.POST("/v1/gruplacs/scraping", handler.Scrap)
}
