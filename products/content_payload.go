package products

import "encoding/json"

type contentPayload struct {
	GrupLACCode string `json:"code"`
	Content     string `json:"content"`
}

func NewContentPayload(content string) (*contentPayload, error) {
	var payload contentPayload
	err := json.Unmarshal([]byte(content), &payload)
	return &payload, err
}

func (p *contentPayload) JSONString() (string, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
