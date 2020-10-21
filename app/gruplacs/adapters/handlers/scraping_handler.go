package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func newScrapingHandler(m ports.CollectorManager) *ScrapingHandler {
	return &ScrapingHandler{manager: m}
}
func (h *ScrapingHandler) Scrap(c *gin.Context) {
	err := h.manager.CollectAll()

	if err != nil {
		// TODO: add properly error.
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
