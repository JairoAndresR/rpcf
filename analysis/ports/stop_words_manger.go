package ports

type StopWordsManager interface {
	GetStopWords() ([]string, error)
}
