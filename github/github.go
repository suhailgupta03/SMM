package github

import (
	"context"
	"cuddly-eureka-/util"
	"fmt"
	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

type GitHubActions interface {
	GetAuthenticatedUserRepos() ([]string, error)
	GetOrgRepos(org string) ([]string, error)
	GetRepoLanguages(repoName []string, owner string) []RepoLanguageDetails
	GetSingleRepoLanguages(repoName string, owner string) chan RepoLanguageResponse
}

type RepoLanguageDetails struct {
	Name      string
	Languages []string
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

// GetOrgRepos returns the repo names identified by the organization name. Until the token
// supplied has privileges, it will return the public repos by default
func (g *GitHub) GetOrgRepos(org string) ([]string, error) {
	repos, _, err := g.client.Repositories.ListByOrg(g.ctx, org, nil)
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

func (g *GitHub) GetSingleRepoLanguages(repoName string, owner string) chan RepoLanguageResponse {
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
