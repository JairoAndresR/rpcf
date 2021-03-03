package products

type GrupLAC struct {
	Code string
	Name string
}

func NewGrupLAC(code, name string) *GrupLAC {
	return &GrupLAC{
		Code: code,
		Name: name,
	}
}
