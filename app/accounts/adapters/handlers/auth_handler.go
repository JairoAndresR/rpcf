package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpcf/accounts/ports"
	"rpcf/core"
)

func init() {
	err := core.Injector.Provide(newAuthHandler)
	core.CheckInjection(err, "newAuthHandler")
}

type AuthHandler struct {
	manager ports.TokenManager
}

func newAuthHandler(manager ports.TokenManager) *AuthHandler {
	return &AuthHandler{
		manager: manager,
	}
}

func (l *AuthHandler) ValidateToken(c *gin.Context) {
	req, err := newAuthValidateToken(c)
	if err != nil {
		GenerateError(c, http.StatusBadRequest, err)
		return
	}

	isValid, err := l.manager.Validate(req.Token)
	if err != nil || !isValid {
		GenerateError(c, http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusOK, req)
}
