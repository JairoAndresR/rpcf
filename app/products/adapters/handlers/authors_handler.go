package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpcf/core"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(NewAuthorsHandlers)
	core.CheckInjection(err, "NewAuthorsHandlers")
}

type AuthorsHandler struct {
	manager ports.AuthorsManager
}

func NewAuthorsHandlers(manager ports.AuthorsManager) *AuthorsHandler {
	return &AuthorsHandler{
		manager: manager,
	}
}

func (h *AuthorsHandler) GetAll(c *gin.Context) {
	req := NewAuthorsRequest(c)

	list, err := h.manager.GetAll(req.Filters)

	if err != nil {
		generateError(c, http.StatusUnprocessableEntity, err)
		return
	}
	response := NewAuthorListResponse(list)
	c.JSON(http.StatusOK, response)
}
