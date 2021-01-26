package products

import "time"

type AuthorDefinition struct {
	ID         string
	Definition string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
