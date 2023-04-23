package initialize

import (
	"cuddly-eureka-/appconstants"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

var (
	repo   *flag.FlagSet
	github *flag.FlagSet
)

func config(yamlToReadFrom *string) (*appconstants.MaturityYAMLStruct, error) {
	fileName := os.Getenv("MATURITY_REPO_YAML")
	if yamlToReadFrom != nil {
		// Override the bash variable with the value passed using the application flag
		fileName = *yamlToReadFrom
	}
	maturityYAML := new(appconstants.MaturityYAMLStruct)
	// If the filename exists, read everything from the YAML file
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Warning!!: Failed to read maturity repo yaml ..")
		fmt.Println("Warning!! File named '" + fileName + "' must be present in the root .. Check execution context")
		return nil, err
	}

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

// getFromEnv returns the constants used across application. The values passed as
// application flags are preferred over the variables exported from the shell
func getFromEnv(flags *appconstants.ApplicationFlags) appconstants.Constants {
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
		test.AWS.LogGroup = os.Getenv("AWS_TEST_LOG_GROUP")
		test.AWS.LogStream = os.Getenv("AWS_TEST_LOG_STREAM")
		stage = "test"
	} else {
		test = nil
	}
	return appconstants.Constants{
		Stage: stage,
		GitHubToken: func() string {
			// TODO: Replace this existing GetOR function in util package
			if flags != nil && len(*(*flags).GitHubToken) > 0 {
				return *flags.GitHubToken
			}
			return os.Getenv("TOKEN")
		}(),
		GitHubOwner: func() string {
			// TODO: Replace this existing GetOR function in util package
			if flags != nil && len(*(*flags).GitHubOwner) > 0 {
				return *flags.GitHubOwner
			}
			return os.Getenv("OWNER")
		}(),
		ScanAllGitHub: scanAllGitHub,
		Test:          test,
	}
}

// initFlags enables the command line flags for the application
func initFlags() appconstants.ApplicationFlags {
	repo = flag.NewFlagSet("repo", flag.ContinueOnError)
	ymlFileName := repo.String("yml", "repo-details.yml", "(optional) input yaml file holding repository details")

	github = flag.NewFlagSet("github", flag.ContinueOnError)
	githubToken := github.String("token", "", "(optional) github token to use to scan a repository")
	githubOwner := github.String("owner", "", "(optional) github account owner / username")

	// Parse the command line flags
	if len(os.Args) > 2 {
		commandParser(os.Args[1:])
	}

	return appconstants.ApplicationFlags{
		YAMLFile:    ymlFileName,
		GitHubOwner: githubOwner,
		GitHubToken: githubToken,
	}
}

// commandParser recursively parses the command line flags
func commandParser(osArgs []string) {
	if len(osArgs) <= 1 {
		return
	}

	switch osArgs[0] {
	case "repo":
		if err := repo.Parse(osArgs[1:]); err != nil {
			log.Fatalf("error loading flags: %v", err)
		}
		osArgs = repo.Args()
		break
	case "github":
		if err := github.Parse(osArgs[1:]); err != nil {
			log.Fatalf("error loading flags: %v", err)
		}
		osArgs = github.Args()
		break
	}
	commandParser(osArgs)

}

// GetAppConstants returns the application constants or the configuration
// variables. If it fails to find conf.toml, it checks the ENV to find
// the matching variables
func GetAppConstants() appconstants.Constants {
	flags := initFlags()
	fmt.Println(flags)
	c := getFromEnv(&flags)
	d, _ := config(flags.YAMLFile)
	if d != nil {
		c.MaturityRepoDetails = &d.Repository
	}

	return c
}
