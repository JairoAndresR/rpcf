package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
	"github.com/thoas/go-funk"
)

type ProductParam string

const (
	titleParam        ProductParam = "title"
	groupCodeParam    ProductParam = "groupCode"
	researcherIdParam ProductParam = "researcherId"
	startYearParam    ProductParam = "startYear"
	endYearParam      ProductParam = "endYear"
	typeNameParam     ProductParam = "typeName"
)

func (p ProductParam) ToString() string {
	return fmt.Sprintf("%s", p)
}

func (p ProductParam) ToSnake() string {
	return strcase.ToSnake(fmt.Sprintf("%s", p))
}

func getProductParams() []ProductParam {
	return []ProductParam{
		titleParam,
		groupCodeParam,
		researcherIdParam,
		startYearParam,
		typeNameParam,
		endYearParam,
	}
}

func GetSupportedRequestParams(c *gin.Context, params []ProductParam) map[string]string {
	filters := make(map[string]string, 0)

	for _, p := range params {
		paramValue := c.Query(p.ToString())
		if !funk.IsEmpty(paramValue) {
			filters[p.ToSnake()] = paramValue
		}
	}

	return filters
}
