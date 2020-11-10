package entities

import (
	"github.com/jinzhu/gorm"
	"rpcf/core/entities"
	"rpcf/gruplacs"
	"time"
)

type GrupLAC struct {
	*entities.Base
	ID        string
	Name      string
	URL       string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (g *GrupLAC) BeforeCreate(scope *gorm.Scope) error {
	id := g.GenerateID()
	err := scope.SetColumn("ID", id)
	if err != nil {
		return err
	}
	return nil
}

func NewGrupLAC(g gruplacs.GruplacDefinition) *GrupLAC {
	return &GrupLAC{
		ID:        g.ID,
		Name:      g.Name,
		URL:       g.URL,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}

func (g *GrupLAC) ToDomain() gruplacs.GruplacDefinition {
	return gruplacs.GruplacDefinition{
		ID:        g.ID,
		Name:      g.Name,
		URL:       g.URL,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}
func MapListToDomain(list []*GrupLAC) []gruplacs.GruplacDefinition {
	result := make([]gruplacs.GruplacDefinition, len(list)-1)
	for _, g := range list {
		result = append(result, g.ToDomain())
	}
	return result
}
