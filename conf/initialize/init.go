package initialize

import (
	"cuddly-eureka-/appconstants"
	"cuddly-eureka-/conf/toml"
	"os"
	"strings"
)

func config() (toml.RawConfig, error) {
	configParser := toml.Parser()
	data, err := os.ReadFile("conf.toml")
	if err != nil {
		return nil, err
	}
	c, cErr := configParser.Unmarshal(data)
	if cErr != nil {
		panic(cErr)
	}

	return c, cErr
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
	//conf, err := config()
	//if err != nil {
	//	fmt.Println("Warning: Failed to read the configuration file. " + err.Error())
	//	return getFromEnv()
	//}
	return getFromEnv()
}
