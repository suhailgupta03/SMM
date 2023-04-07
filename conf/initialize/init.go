package initialize

import (
	"cuddly-eureka-/appconstants"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

func config() (*appconstants.MaturityYAMLStruct, error) {
	fileName := "repo-details.yml" // os.Getenv("MATURITY_REPO_YAML")
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: Failed to read maturity repo yaml ..")
		fmt.Println("File named '" + fileName + "' must be present in the root")
		return nil, err
	}

	maturityYAML := new(appconstants.MaturityYAMLStruct)
	if err = yaml.Unmarshal(data, maturityYAML); err != nil {
		return nil, err
	}

	return maturityYAML, nil
}

func isStageTest(stage string) bool {
	if strings.ToLower(strings.TrimSpace(stage)) == "test" {
		return true
	} else {
		return false
	}
}

func getFromEnv() appconstants.Constants {
	test := new(appconstants.TestConstants)
	stage := ""
	if isStageTest(os.Getenv("STAGE")) {
		test.Repo.Node = os.Getenv("NODE")
		test.Repo.Empty = os.Getenv("EMPTY")
		test.Repo.NVMRC = os.Getenv("NVMRC_ONLY")
		test.Repo.Django = os.Getenv("DJANGO")
		test.Repo.Trivy = os.Getenv("TRIVYREPO")
		stage = "test"
	} else {
		test = nil
	}
	return appconstants.Constants{
		Stage:       stage,
		GitHubToken: os.Getenv("TOKEN"),
		GitHubOwner: os.Getenv("OWNER"),
		Test:        test,
	}
}

// GetAppConstants returns the application constants or the configuration
// variables. If it fails to find conf.toml, it checks the ENV to find
// the matching variables
func GetAppConstants() appconstants.Constants {
	c := getFromEnv()
	if !isStageTest(os.Getenv("STAGE")) {
		d, _ := config()
		c.MaturityRepoDetails = d.Repository
	}

	return c
}
