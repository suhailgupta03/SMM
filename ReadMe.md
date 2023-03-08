### Generate .so Files and Run Code
To generate the `.so` files and run the code, execute the following script
```shell
./build_run.sh
```


### To Create a New Plugin
- Create a new directory inside `plugins` directory
- Implement `GetVersions` as defined in [types/commontypes.go](./types/commontypes.go)
- Import a variable named `Version` of type `T` that implements `GetVersions`

### Example
To create a plugin named `MyPlugin`

```go
package main

import "cuddly-eureka-/types"

// MyPlugin Create a custom type
type MyPlugin struct {
}

// GetVersion Implement GetVersion on the type created above
func (analytics MyPlugin) GetVersion() types.ProductVersion {
	var version types.ProductVersion
	version.Python = "9.1.2"
	version.React = "4.5.1"
	return version
}

// Version Export a variable of type created above
var Version MyPlugin
```

### Running the plugin
[runner.go](./runner.go) reads the `.so` files in all the plugins directory
and invokes `GetVersion` method. It creates a new type [ExtendedProductVersion](./depchecker.go)
with the same underlying type `ProductVersion` but adds new methods
to check the EOL of the stack.