package managers

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"rpcf/analysis/ports"
	"rpcf/core"
)

const (
	stopWordsDirectory = "analysis/stop_words/"
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
	words := make([]string, 0)

	files, err := ioutil.ReadDir(stopWordsDirectory)
	if err != nil {
		fmt.Println(err)
		return words, err
	}
	for _, f := range files {
		filePath := fmt.Sprintf("%s%s", stopWordsDirectory, f.Name())
		lines, err := s.loadFile(filePath)
		if err != nil {
			fmt.Println(err)
			continue
		}
		words = append(words, lines...)
	}

	return words, nil
}

func (s *StopWordsManager) loadFile(filePath string) ([]string, error) {
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
