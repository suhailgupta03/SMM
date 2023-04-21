package types

type MaturityCheck int

// Reference: https://go.dev/ref/spec#Iota
const (
	MaturityValue0 MaturityCheck = iota
	MaturityValue1
	MaturityValue2
)

type MaturityMeta struct {
	Type    string
	Name    string
	EcrType bool
}

const (
	MaturityTypeDependency = "Dependency"
	MaturityTypeDocs       = "Docs"
	MaturityObservability  = "Observability"
)

type Maturity interface {
	Check(input string) MaturityCheck
	Meta() MaturityMeta
}
