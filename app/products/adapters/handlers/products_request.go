package handlers

import (
	"github.com/gin-gonic/gin"
)

type ProductsRequest struct {
	Filters map[string]string
}

func NewProductRequest(c *gin.Context) *ProductsRequest {
	productParams := getProductParams()
	filters := GetSupportedRequestParams(c, productParams)
	return &ProductsRequest{Filters: filters}
}
