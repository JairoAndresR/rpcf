package ports

import "rpcf/analysis"

type WordFrequencyCounter interface {
	Count(text string) []analysis.Word
}
