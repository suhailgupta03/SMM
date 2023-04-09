package initialize

import (
	"cuddly-eureka-/appconstants"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

func config() (*appconstants.MaturityYAMLStruct, error) {
	fileName := os.Getenv("MATURITY_REPO_YAML")
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Warning!!: Failed to read maturity repo yaml ..")
		fmt.Println("Warning!! File named '" + fileName + "' must be present in the root .. Check execution context")
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
	scanAllGitHub := false
	if strings.ToLower(strings.TrimSpace(os.Getenv("SCAN_ALL_GITHUB_REPOS"))) == "true" {
		scanAllGitHub = true
	}

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
		Stage:         stage,
		GitHubToken:   os.Getenv("TOKEN"),
		GitHubOwner:   os.Getenv("OWNER"),
		ScanAllGitHub: scanAllGitHub,
		Test:          test,
	}
}

// GetAppConstants returns the application constants or the configuration
// variables. If it fails to find conf.toml, it checks the ENV to find
// the matching variables
func GetAppConstants() appconstants.Constants {
	c := getFromEnv()
	d, _ := config()
	if d != nil {
		c.MaturityRepoDetails = &d.Repository
	}

	return c
}
