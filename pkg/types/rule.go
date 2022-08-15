package types

type Rule struct {
	FromHeader    string
	ToHeader      string
	HeaderMapping map[string]string
}
