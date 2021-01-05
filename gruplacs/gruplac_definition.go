package gruplacs

import (
	"regexp"
	"strings"
	"time"
)

const (
	codePattern = "(=).*"
)

type GruplacDefinition struct {
	ID        string
	Name      string
	URL       string
	Code      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (g *GruplacDefinition) GetCode() string {
	re := regexp.MustCompile(codePattern)
	code := re.FindString(g.URL)
	return strings.ReplaceAll(code, "=", "")
}
