package handlers

import (
	"github.com/gin-gonic/gin"
	"rpcf/core"
)

func init() {
	err := core.Injector.Provide(NewProductsHandler)
	core.CheckInjection(err, "NewProductsHandler")
}

type ProductsHandler struct {
}

func NewProductsHandler() *ProductsHandler {
	return &ProductsHandler{}
}

func (h *ProductsHandler) GetAll(c *gin.Context) {

}
