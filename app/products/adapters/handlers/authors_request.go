package handlers

import (
	"github.com/gin-gonic/gin"
)

type AuthorRequest struct {
	GroupCode string
}

func NewAuthorsRequest(c *gin.Context) *AuthorRequest {
	groupCode := c.Query(groupCodeParam)
	return &AuthorRequest{GroupCode: groupCode}
}
