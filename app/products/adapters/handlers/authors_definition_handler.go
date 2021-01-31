package handlers

import (
	"github.com/gin-gonic/gin"
	"rpcf/core"
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

func (h *AuthorsDefinitionHandler) Create(c *gin.Context) {

}
