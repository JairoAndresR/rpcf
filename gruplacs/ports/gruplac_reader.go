package ports

import "rpcf/gruplacs"

type GrupLACReader interface {

	// GetAll retrieves all the grupLACs to collect the information.
	GetAll() []gruplacs.GrupLAC
}
