package handlers

import "rpcf/products"

type ProductDefinitionResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Definition string `json:"definition"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
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
