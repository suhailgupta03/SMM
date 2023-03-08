package main

import "cuddly-eureka-/types"

type Integrations struct {
}

func (analytics Integrations) GetVersion() types.ProductVersion {
	var version types.ProductVersion
	version.Python = "9.1.2"
	version.React = "4.5.1"
	version.Node = "14.0.0"
	version.Django = "1.2.9"
	return version
}

var Version Integrations
