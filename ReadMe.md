[![Run Tests](https://github.com/suhailgupta03/cuddly-eureka-/actions/workflows/test.yml/badge.svg)](https://github.com/suhailgupta03/cuddly-eureka-/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/suhailgupta03/cuddly-eureka-/branch/main/graph/badge.svg?token=tNKcOjlxLo)](https://codecov.io/gh/suhailgupta03/cuddly-eureka-)

### Build plugins and run 
To generate the `.so` files and run the code, execute the following script
```shell
make run-local
```

### Generate binary and run SMM binary
```shell
make build
./smm
```

### Create the distribution build
```shell
make dist
```

### Repository scan configuration
and describe the configuration inside [repo-details.yml](assets/repo-details.yml) in a structure that looks similar to the following.
```yaml
name: Repository Details
repository:
  - name: virality
    ecr: xxx.dkr.ecr.us-east-1.amazonaws.com/ci:v1.3.1
    aws:
      log-group-name: playground
      log-stream-name: playground-stream
```

### Passing flags to the binary
```shell
./smm  repo -yml=scan-details.yml
  github -token=SECRET_TOKEN -owner=GITHUB_OWNER
```

### Available Plugins
* [NODE EOL](plugins/nodeeol/nodeeol.go)
* [DJANGO EOL](plugins/djangoeol/djangoeol.go)
* [REACT EOL](plugins/reacteol/reacteol.go)
* [README](plugins/readme/readme.go)
* [PYTHON EOL](plugins/pythoneol/pythoneol.go)
* [REPOVULN](plugins/repovuln/repovuln.go)
  * Uses [trivy](https://github.com/aquasecurity/trivy) to scan repos
  * All linked tests run with [version 0.38.3](https://github.com/aquasecurity/trivy/releases/tag/v0.38.3) 
  * GitHub workflow also assumes the above version
  * For the plugin to be able to scan private repositories, value to `GITHUB_TOKEN` must be provided in the shell that is executing the code. See, [test.env](./test.env) for example. This is used internally by trivy [as written in the documentation here](https://aquasecurity.github.io/trivy/v0.38/docs/target/git-repository/)
* [ECRVULN](plugins/ecrvuln/ecrvuln.go)
  * Uses [trivy](https://github.com/aquasecurity/trivy) to scan repos
  * All linked tests run with [version 0.38.3](https://github.com/aquasecurity/trivy/releases/tag/v0.38.3)
  * For the plugin to able to scan private ECR images values to `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY` and `AWS_DEFAULT_REGION` must be provided in the shell as described in the [trivy docs](https://aquasecurity.github.io/trivy/v0.38/docs/advanced/private-registries/ecr/)
* [LATESTPATCHDJANGO](plugins/latestpatchdjango/latestpatchdjango.go)
* [LATESTPATCHNODE](plugins/latestpatchnode/latestpatchnode.go)
* [LATESTPATCHPYTHON](plugins/latestpatchpython/latestpatchpython.go)
* [HASLOGGING](plugins/haslogging/haslogging.go)
  * Uses AWS config stored inside `~/.aws/config` or the AWS ENV exported in the shell running the program
* [HASJSONLOGGING](plugins/hasjsonlogging/hasjsonlogging.go)
  * Uses AWS config stored inside `~/.aws/config` or the AWS ENV exported in the shell running the program

### Description of Maturity Values
[MaturityValues](types/commontypes.go) are defined here.

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
To open coverage report along with running the test cases
```shell
make test
```
