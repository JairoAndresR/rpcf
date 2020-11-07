package handlers

import "github.com/gin-gonic/gin"

type ProductDefinitionsHandler struct {
}

func newProductDefinitionHandler() *ProductDefinitionsHandler {
	return &ProductDefinitionsHandler{}
}

func (h *ProductDefinitionsHandler) Create(ctx *gin.Context) {}

func (h *ProductDefinitionsHandler) GetAll(ctx *gin.Context) {}

func (h *ProductDefinitionsHandler) GetByName(ctx *gin.Context) {}

func (h *ProductDefinitionsHandler) Delete(ctx *gin.Context) {}

func (h *ProductDefinitionsHandler) Update(ctx *gin.Context) {}
