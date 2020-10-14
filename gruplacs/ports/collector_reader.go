package ports

type CollectorReader interface {

	// GetContent retrieves the content of specific GrupLAC
	// using the URL of it.
	GetContent(URL string) string
}
