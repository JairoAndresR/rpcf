package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpcf/accounts/ports"
	"rpcf/core"
	"rpcf/core/handlers"
)

func init() {
	err := core.Injector.Provide(newRegisterUserHandler)
	core.CheckInjection(err, "RegisterUserHandler")
}

type RegisterUserHandler struct {
	manager ports.RegisterManager
}

func newRegisterUserHandler(manager ports.RegisterManager) *RegisterUserHandler {
	return &RegisterUserHandler{manager: manager}
}

func (h *RegisterUserHandler) Create(c *gin.Context) {

	req, err := newLoginRequestFromContext(c)

	if err != nil {
		GenerateError(c, http.StatusBadRequest, err)
		return
	}

	error := req.IsValid()

	if error != nil {
		handlers.GenerateFullError(c, error)
		return
	}

	account := req.ToDomain()
	acc, token, err := h.manager.Register(account)

	if err != nil {
		GenerateError(c, http.StatusUnprocessableEntity, err)
		return
	}

	response := newRegisterUserResponse(acc, token)
	c.JSON(http.StatusCreated, response)
}
