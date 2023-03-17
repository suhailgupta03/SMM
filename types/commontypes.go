package types

type MaturityCheck int

// Reference: https://go.dev/ref/spec#Iota
const (
	Yes MaturityCheck = iota
	No
	NA
)

type Maturity interface {
	Check(repoName string) MaturityCheck
}
