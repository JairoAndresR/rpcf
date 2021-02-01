package ports

type GRUPLACCollectorReader interface {

	// GetContent retrieves the content of specific GrupLAC
	// using the URL of it.
	GetContent(URL string) (string, error)
}
