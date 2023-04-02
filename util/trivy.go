package util

import (
	"errors"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func doesTrivyExist() (bool, error) {
	cmd := exec.Command("trivy", "version")
	var out strings.Builder
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return false, err
	}
	if strings.Contains(strings.ToLower(out.String()), "version") {
		return true, nil
	}

	return false, nil
}

func parseCriticalVul(trivyResponse string) (*bool, *int) {
	exists := false
	var numberOf *int

	if len(trivyResponse) == 0 {
		return &exists, nil
	}

	trivyOutputRegex := regexp.MustCompile(`(?i)(Total:\s*\d\s*)(\(CRITICAL:\s*\d\s*\))`)
	if trivyOutputRegex.MatchString(trivyResponse) {
		matches := trivyOutputRegex.FindStringSubmatch(trivyResponse)
		if len(matches) == 3 {
			totalVul := strings.TrimSpace(strings.Split(matches[1], ":")[1])
			tVul, _ := strconv.Atoi(totalVul)
			if tVul > 0 {
				exists = true
				numberOf = &tVul
			}
		}
	}

	return &exists, numberOf
}

func IsRepoVulnerable(repoPath string) (*bool, error) {
	exists, eErr := doesTrivyExist()
	if eErr != nil {
		return nil, eErr
	}

	isVulnerable := false

	if exists {
		cmd := exec.Command("trivy", "repository", repoPath, "-q", "-s", "CRITICAL", "-f", "table")
		var out strings.Builder
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			return nil, err
		}
		vulExists, _ := parseCriticalVul(out.String())
		if *vulExists {
			isVulnerable = true
		}

		return &isVulnerable, nil
	} else {
		return nil, errors.New("trivy does not exist")
	}
}

func IsImageVulnerable(imagePath string) (*bool, error) {
	exists, eErr := doesTrivyExist()
	if eErr != nil {
		return nil, eErr
	}

	isVulnerable := false
	if exists {
		cmd := exec.Command("trivy", "image", imagePath, "-q", "-s", "CRITICAL", "-f", "table")
		var out strings.Builder
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			return nil, err
		}
		vulExists, _ := parseCriticalVul(out.String())
		if *vulExists {
			isVulnerable = true
		}
		return &isVulnerable, nil
	} else {
		return nil, errors.New("trivy does not exist")
	}

}
