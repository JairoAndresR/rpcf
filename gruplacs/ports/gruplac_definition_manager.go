package ports

import "rpcf/gruplacs"

type GruplacDefinitionManager interface {
	Create(g *gruplacs.GrupLAC) (*gruplacs.GrupLAC, error)

	Update(g *gruplacs.GrupLAC) (*gruplacs.GrupLAC, error)

	Delete(id string) error

	GetAll()([]*gruplacs.GrupLAC, error)

	GetByName(name string) (*gruplacs.GrupLAC, error)
}