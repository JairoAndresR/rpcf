package products

import "time"

type Product struct {
	ID              string
	SKResearcher    string
	SKResearchGroup string
	TypeId          string
	TypeName        string
	Title           string
	StartYear       string
	EndYear         string
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}
