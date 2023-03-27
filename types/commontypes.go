package types

type MaturityCheck int

// Reference: https://go.dev/ref/spec#Iota
const (
	MaturityValue0 MaturityCheck = iota
	MaturityValue1
	MaturityValue2
)

type Maturity interface {
	Check(repoName string) MaturityCheck
}
