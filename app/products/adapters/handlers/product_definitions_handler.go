package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpcf/core"
	"rpcf/core/handlers"
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

func (h *ProductDefinitionsHandler) Create(ctx *gin.Context) {
	req, err := newProductDefinitionCreateRequest(ctx)

	if err != nil {
		generateError(ctx, http.StatusBadRequest, err)
		return
	}

	error := req.IsValid()
	if error != nil {
		handlers.GenerateFullError(ctx, error)
		return
	}

	definition := req.GetProductDefinition()

	definition, err = h.manager.Create(definition)

	if err != nil {
		generateError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	response := newProductDefinitionResponse(definition)
	ctx.JSON(http.StatusCreated, response)

}

func (h *ProductDefinitionsHandler) GetAll(ctx *gin.Context) {
	definitions, err := h.manager.GetAll()
	if err != nil {
		generateError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	response := newProductDefinitionListResponse(definitions)
	ctx.JSON(http.StatusCreated, response)
}

func (h *ProductDefinitionsHandler) GetByName(ctx *gin.Context) {}

func (h *ProductDefinitionsHandler) Delete(ctx *gin.Context) {

}

func (h *ProductDefinitionsHandler) Update(ctx *gin.Context) {}
