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
func (hat HasAutomatedTests) Check() types.MaturityCheck {
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