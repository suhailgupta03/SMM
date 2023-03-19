package util

import (
	"regexp"
	"strings"
)

// GetVersionFromRequirementsTxt returns the version of the package from the requirements.txt
// file string passed filedata.
func GetVersionFromRequirementsTxt(fileData, packageName string) *string {
	// replacing \r\n with \n just in case the file was prepared in windows os
	attributeList := strings.Split(strings.ReplaceAll(fileData, "\r\n", "\n"), "\n")
	var version *string
	for _, attrib := range attributeList {
		attribDetails := strings.Split(strings.TrimSpace(attrib), "==")
		if len(attribDetails) == 2 && strings.ToLower(attribDetails[0]) == strings.ToLower(packageName) {
			version = &attribDetails[1]
			break
		}
	}

	return version
}

// GetVersionFromPackageJSON takes the package.json and returns the version of the package name
// passed as an argument. Returns nil if the package name does not exist
func GetVersionFromPackageJSON(fileData PackageJson, packageName string) *string {
	dep, fDep := fileData["dependencies"]
	if fDep {
		version, fVer := dep.(map[string]interface{})[packageName]
		if fVer {
			reg := regexp.MustCompile(`^\^?`)
			version = reg.ReplaceAllString(version.(string), "")
			v := version.(string)
			return &v
		}
	}

	return nil
}

type PackageJson map[string]interface{}
