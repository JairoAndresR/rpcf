package handlers

import (
	"net/http"
	"rpcf/core"
	"rpcf/products/ports"

	"github.com/gin-gonic/gin"
)

func init() {
	err := core.Injector.Provide(NewReportsHandler)
	core.CheckInjection(err, "NewReportsHandler")
}

type ReportsHandler struct {
	manager ports.ReportsManager
}

func NewReportsHandler(manager ports.ReportsManager) *ReportsHandler {
	return &ReportsHandler{
		manager: manager,
	}
}

func (h *ReportsHandler) CountAll(c *gin.Context) {
	req, err := NewReportsRequest(c)

	if err != nil {
		generateError(c, http.StatusBadRequest, err)
		return
	}

	list, err := h.manager.CountAll(req.Filters, req.GroupType)

	if err != nil {
		generateError(c, http.StatusUnprocessableEntity, err)
		return
	}
	response := NewListReportResponse(list)
	c.JSON(http.StatusOK, response)
}

func (h *ReportsHandler) CountProductsByCategory(c *gin.Context) {

	list := h.manager.CountProductsByCategory()

	response := NewListReportResponse(list)
	c.JSON(http.StatusOK, response)
}
