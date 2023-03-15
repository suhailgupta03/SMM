package main

import (
	"cuddly-eureka-/types"
	"fmt"
	"os"
	"path/filepath"
	"plugin"
)

func getRepos() []string {
	repos := make([]string, 0)
	repos = append(repos, "repo1")
	repos = append(repos, "repo2")
	return repos
}

func main() {
	entries, err := os.ReadDir("plugins")
	if err != nil {
		panic(err)
	}

	repos := getRepos()
	for _, repo := range repos {
		for _, e := range entries {
			/**
			Load all the plugins and run for each repo
			*/
			pluginDirName := e.Name()
			pluginFileName := pluginDirName + ".so"
			pluginFilePath := filepath.Join("plugins", pluginDirName, pluginFileName)
			plug, plugErr := plugin.Open(pluginFilePath)
			if plugErr != nil {
				panic(plugErr)
			}
			symbol, sErr := plug.Lookup("Check")
			if sErr != nil {
				panic(sErr)
			}
			var maturity types.Maturity
			maturity, ok := symbol.(types.Maturity)
			if !ok {
				panic("Type mismatch")
			}

			/**
			Now we extract the ProductVersion and cast that into the
			new type which is distinct but derives from the ProductVersion
			*/
			maturityValue := maturity.Check(repo)
			maturityCheck := ExtendedMaturityCheck(maturityValue)
			// The following methods IsNodeEOL, IsNodeEOL, etc,. are controlled
			// by runner.go and developers creating a plugin need not
			// implement these methods. Developers creating plugin
			// must only return the version numbers in the GetVersion
			// response
			maturityCheck.Print(repo, pluginDirName, maturityValue)
			fmt.Println("========")
		}

	}
}
