package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpcf/collector/ports"
	"rpcf/core"
)

func init() {
	err := core.Injector.Provide(newProductScrapingHandler)
	core.CheckInjection(err, "newProductScrapingHandler")
}

type ProductsScrapingHandler struct {
	manager ports.GrupLACCollectorManager
}

func newProductScrapingHandler(m ports.GrupLACCollectorManager) *ProductsScrapingHandler {
	return &ProductsScrapingHandler{manager: m}
}
func (h *ProductsScrapingHandler) Scrap(c *gin.Context) {
	err := h.manager.CollectAll()

	if err != nil {
		// TODO: add properly error.
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
