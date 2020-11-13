package ports

import "rpcf/gruplacs"

type GruplacDefinitionWriter interface {
	Create(g *gruplacs.GruplacDefinition) (*gruplacs.GruplacDefinition, error)

	Update(g *gruplacs.GruplacDefinition) (*gruplacs.GruplacDefinition, error)

	Delete(id string) error
}
