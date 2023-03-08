package main

import (
	"cuddly-eureka-/types"
	"fmt"
	"os"
	"path/filepath"
	"plugin"
)

func main() {
	entries, err := os.ReadDir("plugins")
	if err != nil {
		panic(err)
	}

	for _, e := range entries {
		pluginDirName := e.Name()
		pluginFileName := pluginDirName + ".so"
		pluginFilePath := filepath.Join("plugins", pluginDirName, pluginFileName)
		fmt.Println(pluginFilePath)
		plug, plugErr := plugin.Open(pluginFilePath)
		if plugErr != nil {
			panic(plugErr)
		}
		symbol, sErr := plug.Lookup("Version")
		if sErr != nil {
			panic(sErr)
		}
		var version types.Version
		version, ok := symbol.(types.Version)
		if !ok {
			panic("Type mismatch")
		}

		/**
		Now we extract the ProductVersion and cast that into the
		new type which is distinct but derives from the ProductVersion
		*/
		productVersions := ExtendedProductVersion(version.GetVersion())
		// The following methods IsNodeEOL, IsNodeEOL, etc,. are controlled
		// by runner.go and developers creating a plugin need not
		// implement these methods.
		nodeResult := productVersions.IsNodeEOL(productVersions.Node)
		pythonResult := productVersions.IsPythonEOL(productVersions.Python)
		// More methods here ..
		fmt.Println("Results for " + pluginDirName)
		fmt.Printf("Node %t Python %t \n", nodeResult, pythonResult)
		fmt.Println("========")
	}

}
