package entities

import (
	"github.com/jinzhu/gorm"
	"rpcf/core/entities"
	"rpcf/gruplacs"
	"time"
)

type GruplacDefinition struct {
	*entities.Base
	ID        string
	Name      string
	URL       string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (g *GruplacDefinition) BeforeCreate(scope *gorm.Scope) error {
	id := g.GenerateID()
	err := scope.SetColumn("ID", id)
	if err != nil {
		return err
	}
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
	result := make([]*gruplacs.GruplacDefinition, len(list)-1)
	for _, g := range list {
		result = append(result, g.GetDomainReference())
	}
	return result
}
