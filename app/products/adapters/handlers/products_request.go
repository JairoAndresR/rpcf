package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
	"github.com/thoas/go-funk"
)

const (
	titleParam      = "title"
	groupParam      = "group"
	researcherParam = "researcher"
	startYearParam  = "startYear"
	endYearParam    = "endYear"
	typeNameParam   = "typeName"
)

type ProductsRequest struct {
	Filters map[string]string
}

func NewProductRequest(c *gin.Context) *ProductsRequest {
	filters := make(map[string]string, 0)

	title := c.Query(titleParam)
	group := c.Query(groupParam)
	researcher := c.Query(researcherParam)
	startYear := c.Query(startYearParam)
	endYear := c.Query(endYearParam)
	typeName := c.Query(typeNameParam)

	if !funk.IsEmpty(title) {
		filters[titleParam] = title
	}

	if !funk.IsEmpty(group) {
		filters[strcase.ToSnake(groupParam)] = group
	}

	if !funk.IsEmpty(researcher) {
		filters[strcase.ToSnake(researcherParam)] = researcher
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
