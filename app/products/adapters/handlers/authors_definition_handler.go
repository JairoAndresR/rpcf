package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpcf/core"
	"rpcf/core/handlers"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newAuthorsDefinitionHandler)
	core.CheckInjection(err, "newAuthorsDefinitionHandler")
}

type AuthorsDefinitionHandler struct {
	manager ports.AuthorsDefinitionManager
}

func newAuthorsDefinitionHandler(manager ports.AuthorsDefinitionManager) *AuthorsDefinitionHandler {
	return &AuthorsDefinitionHandler{
		manager: manager,
	}
}

func (h *AuthorsDefinitionHandler) Create(ctx *gin.Context) {
	req, err := newAuthorDefinitionCreateRequest(ctx)

	if err != nil {
		generateError(ctx, http.StatusBadRequest, err)
		return
	}

	error := req.IsValid()
	if error != nil {
		handlers.GenerateFullError(ctx, error)
		return
	}

	definition := req.GetAuthorDefinition()
	definition, err = h.manager.Create(definition)

	if err != nil {
		generateError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	response := newAuthorDefinitionResponse(definition)
	ctx.JSON(http.StatusCreated, response)
}
