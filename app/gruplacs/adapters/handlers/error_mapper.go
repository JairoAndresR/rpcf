package handlers

import (
	"github.com/gin-gonic/gin"
	"rpcf/core/handlers"
)

func generateError(c *gin.Context, status int, err error){
	e := mapError(err, status)
	handlers.GenerateError(c, e)
}

func mapError(err error, status int) *handlers.Error {
	title := err.Error()
	error := handlers.NewErrorWithStatus(title, status)
	return error
}