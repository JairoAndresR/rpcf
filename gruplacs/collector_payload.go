package gruplacs

import "encoding/json"

type CollectorPayload struct {
	GrupLACCode string `json:"code"`
	GrupLACName string `json:"name"`
	Content     string `json:"content"`
}

func NewCollectorPayload(content string, code, name string) *CollectorPayload {
	return &CollectorPayload{
		GrupLACCode: code,
		GrupLACName: name,
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
