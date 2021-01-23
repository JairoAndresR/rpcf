package entities

const (
	tableNameGrupLAC = "gruplacs"
)

type GrupLAC struct {
	ID      string `gorm:"primaryKey"`
	Name    string
	Code    string
	Authors []Author `gorm:"many2many:authors_gruplacs;"`
}

func (g *GrupLAC) TableName() string {
	return tableNameGrupLAC
}
