package ports

import "rpcf/gruplacs"

type GruplacDefinitionReader interface {

	// GetAll retrieves all the grupLACs to collect the information.
	GetAll() ([]*gruplacs.GruplacDefinition, error)

	GetById(id string) (*gruplacs.GruplacDefinition, error)

	GetByName(name string) (*gruplacs.GruplacDefinition, error)
}
