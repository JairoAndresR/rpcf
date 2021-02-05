package entities

import "rpcf/products"

const (
	tableNameGrupLAC = "gruplacs"
)

type GrupLAC struct {
	Name    string
	Code    string   `gorm:"primaryKey"`
	Authors []Author `gorm:"many2many:authors_gruplacs;"`
}

func NewGrupLAC(g *products.GrupLAC) *GrupLAC {
	return &GrupLAC{
		Code: g.Code,
		Name: g.Name,
	}
}

func (g *GrupLAC) TableName() string {
	return tableNameGrupLAC
}

func (g *GrupLAC) ToDomain() *products.GrupLAC {
	return products.NewGrupLAC(g.Code, g.Name)
}

func MapGrupLACsToLisDomain(list []*GrupLAC) []*products.GrupLAC {
	results := make([]*products.GrupLAC, 0)

	for _, g := range list {
		results = append(results, g.ToDomain())
	}
	return results
}
