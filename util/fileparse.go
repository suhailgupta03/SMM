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

type PackageJson map[string]interface{}

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

type FromCommand struct {
	Platform *string
	Image    *string
	Tag      *string
	Digest   *string
	As       *string
}

// ParseDockerFileFromCommand takes in the dockerfile as an input and
// parses the 'From' command in the dockerfile. Returns nil if the
// 'From' command is not found. This method currently does not process
// the digest.
func ParseDockerFileFromCommand(dockerFile string) []*FromCommand {
	/**
	FROM [--platform=<platform>] <image> [AS <name>]
	FROM [--platform=<platform>] <image>[:<tag>] [AS <name>]
	FROM [--platform=<platform>] <image>[@<digest>] [AS <name>]
	*/
	dockerFileNewLineSplit := strings.Split(strings.ReplaceAll(dockerFile, "\r\n", "\n"), "\n")
	fromCommandSlice := make([]*FromCommand, 0)
	re := regexp.MustCompile(`^FROM\s+`)
	platformRegex := regexp.MustCompile(`(--platform=)\S+`)
	asRegex := regexp.MustCompile(`(AS\s+)(\S+)`)
	imageRegex := regexp.MustCompile(`(\S+:)(\S+)`)
	imageRegexWithoutTag := regexp.MustCompile(`(^FROM\s+)(\S+)`)
	colonRegex := regexp.MustCompile(`:$`)
	for _, split := range dockerFileNewLineSplit {
		split = strings.TrimSpace(split)
		if re.MatchString(split) {
			from := new(FromCommand)
			if platformRegex.MatchString(split) {
				pGroups := platformRegex.FindStringSubmatch(split)
				if len(pGroups) == 3 {
					pName := pGroups[2]
					from.Platform = &pName
				}
			}

			if asRegex.MatchString(split) {
				aGroups := asRegex.FindStringSubmatch(split)
				if len(aGroups) == 3 {
					aName := aGroups[2]
					from.As = &aName
				}
			}

			if imageRegex.MatchString(split) {
				iGroups := imageRegex.FindStringSubmatch(split)
				if len(iGroups) == 3 {
					iName := colonRegex.ReplaceAllString(iGroups[1], "")
					iTag := iGroups[2]
					from.Image = &iName
					from.Tag = &iTag
				}
			} else if imageRegexWithoutTag.MatchString(split) {
				iGroups := imageRegexWithoutTag.FindStringSubmatch(split)
				if len(iGroups) == 3 {
					from.Image = &iGroups[2]
				}
			}

			fromCommandSlice = append(fromCommandSlice, from)
		}
	}

	return fromCommandSlice

}
