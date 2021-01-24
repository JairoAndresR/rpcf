package handlers

import (
	"rpcf/products"
	"time"
)

type ProductResponse struct {
	ID        string            `json:"id"`
	GroupCode string            `json:"group_code"`
	TypeId    string            `json:"type_id"`
	TypeName  string            `json:"type_name"`
	Title     string            `json:"title"`
	StartYear string            `json:"start_year"`
	EndYear   string            `json:"end_year"`
	Authors   []*AuthorResponse `json:"authors"`
	CreatedAt *time.Time        `json:"created_at"`
	UpdatedAt *time.Time        `json:"updated_at"`
}

func NewProductResponse(p *products.Product) *ProductResponse {
	return &ProductResponse{
		ID:        p.ID,
		GroupCode: p.GrouplacCode,
		TypeId:    p.TypeId,
		TypeName:  p.TypeName,
		Title:     p.Title,
		StartYear: p.StartYear,
		EndYear:   p.EndYear,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
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

func NewProductListResponse(list []*products.Product) *ProductListResponse {
	return &ProductListResponse{
		Products: MapListResponse(list),
		Total:    len(list), // TODO: change this value from the query result
	}
}
