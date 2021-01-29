package products

import (
	"time"
)

type Product struct {
	ID           string `json:"id"`
	GrouplacCode string
	GroupName    string
	TypeId       string
	TypeName     string
	Title        string
	StartYear    string
	EndYear      string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	Authors      []*Author
	GrupLACS     []GrupLAC
}

func (p Product) GetID() string {
	return p.ID
}
