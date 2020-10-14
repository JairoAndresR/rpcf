package gruplacs

import (
	"regexp"
	"strings"
	"time"
)

const (
	codePattern = "(=).*"
)

type GrupLAC struct {
	ID        string
	Name      string
	URL       string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (g *GrupLAC) GetCode() string {
	re := regexp.MustCompile(codePattern)
	code := re.FindString(g.URL)
	return strings.ReplaceAll(code, "=", "")
}
