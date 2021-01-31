package entities

import (
	"rpcf/core/entities"
	"rpcf/products"
	"time"
)

type AuthorDefinition struct {
	*entities.Base
	ID         string `gorm:primaryKey"`
	Definition string `gorm:"type:text;not null"`
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

func NewAuthorDefinition(d *products.AuthorDefinition) *AuthorDefinition {
	return &AuthorDefinition{
		ID:         d.ID,
		Definition: d.Definition,
		CreatedAt:  d.CreatedAt,
		UpdatedAt:  d.UpdatedAt,
	}
}

func (a *AuthorDefinition) GetDomainReference() *products.AuthorDefinition {
	return &products.AuthorDefinition{
		ID:         a.ID,
		Definition: a.Definition,
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt,
	}
}
