package products

import "time"

type ProductDefinition struct {
	ID         string
	Name       string
	Definition string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
