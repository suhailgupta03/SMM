package main

import "cuddly-eureka-/types"

type Analytics struct {
}

func (analytics Analytics) GetVersion() types.ProductVersion {
	var version types.ProductVersion
	version.Python = "9.1.2"
	version.Angular = "4.5.1"
	return version
}

var Version Analytics
