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
	Name string `yaml:"name"`
	ECR  string `yaml:"ecr"`
}

type MaturityYAMLStruct struct {
	Name       string               `yaml:"name"`
	Repository []MaturityRepoDetail `yaml:"repository"`
}
