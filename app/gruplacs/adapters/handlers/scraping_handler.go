package handlers

import (
	"github.com/gin-gonic/gin"
	"rpcf/core"
	"rpcf/gruplacs/ports"
)

func init() {
	err := core.Injector.Provide(newScrapingHandler)
	core.CheckInjection(err, "newScrapingHandler")
}

type ScrapingHandler struct {
	manager ports.CollectorManager
}

func newScrapingHandler() *ScrapingHandler {
	return &ScrapingHandler{}
}
func (h *ScrapingHandler) Scrap(c *gin.Context) {

}
