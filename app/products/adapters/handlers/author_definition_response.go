package handlers

import "rpcf/products"

type AuthorDefinitionResponse struct {
	ID         string `json:"id"`
	Definition string `json:"definition"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

type AuthorDefinitionListResponse struct {
	Definitions []*AuthorDefinitionResponse `json:"Definitions"`
	Total       int                         `json:"Total"`
}

func newAuthorDefinitionResponse(ad *products.AuthorDefinition) *AuthorDefinitionResponse {
	return &AuthorDefinitionResponse{
		ID:         ad.ID,
		Definition: ad.Definition,
		CreatedAt:  ad.CreatedAt.Unix(),
		UpdatedAt:  ad.UpdatedAt.Unix(),
	}
}

func newAuthorDefinitionListResponse(definitions []*products.AuthorDefinition) *AuthorDefinitionListResponse {
	results := make([]*AuthorDefinitionResponse, 0)

	for _, d := range definitions {
		results = append(results, newAuthorDefinitionResponse(d))
	}

	return &AuthorDefinitionListResponse{
		Definitions: results,
		Total:       len(results),
	}
}
