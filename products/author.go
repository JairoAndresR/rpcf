package products

import "encoding/json"

type Author struct {
	ID       string
	Names    string
	GrupLACs []*GrupLAC
}

func NewAuthorForResult(result map[string]string) (*Author, error) {
	var p Author

	jsonBody, err := json.Marshal(result)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBody, &p)

	if err != nil {
		return nil, err
	}
	return &p, nil
}

func NewAuthorsFromResults(results []map[string]string) ([]*Author, []error) {
	as := make([]*Author, 0)
	es := make([]error, 0)

	for _, r := range results {
		a, err := NewAuthorForResult(r)

		if err != nil {
			es = append(es, err)
			continue
		}
		as = append(as, a)
	}

	return as, es
}
