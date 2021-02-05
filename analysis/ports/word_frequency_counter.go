package ports

type WordFrequencyCounter interface {
	Count(text string) map[string]int
}
