[![Run Tests](https://github.com/suhailgupta03/cuddly-eureka-/actions/workflows/test.yml/badge.svg)](https://github.com/suhailgupta03/cuddly-eureka-/actions/workflows/test.yml)

### Generate .so Files and Run Code
To generate the `.so` files and run the code, execute the following script
```shell
./build_run.sh
```


### To Create a New Plugin
- Create a new directory inside `plugins` directory
- Implement `Check` as defined in [types/commontypes.go](./types/commontypes.go)
- Import a variable named `Check` of type `T` that implements `Check` method

### Example
To create a plugin named `HasAutomatedTests`

```go
package main

import "cuddly-eureka-/types"

// HasAutomatedTests creates a custom type
type HasAutomatedTests struct {
}

// Check holds the logic that decides the value of MaturityCheck
func (hat HasAutomatedTests) Check(repoPath string) types.MaturityCheck {
	// Custom Logic Inside the Check Method
	return types.Yes
}

// Check is exported from this plugin file
var Check HasAutomatedTests
```

### Running the plugin
[runner.go](./runner.go) reads the `.so` files in all the plugins directory
and invokes `Check` method. It creates a new type [ExtendedMaturityCheck](./depchecker.go)
with the same underlying type `MaturityCheck` but adds new methods
to check the EOL of the stack.

### Working with the configuration variables

There are two ways by which configuration values could be passed.
* Updating the values in `conf.toml`
* Exporting the values in the shell that will execute the program

Values from `conf.toml` will be prioritized over values exported
from the executing shell. The snippet from [init.go](conf/initialize/init.go)
shows the handling of configuration variables.
```go
util.GetOR(config["TOKEN"].(string), os.Getenv("TOKEN"))
```

### Running the test cases
```shell
go test -v ./...
```
To also print the code coverage use the following command
```shell
go test -v ./... -cover
```

_Note: If you are not exporting env variables through shell, you'll need to copy conf.toml inside the PLUGIN_NAME directory to see if GitHub dependent plugin test cases are passing_
