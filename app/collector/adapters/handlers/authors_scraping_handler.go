package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpcf/collector/ports"
	"rpcf/core"
)

func init() {
	err := core.Injector.Provide(newAuthorsScrapingHandler)
	core.CheckInjection(err, "newAuthorsScrapingHandler")
}

type AuthorsScrapingHandler struct {
	manager ports.AuthorsCollectorManager
}

func newAuthorsScrapingHandler() *AuthorsScrapingHandler {
	return &AuthorsScrapingHandler{}
}

func (h *AuthorsScrapingHandler) Scrap(c *gin.Context) {
	err := h.manager.CollectAll()

	if err != nil {
		c.Error(err)
		return
	}

	res := NewScrapingResponse()
	c.JSON(http.StatusOK, res)
}
