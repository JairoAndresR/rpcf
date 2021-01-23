package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
	"github.com/thoas/go-funk"
)

type AuthorRequest struct {
	Filters map[string]string
}

func NewAuthorsRequest(c *gin.Context) *AuthorRequest {
	filters := make(map[string]string, 0)

	groupCode := c.Query(groupCodeParam)
	if !funk.IsEmpty(groupCode) {
		filters[strcase.ToSnake(groupCodeParam)] = groupCode
	}
	return &AuthorRequest{Filters: filters}
}
