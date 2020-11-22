package handlers

import "rpcf/products"

type ProductDefinitionResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Definition string `json:"definition"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

type ProductDefinitionListResponse struct {
	Definitions []*ProductDefinitionResponse `json:"Definitions"`
	Total       int                          `json:"Total"`
}

func newProductDefinitionResponse(p *products.ProductDefinition) *ProductDefinitionResponse {
	return &ProductDefinitionResponse{
		ID:         p.ID,
		Name:       p.Name,
		Definition: p.Definition,
		CreatedAt:  p.CreatedAt.Unix(),
		UpdatedAt:  p.UpdatedAt.Unix(),
	}
}

func newProductDefinitionListResponse(p []*products.ProductDefinition) *ProductDefinitionListResponse {
	definitions := make([]*ProductDefinitionResponse, 0)

	for _, d := range p {
		definitions = append(definitions, newProductDefinitionResponse(d))
	}

	return &ProductDefinitionListResponse{
		Definitions: definitions,
		Total:       len(definitions)}
}
