package types

type MaturityCheck int

// Reference: https://go.dev/ref/spec#Iota
const (
	MaturityValue0 MaturityCheck = iota
	MaturityValue1
	MaturityValue2
)

type MaturityMeta struct {
	Type        string
	Name        string
	EcrType     bool
	CodeCovType bool
}

const (
	MaturityTypeDependency = "Dependency"
	MaturityTypeDocs       = "Docs"
	MaturityObservability  = "Observability"
	MaturityCI             = "CI"
)

type Maturity interface {
	Check(input string, opts ...*string) MaturityCheck
	Meta() MaturityMeta
}

func MValueToString(mval MaturityCheck) string {
	if mval == 2 {
		return "True"
	} else if mval == 1 {
		return "False"
	} else {
		return "N/A"
	}
}
