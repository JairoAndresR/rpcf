package handlers

import (
	"github.com/gin-gonic/gin"
)

const (
	groupNameParam = "groupType"
)

type ReportRequest struct {
	GroupType string
	Filters   map[string]string
}

func NewReportsRequest(c *gin.Context) (*ReportRequest, error) {
	groupType := c.Query(groupNameParam)
	productParams := getProductParams()
	filters := GetSupportedRequestParams(c, productParams)
	req := &ReportRequest{
		GroupType: groupType,
		Filters:   filters,
	}
	return req, nil
}
