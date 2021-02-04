package managers

import (
	"fmt"
	"regexp"
	"rpcf/analysis/ports"
	"rpcf/core"
	"strings"
)

func init() {
	err := core.Injector.Provide(newWordFrequencyCounter)
	core.CheckInjection(err, "newWordFrequencyCounter")
}

//Word type struct defined.
type Word struct {
	key   string
	value int
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

func (c *wordFrequencyCounter) Count(text string) map[string]int {
	wordList := strings.Fields(text)
	wordCounts := make(map[string]int)

	for _, word := range wordList {
		cleanWord := c.clean(strings.ToLower(word))
		found := c.checkStopWords(cleanWord)
		if !found {
			wordCounts[cleanWord]++
		}
	}
	return wordCounts
}

func (c *wordFrequencyCounter) clean(text string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		// TODO: return the error and handle it
		fmt.Println(err)
	}
	processedWord := reg.ReplaceAllString(text, "")
	return processedWord
}

func (c *wordFrequencyCounter) checkStopWords(word string) bool {
	for _, item := range c.stopWords {
		if item == word {
			return true
		}
	}
	return false
}
