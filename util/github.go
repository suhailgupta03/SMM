package util

import "cuddly-eureka-/conf/initialize"

func GenerateUrlFromRepoName(repoName string) string {
	app := initialize.GetAppConstants()
	return "https://github.com/" + app.GitHubOwner + "/" + repoName
}
