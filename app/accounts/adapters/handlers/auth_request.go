package handlers

import "github.com/gin-gonic/gin"

type ValidateToken struct {
	Token string `json:"token"`
}

func newAuthValidateToken(c *gin.Context) (*ValidateToken, error) {
	var req ValidateToken
	err := c.BindJSON(&req)

	if err != nil {
		return nil, err
	}

	return &req, nil
}
