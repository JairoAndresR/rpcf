package gruplacs

import "encoding/json"

type CollectorPayload struct {
	GrupLACCode string `json:"code"`
	Content     string `json:"content"`
}

func NewCollectorPayload(content string, code string) *CollectorPayload {
	return &CollectorPayload{
		GrupLACCode: code,
		Content:     content,
	}
}

func (p *CollectorPayload) JSONString() (string, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
