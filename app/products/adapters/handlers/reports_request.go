package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

const (
	groupTypeParam = "groupType"
)

type ReportRequest struct {
	GroupType string
	Filters   map[string]string
}

func getSupportedParameterFiltersForReports() []string {
	return []string{"group_name", "group_code", "type_name", "title", "start_year", "end_year"}
}

func NewReportsRequest(c *gin.Context) (*ReportRequest, error) {
	groupType := c.Query(groupTypeParam)

	if funk.IsEmpty(groupType) {
		return nil, errors.New(InvalidGroupTypeParam)
	}

	filters := make(map[string]string, 0)
	supportedFilters := getSupportedParameterFiltersForReports()

	for _, paramName := range supportedFilters {
		paramValue := c.Query(paramName)

		if !funk.IsEmpty(paramValue) {
			filters[paramName] = paramValue
		}
	}

	req := &ReportRequest{
		GroupType: groupType,
		Filters:   filters,
	}

	return req, nil
}
