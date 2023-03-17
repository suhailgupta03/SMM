package main

import (
	"cuddly-eureka-/appconstants"
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/types"
	"fmt"
	"os"
	"path/filepath"
	"plugin"
)

var (
	appConstants appconstants.Constants
)

func getRepos(token string, owner string) []github.RepoLanguageDetails {
	g := &github.GitHub{}
	g = g.Init(token)
	repoNames, rErr := g.GetAuthenticatedUserRepos()
	if rErr != nil {
		panic("Failed to fetch the repo names. Check token." + rErr.Error())
	}
	fmt.Printf("Fetched %d repos for %s\n", len(repoNames), owner)
	repoLanguageDetails := g.GetRepoLanguages(repoNames, owner)
	return repoLanguageDetails
}

func main() {
	entries, err := os.ReadDir("plugins")
	if err != nil {
		panic(err)
	}
	appConstants = initialize.GetAppConstants()
	repos := getRepos(appConstants.GitHubToken, appConstants.GitHubOwner)

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
			maturityValue := maturity.Check(repo.Name)
			//
			//maturityCheck := ExtendedMaturityCheck(maturityValue)
			//maturityCheck.Print(repo.Name, pluginDirName, maturityValue)
			fmt.Println("========", maturityValue)
		}

	}
}
