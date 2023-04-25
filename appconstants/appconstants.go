package appconstants

type Constants struct {
	Stage               string
	GitHubToken         string
	GitHubOwner         string
	MaturityRepoDetails *[]MaturityRepoDetail
	ScanAllGitHub       bool
	Test                *TestConstants
}

type MaturityRepoDetail struct {
	Name    string         `yaml:"name"`
	ECR     string         `yaml:"ecr"`
	AWS     RepoAWSDetails `yaml:"aws"`
	CodeCov CodeCov        `yaml:"codecov"`
}

type MaturityYAMLStruct struct {
	Name       string               `yaml:"name"`
	Repository []MaturityRepoDetail `yaml:"repository"`
}

type RepoAWSDetails struct {
	LogStream string `yaml:"log-stream-name"`
	LogGroup  string `yaml:"log-group-name"`
}

type CodeCov struct {
	Bearer string `yaml:"bearer"`
}

type ApplicationFlags struct {
	YAMLFile    *string
	GitHubToken *string
	GitHubOwner *string
}
