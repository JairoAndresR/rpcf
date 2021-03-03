package products

type Report struct {
	Value string
	Count int64
}

func NewReport(value string, count int64) *Report {
	return &Report{
		Value: value,
		Count: count,
	}
}
