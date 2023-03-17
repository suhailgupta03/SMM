package initialize

import (
	"cuddly-eureka-/appconstants"
	"cuddly-eureka-/conf/toml"
	"cuddly-eureka-/util"
	"fmt"
	"os"
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

// Constants returns configuration structure, prioritizing variables in
// .toml file over .env
func getFromConfig(config toml.RawConfig) appconstants.Constants {
	return appconstants.Constants{
		GitHubToken: util.GetOR(config["TOKEN"].(string), os.Getenv("TOKEN")),
		GitHubOwner: util.GetOR(config["OWNER"].(string), os.Getenv("OWNER")),
	}
}

func getFromEnv() appconstants.Constants {
	return appconstants.Constants{
		GitHubToken: os.Getenv("TOKEN"),
		GitHubOwner: os.Getenv("OWNER"),
	}
}

// GetAppConstants returns the application constants or the configuration
// variables. If it fails to find conf.toml, it checks the ENV to find
// the matching variables
func GetAppConstants() appconstants.Constants {
	conf, err := config()
	if err != nil {
		fmt.Println("Warning: Failed to read the configuration file. " + err.Error())
		return getFromEnv()
	}
	return getFromConfig(conf)
}
