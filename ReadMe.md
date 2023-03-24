[![Run Tests](https://github.com/suhailgupta03/cuddly-eureka-/actions/workflows/test.yml/badge.svg)](https://github.com/suhailgupta03/cuddly-eureka-/actions/workflows/test.yml)

### Generate .so Files and Run Code
To generate the `.so` files and run the code, execute the following script
```shell
./build_run.sh
```

### Available Plugins
* [NODE EOL](plugins/nodeeol/nodeeol.go)
* [DJANGO EOL](plugins/djangoeol/djangoeol.go)
* [REACT EOL](plugins/reacteol/reacteol.go)
* [README](plugins/readme/readme.go)
* [PYTHON EOL](plugins/pythoneol/pythoneol.go)

### Description of Maturity Values
[MaturityValues](types/commontypes.go) are defined here.

| 1   | YES | EOL / FOUND         | 
|-----|-----|---------------------|
| 2   | NO  | NOT EOL / NOT FOUND |
| 0   | NA  | -                   |

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
To export the environment variables, run
```shell
source test.env
```

`test.env` will have variables as shown below exported to the shell running the code

```shell
export STAGE=test
export TOKEN=
export OWNER=
export NODE=issue-test
export EMPTY=
```

### Running the test cases
Once the variables have been exported, the tests could be run as follows
```shell
go test -v ./...
```
To also print the code coverage use the following command
```shell
go test -v ./... -cover
```
