package handlers

import (
	"github.com/gin-gonic/gin"
	"rpcf/core"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newProductDefinitionHandler)
	core.CheckInjection(err, "newProductDefinitionHandler")
}

type ProductDefinitionsHandler struct {
	manager ports.ProductDefinitionManager
}

func newProductDefinitionHandler(manager ports.ProductDefinitionManager) *ProductDefinitionsHandler {
	return &ProductDefinitionsHandler{manager: manager}
}

func (h *ProductDefinitionsHandler) Create(ctx *gin.Context) {}

func (h *ProductDefinitionsHandler) GetAll(ctx *gin.Context) {}

func (h *ProductDefinitionsHandler) GetByName(ctx *gin.Context) {}

func (h *ProductDefinitionsHandler) Delete(ctx *gin.Context) {}

func (h *ProductDefinitionsHandler) Update(ctx *gin.Context) {}
