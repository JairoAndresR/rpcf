package handlers

import (
	"rpcf/products"
	"time"
)

type ProductResponse struct {
	ID              string     `json:"id"`
	SKResearcher    string     `json:"sk_researcher"`
	SKResearchGroup string     `json:"sk_research_group"`
	TypeId          string     `json:"type_id"`
	TypeName        string     `json:"type_name"`
	Title           string     `json:"title"`
	StartYear       string     `json:"start_year"`
	EndYear         string     `json:"end_year"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

func NewProductResponse(p *products.Product) *ProductResponse {
	return &ProductResponse{
		ID:              p.ID,
		SKResearcher:    p.SKResearcher,
		SKResearchGroup: p.SKResearchGroup,
		TypeId:          p.TypeId,
		TypeName:        p.TypeName,
		Title:           p.Title,
		StartYear:       p.StartYear,
		EndYear:         p.EndYear,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
	}
}

func MapListResponse(list []*products.Product) []*ProductResponse {
	results := make([]*ProductResponse, 0)

	for _, p := range list {
		product := NewProductResponse(p)

		results = append(results, product)
	}

	return results
}

type ProductListResponse struct {
	Products []*ProductResponse `json:"products"`
	Total    int                `json:"total"`
}
