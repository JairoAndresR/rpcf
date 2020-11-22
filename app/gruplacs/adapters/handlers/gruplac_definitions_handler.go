package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"rpcf/core"
	"rpcf/core/handlers"
	"rpcf/gruplacs/ports"
)

func init() {
	err := core.Injector.Provide(newGruplacDefinitionHandler)
	core.CheckInjection(err, "newGruplacDefinitionHandler")
}

type GruplacDefinitionsHandler struct {
	manager ports.GruplacDefinitionManager
}

func newGruplacDefinitionHandler(manager ports.GruplacDefinitionManager) *GruplacDefinitionsHandler {
	return &GruplacDefinitionsHandler{manager: manager}
}

func (h *GruplacDefinitionsHandler) Create(ctx *gin.Context) {
	req, err := newGruplacDefinitionCreateRequest(ctx)

	if err != nil {
		generateError(ctx, http.StatusBadRequest, err)
		return
	}

	error := req.IsValid()
	if error != nil {
		handlers.GenerateFullError(ctx, error)
		return
	}

	definition := req.GetGruplacDefinition()

	definition, err = h.manager.Create(definition)

	if err != nil {
		generateError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	response := newGruplacDefinitionResponse(definition)
	ctx.JSON(http.StatusCreated, response)
}

func (h *GruplacDefinitionsHandler) GetAll(ctx *gin.Context) {
	definitions, err := h.manager.GetAll()
	if err != nil {
		generateError(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	response := newGruplacDefinitionListResponse(definitions)
	ctx.JSON(http.StatusOK, response)
}

func (h *GruplacDefinitionsHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	definition, err := h.manager.GetById(id)
	if err != nil {
		generateError(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	response := newGruplacDefinitionResponse(definition)
	ctx.JSON(http.StatusOK, response)
}

func (h *GruplacDefinitionsHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if len(id) == 0 {
		err := errors.New("gruplac definition id empty")
		generateError(ctx, http.StatusNotFound, err)
		return
	}

	err := h.manager.Delete(id)
	if err != nil {
		err := errors.New("gruplac definition id doesn't exist")
		generateError(ctx, http.StatusUnprocessableEntity, err)

		return
	}
	ctx.Writer.WriteHeader(http.StatusOK)
}

func (h *GruplacDefinitionsHandler) Update(ctx *gin.Context) {
	req, err := newGruplacDefinitionUpdateRequest(ctx)

	if err != nil {
		generateError(ctx, http.StatusBadRequest, err)
		return
	}

	error := req.IsValid()
	if error != nil {
		handlers.GenerateFullError(ctx, error)
		return
	}

	definition := req.GetGruplacDefinition()

	_, err = h.manager.Update(definition)
	if err != nil {
		generateError(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	ctx.Writer.WriteHeader(http.StatusOK)
}
