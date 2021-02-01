package collector

import "encoding/json"

type Payload struct {
	GrupLACCode string `json:"code"`
	GrupLACName string `json:"name"`
	Content     string `json:"content"`
}

func NewCollectorPayload(content string, code, name string) *Payload {
	return &Payload{
		GrupLACCode: code,
		GrupLACName: name,
		Content:     content,
	}
}

func (p *Payload) JSONString() (string, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
