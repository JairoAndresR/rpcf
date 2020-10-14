package gruplacs

import (
	"regexp"
	"strings"
)

const (
	codePattern = "(=).*"
)

type GrupLAC struct {
	Name string
	URL  string
}

func (g *GrupLAC) GetCode() string {
	re := regexp.MustCompile(codePattern)
	code := re.FindString(g.URL)
	return strings.ReplaceAll(code, "=", "")
}
