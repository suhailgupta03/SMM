package main

import (
	"cuddly-eureka-/http"
	"cuddly-eureka-/types"
	"fmt"
)

type ExtendedProductVersion types.ProductVersion

func (pv *ExtendedProductVersion) IsNodeEOL(version string) bool {
	details := http.GetEOLDetails("node")
	fmt.Println(details.EOLDate, " for node")
	// add logic to compute EOL
	return false
}

func (pv *ExtendedProductVersion) IsPythonEOL(version string) bool {
	details := http.GetEOLDetails("python")
	fmt.Println(details.EOLDate, " for python")
	// add logic to compute EOL
	return true
}
