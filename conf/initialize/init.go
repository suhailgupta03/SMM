package initialize

import (
	"cuddly-eureka-/appconstants"
	"cuddly-eureka-/conf/toml"
	"os"
)

func Config() (toml.RawConfig, error) {
	configParser := toml.Parser()
	data, err := os.ReadFile("conf.toml")
	if err != nil {
		panic(err)
	}
	c, cErr := configParser.Unmarshal(data)
	if cErr != nil {
		panic(cErr)
	}

	return c, cErr
}

func Constants(config toml.RawConfig) appconstants.Constants {
	return appconstants.Constants{
		GitHubToken: config["TOKEN"].(string),
		GitHubOwner: config["OWNER"].(string),
	}
}
