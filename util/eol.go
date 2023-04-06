package util

import (
	"cuddly-eureka-/http"
	"cuddly-eureka-/types"
	"fmt"
	"github.com/hashicorp/go-version"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"
)

func getVersionWithRTrimmedDots(version string) string {
	re := regexp.MustCompile(`\.\d+$`)
	versionWithTrimmedDots := re.ReplaceAllString(version, "")
	return versionWithTrimmedDots
}

/*
*
normalizeVersionString removes :
  - leading characters to make sure that the version always starts with a number
  - training characters to make sure that the version always ends with a number
  - trims any leading or trailing white space in the version string
*/
func normalizeVersionString(version string) string {
	re := regexp.MustCompile(`(?m)(\d+\.?)`)
	vGroups := re.FindAllStringSubmatch(version, -1)
	norm := ""
	if len(vGroups) > 0 {
		for _, group := range vGroups {
			if len(group[1]) > 0 {
				norm += group[1]
			}
		}
	}

	return norm
}

// findMatchingVersion find the details of the product matching the passed version. Currently,
// it matches the following:
// 3.2.14 will be matched with 3.2
// 3.2 will be matched with 3.2
// 3 will be matched with 3
// Note: 3.2.14 will not be matched with 3
func findMatchingVersion(versionToFind string, eolList []http.ProductEOLDetails) http.ProductEOLDetails {
	details := new(http.ProductEOLDetails)
	versionList := make([]*version.Version, 0)

	versionWithTrimmedDots := getVersionWithRTrimmedDots(versionToFind)
	for _, d := range eolList {
		currentVersion, _ := version.NewVersion(d.Cycle)
		versionList = append(versionList, currentVersion)
		if d.Cycle == versionToFind || d.Cycle == versionWithTrimmedDots {
			details = &d
			break
		}
	}

	if len(details.Cycle) == 0 {
		// If the release cycle was not found in the existing
		// eolList, check if the version we are trying to filter is lower than
		// the lowest version reported in the EOLDetails
		sort.Sort(version.Collection(versionList))
		vTFind, _ := version.NewVersion(versionToFind)
		if vTFind.LessThan(versionList[0]) {
			details.Cycle = "-1"
		}
	}
	return *details
}

func isVersionEOL(versionToCheck string, eolDetails http.ProductEOLDetails) bool {
	versionWithTrimmedDots := getVersionWithRTrimmedDots(versionToCheck)
	if eolDetails.Cycle == versionToCheck || versionWithTrimmedDots == eolDetails.Cycle {
		typeofEOL := reflect.TypeOf(eolDetails.EOL).String()
		if typeofEOL == "bool" {
			return eolDetails.EOL.(bool)
		} else {
			eolTime, tErr := time.Parse(time.DateOnly, eolDetails.EOL.(string))
			if tErr != nil {
				panic("Failed to parse the data " + eolDetails.EOL.(string))
			}
			currentDateGreater := IsDateGreater(time.Now(), eolTime)
			if currentDateGreater {
				return true
			}
		}
		return false
	} else {
		fmt.Println(eolDetails.Cycle, "..", versionToCheck, "..", versionWithTrimmedDots)
		panic("Version to check does not match with the provided EOL details")
	}
}

func isOnLatestVersionPatch(versionToCheck string, eolDetails http.ProductEOLDetails) bool {
	if versionToCheck != eolDetails.Latest {
		return false
	}
	return true
}

func CheckEOL(versionToFind string, eolList []http.ProductEOLDetails) types.MaturityCheck {
	versionToFind = normalizeVersionString(versionToFind)
	matchingVersionDetails := findMatchingVersion(versionToFind, eolList)
	if matchingVersionDetails.Cycle != "-1" {
		if isVersionEOL(versionToFind, matchingVersionDetails) {
			return types.MaturityValue1
		} else {
			return types.MaturityValue2
		}
	}

	return types.MaturityValue1
}

func IsUsingLatestPatchVersion(versionToCheck string, eolList []http.ProductEOLDetails) types.MaturityCheck {
	if len(strings.TrimSpace(versionToCheck)) == 0 {
		return types.MaturityValue0
	}

	versionToCheck = normalizeVersionString(versionToCheck)
	matchingVersionDetails := findMatchingVersion(versionToCheck, eolList)
	if matchingVersionDetails.Cycle == "-1" {
		/**
		If the version to find does not exist with the EOLList, it most probably means that the version is too
		old or the version does not exist. In that case, we'll simply report the software is not
		using the latest patch version
		*/
		return types.MaturityValue1
	} else {
		if isOnLatestVersionPatch(versionToCheck, matchingVersionDetails) {
			return types.MaturityValue2
		} else {
			return types.MaturityValue1
		}
	}
}
