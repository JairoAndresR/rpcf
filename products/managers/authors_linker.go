package managers

import (
	"rpcf/core"
	"rpcf/products"
	"rpcf/products/ports"
)

func init() {
	err := core.Injector.Provide(newAuthorsLinker)
	core.CheckInjection(err, "newAuthorsLinker")
}

type authorsLinker struct {
	gruplacs map[string]*products.GrupLAC
}

func newAuthorsLinker(gr ports.GrupLACReader) ports.AuthorsLinker {
	// TODO: handler the error
	gs, _ := gr.GetAll()
	gruplacs := make(map[string]*products.GrupLAC, 0)

	for _, g := range gs {
		gruplacs[g.Code] = g
	}
	return &authorsLinker{gruplacs}
}

func (l *authorsLinker) Link(authors []*products.Author, grupLACCode string) []*products.Author {

	g, exist := l.gruplacs[grupLACCode]

	if !exist {
		return authors
	}

	for _, a := range authors {
		a.GrupLACs = append(a.GrupLACs, g)
	}

	return authors
}
