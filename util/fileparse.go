package util

import (
	"strings"
)

// GetVersionFromRequirementsTxt returns the version of the package from the requirements.txt
// file string passed filedata.
func GetVersionFromRequirementsTxt(filedata string, packageName string) *string {
	// replacing \r\n with \n just in case the file was prepared in windows os
	attributeList := strings.Split(strings.ReplaceAll(filedata, "\r\n", "\n"), "\n")
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
