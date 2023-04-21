package appconstants

type TestConstants struct {
	Repo TestRepos
	AWS  TestAWS
}

type TestRepos struct {
	Node   string
	NVMRC  string
	Django string
	Empty  string
	Trivy  string
}

type TestAWS struct {
	LogGroup  string
	LogStream string
}
