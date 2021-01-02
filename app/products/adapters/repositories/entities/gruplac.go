package entities

type GrupLAC struct {
	ID      string `gorm:"primaryKey"`
	Name    string
	Authors []Author `gorm:"many2many:authors_gruplacs;"`
}
