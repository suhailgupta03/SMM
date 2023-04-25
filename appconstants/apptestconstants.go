package appconstants

type TestConstants struct {
	Repo    TestRepos
	AWS     TestAWS
	CodeCov TestCodeCov
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

type TestCodeCov struct {
	RepoName    string
	RepoOwner   string
	BearerToken string
}
