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
	repoLanguageDetails := make([]github.RepoLanguageDetails, 0)
	if appConstants.ScanAllGitHub {
		fmt.Println("Info: Configuration to scan all github set to TRUE .. Will scan all repos for " + owner)
		g := &github.GitHub{}
		g = g.Init(token)
		repoNames, rErr := g.GetAuthenticatedUserRepos()
		if rErr != nil {
			panic("Failed to fetch the repo names. Check token." + rErr.Error())
		}
		fmt.Printf("Fetched %d repos for %s\n", len(repoNames), owner)
		repoLanguageDetails = g.GetRepoLanguages(repoNames, owner)
	} else {
		fmt.Println("Info: Configuration to scan all github set to FALSE .. Will not scan github for " + owner)
		if appConstants.MaturityRepoDetails != nil {
			for _, b := range *appConstants.MaturityRepoDetails {
				d := github.RepoLanguageDetails{
					Name:      b.Name,
					Languages: []string{},
					ECR:       b.ECR,
					AWS:       b.AWS,
				}
				repoLanguageDetails = append(repoLanguageDetails, d)
			}
		}
	}
	return repoLanguageDetails
}

func write(data [][]string, repoName string) {
	headers := make([]string, 0)
	headers = append(headers, "Type", "Name", "Maturity Value")
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
			maturityMeta := maturity.Meta()
			var maturityValue types.MaturityCheck
			var maturityInput string
			if maturityMeta.EcrType {
				if len(repo.ECR) > 0 {
					maturityInput = repo.ECR
				} else {
					fmt.Println("Warning: ECR value not supplied for " + repo.Name)
					maturityInput = ""
				}
			} else if maturityMeta.Type == types.MaturityObservability {
				maturityInput = repo.AWS.LogGroup + "_" + repo.AWS.LogStream
			} else {
				maturityInput = repo.Name
			}
			maturityValue = maturity.Check(maturityInput)
			pluginResult = append(pluginResult, maturityMeta.Type, maturityMeta.Name, strconv.Itoa(int(maturityValue)))
			repoMaturityValues = append(repoMaturityValues, pluginResult)
		}
		write(repoMaturityValues, repo.Name)
	}
}
