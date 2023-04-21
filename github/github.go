package github

import (
	"context"
	"cuddly-eureka-/appconstants"
	"cuddly-eureka-/util"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
	"strings"
)

type GitHubActions interface {
	GetAuthenticatedUserRepos() ([]string, error)
	GetOrgRepos(org string) ([]string, error)
	GetRepoLanguages(repoName []string, owner string) []RepoLanguageDetails
	GetSingleRepoLanguages(repoName, owner string) chan RepoLanguageResponse
	GetRepoContent(repoName, owner, filename string) (*string, error)
}

type RepositoryActions interface {
	GetPackageJSON(repoName, owner string) (util.PackageJson, error)
	GetDotNVMRC(repoName, owner string) (*string, error)
	GetRequirementsTxt(repoName, owner string) (*string, error)
	DoesReadMeExist(repoName, owner string) (*bool, error)
	GetDockerFile(repoName, owner string) (*string, error)
}

type RepoLanguageDetails struct {
	Name      string
	Languages []string
	ECR       string
	AWS       appconstants.RepoAWSDetails
}

type RepoLanguageResponse struct {
	Error   error
	Details RepoLanguageDetails
}

type GitHub struct {
	ctx    context.Context
	client *github.Client
}

func (g *GitHub) Init(token string) *GitHub {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return &GitHub{
		client: client,
		ctx:    ctx,
	}
}

// GetAuthenticatedUserRepos returns the repo names of the user for which
// the token was supplied
func (g *GitHub) GetAuthenticatedUserRepos() ([]string, error) {
	repos, _, err := g.client.Repositories.List(g.ctx, "", nil)
	repoNames := make([]string, 0)
	if err != nil {
		return nil, err
	} else {
		for _, r := range repos {
			repoNames = append(repoNames, *r.Name)
		}
	}

	return repoNames, nil
}

// GetRepoLanguages returns a slice that contains a structure holding the repo language
// details
func (g *GitHub) GetRepoLanguages(repoNames []string, owner string) []RepoLanguageDetails {
	languages := make([]RepoLanguageDetails, 0)
	repoLanguagesChan := make([]chan RepoLanguageResponse, 0)
	for _, name := range repoNames {
		// Concurrent calls to fetch the repo languages
		repoLanguagesChan = append(repoLanguagesChan, g.GetSingleRepoLanguages(name, owner))
	}

	for _, ch := range repoLanguagesChan {
		repoLangResponse := <-ch
		if repoLangResponse.Error != nil {
			// TODO: Handle this better
			fmt.Println("Failed to fetch language details for " + repoLangResponse.Details.Name)
		} else {
			languages = append(languages, RepoLanguageDetails{
				Name:      repoLangResponse.Details.Name,
				Languages: repoLangResponse.Details.Languages,
			})
		}
	}

	return languages
}

func (g *GitHub) GetSingleRepoLanguages(repoName, owner string) chan RepoLanguageResponse {
	langDetailsChannel := make(chan RepoLanguageResponse)
	go func() {
		langMap, _, err := g.client.Repositories.ListLanguages(g.ctx, owner, repoName)
		langDetails := new(RepoLanguageResponse)
		if err != nil {
			langDetails.Error = err
			langDetailsChannel <- *langDetails
			langDetails.Details.Name = repoName
		} else {
			langDetails.Details = RepoLanguageDetails{
				Name:      repoName,
				Languages: util.GetAllKeys(langMap),
			}
			langDetailsChannel <- *langDetails
		}
	}()
	return langDetailsChannel
}

func (g *GitHub) GetRepoContent(repoName, owner, filename string) (*string, error) {
	c, _, _, err := g.client.Repositories.GetContents(g.ctx, owner, repoName, filename, nil)
	if err != nil {
		return nil, err
	} else {
		con, _ := c.GetContent()
		return &con, nil
	}
}

func (g *GitHub) GetPackageJSON(repoName, owner string) (util.PackageJson, error) {
	content, err := g.GetRepoContent(repoName, owner, "package.json")
	if err != nil {
		return nil, err
	} else {
		var jsonMap util.PackageJson
		json.Unmarshal([]byte(*content), &jsonMap)
		return jsonMap, nil
	}
}

// GetDotNVMRC returns the content of the .nvmrc file from the given
// repo name
func (g *GitHub) GetDotNVMRC(repoName, owner string) (*string, error) {
	content, err := g.GetRepoContent(repoName, owner, ".nvmrc")
	if err != nil {
		return nil, err
	}
	return content, nil
}

// GetRequirementsTxt fetches the requirements.txt file from the given repo name
func (g *GitHub) GetRequirementsTxt(repoName, owner string) (*string, error) {
	content, err := g.GetRepoContent(repoName, owner, "requirements.txt")
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (g *GitHub) DoesReadMeExist(repoName, owner string) (*bool, error) {
	content, response, err := g.client.Repositories.GetReadme(g.ctx, owner, repoName, nil)
	exists := false

	if err != nil {
		if response.StatusCode == 404 {
			return &exists, nil
		}
		return nil, err
	}

	readmeContent, readMeErr := content.GetContent()
	if readMeErr != nil {
		return nil, readMeErr
	}
	readmeContent = strings.TrimSpace(readmeContent)

	if len(readmeContent) > 0 {
		exists = true
		return &exists, nil
	}

	return &exists, nil
}

func (g *GitHub) GetDockerFile(repoName, owner string) (*string, error) {
	content, err := g.GetRepoContent(repoName, owner, "Dockerfile")
	if err != nil {
		return nil, err
	}

	return content, nil
}
