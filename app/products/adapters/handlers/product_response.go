package handlers

import (
	"rpcf/products"
	"time"
)

type ProductResponse struct {
	ID        string            `json:"id"`
	GroupCode string            `json:"groupCode"`
	GroupName string            `json:"groupName"`
	TypeId    string            `json:"typeId"`
	TypeName  string            `json:"typeName"`
	Title     string            `json:"title"`
	StartYear string            `json:"startYear"`
	EndYear   string            `json:"endYear"`
	Authors   []*AuthorResponse `json:"authors"`
	CreatedAt *time.Time        `json:"createdAt"`
	UpdatedAt *time.Time        `json:"updatedAt"`
}

func NewProductResponse(p *products.Product) *ProductResponse {
	authors := NewAuthorListResponse(p.Authors)
	return &ProductResponse{
		ID:        p.ID,
		GroupCode: p.GrouplacCode,
		GroupName: p.GroupName,
		TypeId:    p.TypeId,
		TypeName:  p.TypeName,
		Title:     p.Title,
		StartYear: p.StartYear,
		EndYear:   p.EndYear,
		Authors: authors.Authors,
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
