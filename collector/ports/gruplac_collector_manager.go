package ports

type GrupLACCollectorManager interface {

	// It is in charge to collect the products related a research group
	CollectAll() error
}
