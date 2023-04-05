package main

import (
	"cuddly-eureka-/appconstants"
	"cuddly-eureka-/conf/initialize"
	"cuddly-eureka-/github"
	"cuddly-eureka-/output/csv"
	"cuddly-eureka-/types"
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"strconv"
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

func write(data [][]string, repoName string) {
	headers := make([]string, 0)
	headers = append(headers, "plugin", "maturity")
	filename := repoName + "_out.csv"
	csv.Generate(headers, data, filename)
	fmt.Printf("Generated %s 📝\n", filename)
}

func main() {
	entries, err := os.ReadDir("plugins")
	if err != nil {
		panic(err)
	}
	appConstants = initialize.GetAppConstants()
	repos := getRepos(appConstants.GitHubToken, appConstants.GitHubOwner)

	for _, repo := range repos {
		repoMaturityValues := make([][]string, 0)
		for _, e := range entries {
			pluginResult := make([]string, 0)
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
			pluginResult = append(pluginResult, pluginDirName, strconv.Itoa(int(maturityValue)))
			repoMaturityValues = append(repoMaturityValues, pluginResult)
		}
		write(repoMaturityValues, repo.Name)
	}
}
