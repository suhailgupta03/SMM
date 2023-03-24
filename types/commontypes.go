package types

type MaturityCheck int

// Reference: https://go.dev/ref/spec#Iota
const (
	NA MaturityCheck = iota
	Yes
	No
)

type Maturity interface {
	Check(repoName string) MaturityCheck
}
