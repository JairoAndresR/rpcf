package entities

import (
	"crypto/sha1"
	"gorm.io/gorm"
	"encoding/hex"
	"rpcf/core/entities"
	"rpcf/products"
	"strings"
)

type Author struct {
	*entities.Base
	ID       string `gorm:"primaryKey"`
	Names    string
	GrupLACS []*GrupLAC `gorm:"many2many:authors_gruplacs;"`
	Products []*Product `gorm:"many2many:authors_products;"`
}

func NewAuthor(a *products.Author) *Author {

	gruplacs := make([]*GrupLAC, 0)

	for _, g := range a.GrupLACs {
		gruplacs = append(gruplacs, NewGrupLAC(g))
	}
	return &Author{
		ID:       a.ID,
		Names:    a.Names,
		GrupLACS: gruplacs,
	}
}
func (a *Author) BeforeCreate(tx *gorm.DB) error {
	h := sha1.New()
	h.Write([]byte(strings.ToUpper(a.Names)))
	hb := hex.EncodeToString(h.Sum(nil))
	id := string(hb[:])
	a.ID = id
	a.Names = strings.ToUpper(a.Names)
	return nil
}

func (a *Author) ToDomain() *products.Author {
	gruplacs := make([]*products.GrupLAC, 0)

	for _, g := range a.GrupLACS {
		gruplacs = append(gruplacs, g.ToDomain())
	}

	return &products.Author{
		ID:       a.ID,
		Names:    a.Names,
		GrupLACs: gruplacs,
	}
}

func MapAuthorsToDomain(list []*Author) []*products.Author {

	authors := make([]*products.Author, 0)

	for _, a := range list {
		authors = append(authors, a.ToDomain())
	}
	return authors
}

func NewAuthors(authors []*products.Author) []*Author {

	list := make([]*Author, 0)
	for _, a := range authors {
		list = append(list, NewAuthor(a))
	}
	return list
}
