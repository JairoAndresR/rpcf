package entities

import (
	"gorm.io/gorm"
	"rpcf/core/entities"
	"rpcf/gruplacs"
	"time"
)

type GruplacDefinition struct {
	*entities.Base
	ID        string `gorm:"primaryKey"`
	Name      string `gorm:"unique; not null"`
	URL       string `gorm:"not null"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (g *GruplacDefinition) BeforeCreate(tx *gorm.DB) error {
	id := g.GenerateID()
	g.ID = id
	return nil
}

func NewGruplacDefinition(g *gruplacs.GruplacDefinition) *GruplacDefinition {
	return &GruplacDefinition{
		ID:        g.ID,
		Name:      g.Name,
		URL:       g.URL,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}

func (g *GruplacDefinition) ToDomain() gruplacs.GruplacDefinition {
	return gruplacs.GruplacDefinition{
		ID:        g.ID,
		Name:      g.Name,
		URL:       g.URL,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}

func (g *GruplacDefinition) GetDomainReference() *gruplacs.GruplacDefinition {
	return &gruplacs.GruplacDefinition{
		ID:        g.ID,
		Name:      g.Name,
		URL:       g.URL,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}

func MapGruplacDefinitionListToDomain(list []*GruplacDefinition) []*gruplacs.GruplacDefinition {
	result := make([]*gruplacs.GruplacDefinition, 0)
	for _, g := range list {
		result = append(result, g.GetDomainReference())
	}
	return result
}
