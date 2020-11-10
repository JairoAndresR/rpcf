package handlers

import (
	"rpcf/gruplacs"
	"time"
)

type GruplacDefinitionResponse struct{
	ID         string `json:"id"`
	Name       string `json:"name"`
	URL string `json:"definition"`
	CreatedAt  *time.Time  `json:"created_at"`
	UpdatedAt  *time.Time  `json:"updated_at"`
}

type GruplacDefinitionListResponse struct {
	definitions []*GruplacDefinitionResponse `json:"definitions"`
	total int `json:"total"`
}

func newGruplacDefinitionResponse(g *gruplacs.GrupLAC) *GruplacDefinitionResponse{
	return &GruplacDefinitionResponse{
		ID: g.ID,
		Name: g.Name,
		URL: g.URL,
		CreatedAt:  g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}

func newGruplacDefinitionListResponse(g []*gruplacs.GrupLAC)  *GruplacDefinitionListResponse{
	definitions := make([]*GruplacDefinitionResponse, 0)

	for _, d := range g {
		definitions = append(definitions, newGruplacDefinitionResponse(d))
	}
	return &GruplacDefinitionListResponse{
		definitions: definitions,
		total: len(definitions),
	}
}