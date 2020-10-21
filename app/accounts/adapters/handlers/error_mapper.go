package handlers

import (
	"github.com/gin-gonic/gin"
	"rpcf/core/handlers"
)

func GenerateError(c *gin.Context, status int, err error) {
	e := mapError(err, status)
	handlers.GenerateError(c, e)
}

func AbortWithError(c *gin.Context, status int, err error) {
	e := mapError(err, status)
	handlers.AbortWithError(c, status, e)
}

func mapError(err error, status int) *handlers.Error {
	title := err.Error()
	error := handlers.NewErrorWithStatus(title, status)
	return error
}
