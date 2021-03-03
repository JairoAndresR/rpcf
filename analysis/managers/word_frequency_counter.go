package managers

import (
	"fmt"
	"regexp"
	"rpcf/analysis"
	"rpcf/analysis/ports"
	"rpcf/core"
	"sort"
	"strings"
)

func init() {
	err := core.Injector.Provide(newWordFrequencyCounter)
	core.CheckInjection(err, "newWordFrequencyCounter")
}

type wordFrequencyCounter struct {
	stopWords []string
}

func newWordFrequencyCounter(swm ports.StopWordsManager) ports.WordFrequencyCounter {
	stopWords, _ := swm.GetStopWords()
	return &wordFrequencyCounter{
		stopWords: stopWords,
	}
}

func (c *wordFrequencyCounter) Count(text string) []analysis.Word {
	wordList := strings.Fields(text)
	wordCounts := make(map[string]int)

	for _, word := range wordList {
		cleanWord := c.clean(strings.ToLower(word))
		isStopWord := c.isStopWord(cleanWord)
		if !isStopWord {
			wordCounts[cleanWord]++
		}
	}
	return c.sortedWords(wordCounts)
}

func (c *wordFrequencyCounter) clean(text string) string {
	reg, err := regexp.Compile("[^-_/.,\\p{L}0-9 ]+")
	if err != nil {
		// TODO: return the error and handle it
		fmt.Println(err)
	}
	processedWord := reg.ReplaceAllString(text, "")
	return processedWord
}

func (c *wordFrequencyCounter) isStopWord(word string) bool {
	for _, item := range c.stopWords {
		if strings.ToLower(item) == strings.ToLower(word) {
			return true
		}
	}
	return false
}

func (c *wordFrequencyCounter) sortedWords(words map[string]int) []analysis.Word {
	var sorted []analysis.Word
	for k, v := range words {
		sorted = append(sorted, analysis.Word{Key: k, Value: v})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})
	return sorted
}
