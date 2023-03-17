package util

import (
	"cuddly-eureka-/http"
	"reflect"
	"time"
)

func FindMatchingVersion(versionToFind string, eolList []http.ProductEOLDetails) http.ProductEOLDetails {
	details := new(http.ProductEOLDetails)
	for _, d := range eolList {
		if d.Cycle == versionToFind {
			details = &d
			break
		}
	}

	return *details
}

func IsVersionEOL(versionToCheck string, eolDetails http.ProductEOLDetails) bool {
	if eolDetails.Cycle == versionToCheck {
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
		panic("Version to check does not match with the provided EOL details")
	}
}
