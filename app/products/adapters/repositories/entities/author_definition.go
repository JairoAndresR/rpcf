package entities

import (
	"gorm.io/gorm"
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

func (a *AuthorDefinition) GetDomainReference() *products.AuthorDefinition {
	return &products.AuthorDefinition{
		ID:         a.ID,
		Definition: a.Definition,
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt,
	}
}

func MapAuthorDefinitionToDomain(AuthorDefinition *AuthorDefinition) *products.AuthorDefinition {
	result := make([]*products.AuthorDefinition, 0)
	for _, p := range list {
		result = append(result, p.GetDomainReference())
	}
	return result[0]
}
