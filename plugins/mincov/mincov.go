package main

import (
	"cuddly-eureka-/types"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type MinCov struct {
}

type CodeCovRepositoryDetail struct {
	Name        string    `json:"name"`
	Private     bool      `json:"private"`
	Updatestamp time.Time `json:"updatestamp"`
	Author      struct {
		Service  string `json:"service"`
		Username string `json:"username"`
		Name     string `json:"name"`
	} `json:"author"`
	Language  string `json:"language"`
	Branch    string `json:"branch"`
	Active    bool   `json:"active"`
	Activated bool   `json:"activated"`
}

// IsCodeCovActivated checks if the repository is connected with CodeCov and returns true
// if it is. Uses https://docs.codecov.com/reference/repos_retrieve to check for the activation
func isCodeCovActivated(repoName, owner, service, bearerToken string) (*bool, error) {
	apiURL := fmt.Sprintf("https://api.codecov.io/api/v2/%s/%s/repos/%s/", service, owner, repoName)
	authToken := fmt.Sprintf("Bearer %s", bearerToken)
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", authToken)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("CodeCov API did returned a status code of " + strconv.Itoa(resp.StatusCode))
	}

	b, _ := io.ReadAll(resp.Body)
	repoDetails := new(CodeCovRepositoryDetail)
	if err := json.Unmarshal(b, &repoDetails); err != nil {
		return nil, err
	}

	return &repoDetails.Active, nil
}

// Check accepts the _ separated repoName, repoOwner (from the service provider like GitHub), serviceProvider, codecov bearer token
// Reference document for parameters https://docs.codecov.com/reference/repos_retrieve
func (minc *MinCov) Check(repoNameOwnerNameServiceProviderBearerToken string, opts ...*string) types.MaturityCheck {
	argSplit := strings.Split(repoNameOwnerNameServiceProviderBearerToken, "_")
	if len(argSplit) < 4 {
		return types.MaturityValue0
	}

	repoName := argSplit[0]
	repoOwner := argSplit[1]
	serviceProvider := argSplit[2]
	bearerToken := strings.ToLower(argSplit[3])
	activated, _ := isCodeCovActivated(repoName, repoOwner, serviceProvider, bearerToken)
	if activated != nil {
		if *activated == true {
			return types.MaturityValue2
		}
		return types.MaturityValue1
	}

	return types.MaturityValue0
}

func (minc *MinCov) Meta() types.MaturityMeta {
	return types.MaturityMeta{
		Type:        types.MaturityCI,
		Name:        "Unit tests with min coverage enforced",
		CodeCovType: true,
	}
}

var Check MinCov
