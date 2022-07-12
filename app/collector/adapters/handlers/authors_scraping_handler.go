package handlers

import (
	"fmt"
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

func newAuthorsScrapingHandler(manager ports.AuthorsCollectorManager) *AuthorsScrapingHandler {
	return &AuthorsScrapingHandler{manager: manager}
}

func (h *AuthorsScrapingHandler) Scrap(c *gin.Context) {
	err := h.manager.CollectAll()

	if err != nil {
		fmt.Println(err)
		c.Error(err)
		return
	}

	res := NewScrapingResponse()
	c.JSON(http.StatusOK, res)
}
