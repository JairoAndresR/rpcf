package handlers

import (
	"github.com/gin-gonic/gin"
)

type ReportRequest struct {
	GroupType string
	Filters   map[string]string
}

func NewReportsRequest(c *gin.Context) (*ReportRequest, error) {
	groupType := c.Query(groupTypeParam.ToString())
	productParams := getProductParams()
	filters := GetSupportedRequestParams(c, productParams)
	req := &ReportRequest{
		GroupType: groupType,
		Filters:   filters,
	}
	return req, nil
}
