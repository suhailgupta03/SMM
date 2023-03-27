package util

import (
	"cuddly-eureka-/conf/initialize"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoesTrivyExist(t *testing.T) {
	exist, err := doesTrivyExist()
	assert.True(t, exist)
	assert.Nil(t, err)
}

func TestParseCriticalVul(t *testing.T) {
	trivyResponse1 := `
			requirements.txt (pip)
			
			Total: 2 (CRITICAL: 2)
			
			┌──────────────┬────────────────┬──────────┬───────────────────┬───────────────┬──────────────────────────────────────────────────────────────┐
			│   Library    │ Vulnerability  │ Severity │ Installed Version │ Fixed Version │                            Title                             │
			├──────────────┼────────────────┼──────────┼───────────────────┼───────────────┼──────────────────────────────────────────────────────────────┤
			│ pandas       │ CVE-2020-13091 │ CRITICAL │ 0.23.0            │ 1.0.4         │ ** DISPUTED ** pandas through 1.0.3 can unserialize and      │
			│              │                │          │                   │               │ execute comman ......                                        │
			│              │                │          │                   │               │ https://avd.aquasec.com/nvd/cve-2020-13091                   │
			├──────────────┼────────────────┤          ├───────────────────┼───────────────┼──────────────────────────────────────────────────────────────┤
			│ scikit_learn │ CVE-2020-13092 │          │ 0.19.1            │ 0.23.1        │ ** DISPUTED ** scikit-learn (aka sklearn) through 0.23.0 can │
			│              │                │          │                   │               │ unseriali ...                                                │
			│              │                │          │                   │               │ https://avd.aquasec.com/nvd/cve-2020-13092                   │
			└──────────────┴────────────────┴──────────┴───────────────────┴───────────────┴──────────────────────────────────────────────────────────────┘`

	exists, totalVul := parseCriticalVul(trivyResponse1)
	assert.True(t, *exists)
	assert.Equal(t, 2, *totalVul)

	trivyResponse2 := `
			requirements.txt (pip)
			
			Total: 4 (CRITICAL: 4)
			
			┌──────────────┬────────────────┬──────────┬───────────────────┬───────────────────────┬──────────────────────────────────────────────────────────────┐
			│   Library    │ Vulnerability  │ Severity │ Installed Version │     Fixed Version     │                            Title                             │
			├──────────────┼────────────────┼──────────┼───────────────────┼───────────────────────┼──────────────────────────────────────────────────────────────┤
			│ Django       │ CVE-2022-41323 │ HIGH     │ 3.2.15            │ 3.2.16, 4.0.8, 4.1.2  │ python-django: Potential denial-of-service vulnerability in  │
			│              │                │          │                   │                       │ internationalized URLs                                       │
			│              │                │          │                   │                       │ https://avd.aquasec.com/nvd/cve-2022-41323                   │
			│              ├────────────────┤          │                   ├───────────────────────┼──────────────────────────────────────────────────────────────┤
			│              │ CVE-2023-23969 │          │                   │ 4.1.6, 4.0.9, 3.2.17  │ python-django: Potential denial-of-service via               │
			│              │                │          │                   │                       │ Accept-Language headers                                      │
			│              │                │          │                   │                       │ https://avd.aquasec.com/nvd/cve-2023-23969                   │
			│              ├────────────────┤          │                   ├───────────────────────┼──────────────────────────────────────────────────────────────┤
			│              │ CVE-2023-24580 │          │                   │ 4.0.10, 4.1.7, 3.2.18 │ python-django: Potential denial-of-service vulnerability in  │
			│              │                │          │                   │                       │ file uploads                                                 │
			│              │                │          │                   │                       │ https://avd.aquasec.com/nvd/cve-2023-24580                   │
			├──────────────┼────────────────┤          ├───────────────────┼───────────────────────┼──────────────────────────────────────────────────────────────┤
			│ scikit_learn │ CVE-2020-28975 │          │ 0.19.1            │ 0.24.dev0             │ ** DISPUTED ** svm_predict_values in svm.cpp in Libsvm v324, │
			│              │                │          │                   │                       │ as used in...                                                │
			│              │                │          │                   │                       │ https://avd.aquasec.com/nvd/cve-2020-28975                   │
			└──────────────┴────────────────┴──────────┴───────────────────┴───────────────────────┴──────────────────────────────────────────────────────────────┘
			
			yarn.lock (yarn)
			
			Total: 2 (CRITICAL: 2)
			
			┌─────────┬────────────────┬──────────┬───────────────────┬─────────────────────────────────────────────────────────┬──────────────────────────────────────────────────────────┐
			│ Library │ Vulnerability  │ Severity │ Installed Version │                      Fixed Version                      │                          Title                           │
			├─────────┼────────────────┼──────────┼───────────────────┼─────────────────────────────────────────────────────────┼──────────────────────────────────────────────────────────┤
			│ qs      │ CVE-2022-24999 │ HIGH     │ 6.5.1             │ 6.2.4, 6.3.3, 6.4.1, 6.5.3, 6.6.1, 6.7.3, 6.8.3, 6.9.7, │ express: "qs" prototype poisoning causes the hang of the │
			│         │                │          │                   │ 6.10.3                                                  │ node process                                             │
			│         │                │          │                   │                                                         │ https://avd.aquasec.com/nvd/cve-2022-24999               │
			│         │                │          ├───────────────────┤                                                         │                                                          │
			│         │                │          │ 6.5.2             │                                                         │                                                          │
			│         │                │          │                   │                                                         │                                                          │
			│         │                │          │                   │                                                         │                                                          │
			└─────────┴────────────────┴──────────┴───────────────────┴─────────────────────────────────────────────────────────┴──────────────────────────────────────────────────────────┘`

	exists, totalVul = parseCriticalVul(trivyResponse2)
	assert.True(t, *exists)
	// Note: currently it will report the first encounter
	assert.Equal(t, 4, *totalVul)

	trivyResponse3 := `
			requirements.txt (pip)
			
			Total: 2 (HIGH: 2)
			
			┌──────────────┬────────────────┬──────────┬───────────────────┬───────────────┬──────────────────────────────────────────────────────────────┐
			│   Library    │ Vulnerability  │ Severity │ Installed Version │ Fixed Version │                            Title                             │
			├──────────────┼────────────────┼──────────┼───────────────────┼───────────────┼──────────────────────────────────────────────────────────────┤
			│ pandas       │ CVE-2020-13091 │ CRITICAL │ 0.23.0            │ 1.0.4         │ ** DISPUTED ** pandas through 1.0.3 can unserialize and      │
			│              │                │          │                   │               │ execute comman ......                                        │
			│              │                │          │                   │               │ https://avd.aquasec.com/nvd/cve-2020-13091                   │
			├──────────────┼────────────────┤          ├───────────────────┼───────────────┼──────────────────────────────────────────────────────────────┤
			│ scikit_learn │ CVE-2020-13092 │          │ 0.19.1            │ 0.23.1        │ ** DISPUTED ** scikit-learn (aka sklearn) through 0.23.0 can │
			│              │                │          │                   │               │ unseriali ...                                                │
			│              │                │          │                   │               │ https://avd.aquasec.com/nvd/cve-2020-13092                   │
			└──────────────┴────────────────┴──────────┴───────────────────┴───────────────┴──────────────────────────────────────────────────────────────┘`

	exists, totalVul = parseCriticalVul(trivyResponse3)
	assert.False(t, *exists)
	assert.Nil(t, totalVul)

	trivyResponse4 := ""
	exists, totalVul = parseCriticalVul(trivyResponse4)
	assert.False(t, *exists)
	assert.Nil(t, totalVul)

	trivyResponse5 := `2023-03-27T20:09:55.071+0530	FATAL	repository scan error: scan error: unable to initialize a scanner: unable to initialize a repository scanner: git clone error: authentication required`
	exists, totalVul = parseCriticalVul(trivyResponse5)
	assert.False(t, *exists)
	assert.Nil(t, totalVul)

}

func TestIsRepoVulnerable(t *testing.T) {
	app := initialize.GetAppConstants()
	repoPath := GenerateUrlFromRepoName(app.Test.Repo.Trivy)
	isVul, err := IsRepoVulnerable(repoPath)
	assert.True(t, *isVul)
	assert.Nil(t, err)

	randomRepoPath := GenerateUrlFromRepoName("uvRmcTJbpUUjYLZszfXf")
	isVul, err = IsRepoVulnerable(randomRepoPath)
	assert.Nil(t, isVul)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "exit status 1")
}
