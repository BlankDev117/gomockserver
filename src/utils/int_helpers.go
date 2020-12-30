package utils

import (
	"reflect"
	"strconv"
)

// GetIntFromGeneric attempts to parse a given object and return its numeric value
func GetIntFromGeneric(input interface{}) (int, error) {
	if statusCode, isString := input.(string); isString {
		return strconv.Atoi(statusCode)
	}
	if statusCode, isInt := input.(int); isInt {
		return statusCode, nil
	}
	if statusCode, isFloat := input.(float64); isFloat {
		return int(statusCode), nil
	}

	panic("Provided status code could not be parsed into numeric form: " + reflect.TypeOf(input).Name())
}
