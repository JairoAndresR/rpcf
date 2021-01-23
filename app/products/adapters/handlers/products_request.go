package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
	"github.com/thoas/go-funk"
)

type ProductsRequest struct {
	Filters map[string]string
}

func NewProductRequest(c *gin.Context) *ProductsRequest {
	filters := make(map[string]string, 0)

	title := c.Query(titleParam)
	groupId := c.Query(groupCodeParam)
	researcherId := c.Query(researcherIdParam)
	startYear := c.Query(startYearParam)
	endYear := c.Query(endYearParam)
	typeName := c.Query(typeNameParam)

	if !funk.IsEmpty(title) {
		filters[titleParam] = title
	}

	if !funk.IsEmpty(groupId) {
		filters[strcase.ToSnake(groupCodeParam)] = groupId
	}

	if !funk.IsEmpty(researcherId) {
		filters[strcase.ToSnake(researcherIdParam)] = researcherId
	}

	if !funk.IsEmpty(startYear) {
		filters[strcase.ToSnake(startYearParam)] = startYear
	}

	if !funk.IsEmpty(endYear) {
		filters[strcase.ToSnake(endYearParam)] = startYear
	}

	if !funk.IsEmpty(typeName) {
		filters[strcase.ToSnake(typeNameParam)] = typeName
	}

	return &ProductsRequest{Filters: filters}
}
