package handlers

import "rpcf/products"

type ReportResponse struct {
	Value string `json:"value"`
	Count int64  `json:"count"`
}

func newReportResponse(r *products.Report) *ReportResponse {
	return &ReportResponse{
		Value: r.Value,
		Count: r.Count,
	}
}

type ReportListResponse struct {
	Type    string            `json:"type"`
	Reports []*ReportResponse `json:"reports"`
}

func NewListReportResponse(list []*products.Report) *ReportListResponse {
	reports := make([]*ReportResponse, 0)

	for _, r := range list {
		report := newReportResponse(r)
		reports = append(reports, report)
	}

	return &ReportListResponse{
		Reports: reports,
	}
}
