package managers

import (
	"bufio"
	"os"
	"rpcf/analysis/ports"
	"rpcf/core"
)

const (
	filePath = "analysis/stop_words/stop-words-es.txt"
)

func init() {
	err := core.Injector.Provide(newStopWordsManager)
	core.CheckInjection(err, "newStopWordsManager")
}

type StopWordsManager struct {
}

func newStopWordsManager() ports.StopWordsManager {
	return &StopWordsManager{}
}

func (s *StopWordsManager) GetStopWords() ([]string, error) {
	lines := make([]string, 0)
	f, err := os.Open(filePath)
	if err != nil {
		return []string{}, err
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
