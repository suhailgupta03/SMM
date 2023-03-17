package util

import "reflect"

// GetOR returns the result of running logical OR operator on
// two arguments. If the first argument is empty then the second argument
// will be returned by default and vice versa.
func GetOR[T interface{}](itemOne, itemTwo T) T {
	t := reflect.TypeOf(itemOne).String()
	iOne := interface{}(itemOne)
	switch t {
	case "string":
		if iOne.(string) != "" {
			return itemOne
		}
		return itemTwo
	default:
		// Note: add a test case when handling a new type above
		panic("Type not handled yet!")
	}
}
