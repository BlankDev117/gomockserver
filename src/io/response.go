package io

import (
	"encoding/json"
	"strconv"
	"strings"
)

// Response contains the status code and body to return for a given api call as specified by the user in the api settings file
type Response struct {
	StatusCode int
	Body       map[string]interface{}
}

// #region Constructors

// DefaultResponse generates a response with default values set
func DefaultResponse() Response {
	return Response{}
}

// NewResponse creates a response with the specified values
func NewResponse(statusCode int, body map[string]interface{}) Response {
	return Response{statusCode, body}
}

// #endregion

// #region Helpers

// GetJSONResponse Returns the string json for the response object with the provided url params substituted
func GetJSONResponse(response Response, urlParams map[string]string) (string, error) {

	if response.Body == nil || len(response.Body) == 0 {
		return "", nil
	}

	jsonData, err := json.Marshal(response.Body)

	if err != nil {
		return "", err
	}

	if urlParams == nil || len(urlParams) == 0 {
		return string(jsonData), nil
	}

	replacements := []string{}

	for paramKey, paramValue := range urlParams {
		if isNumericParameter(paramValue) {
			replacements = append(replacements, "\"{"+paramKey+"}\"")
		} else {
			replacements = append(replacements, "{"+paramKey+"}")
		}

		replacements = append(replacements, paramValue)
	}

	replacer := strings.NewReplacer(replacements...)

	return replacer.Replace(string(jsonData)), nil
}

func isNumericParameter(param string) bool {
	_, err := strconv.Atoi(param)

	return err == nil
}

// #endregion
