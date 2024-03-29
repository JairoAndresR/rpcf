package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpcf/core"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(NewProductsHandler)
	core.CheckInjection(err, "NewProductsHandler")
}

type ProductsHandler struct {
	manager ports.ProductsManager
}

func NewProductsHandler(manager ports.ProductsManager) *ProductsHandler {
	return &ProductsHandler{
		manager: manager,
	}
}

func (h *ProductsHandler) GetAll(c *gin.Context) {
	req := NewProductRequest(c)
	ps, err := h.manager.GetAll(req.Filters)

	if err != nil {
		generateError(c, http.StatusUnprocessableEntity, err)
		return
	}

	response := NewProductListResponse(ps)
	c.JSON(http.StatusOK, response)
}
