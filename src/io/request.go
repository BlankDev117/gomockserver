package io

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/BlankDev117/gomockserver/src/utils"
)

type request struct {
	Path            string
	Method          string
	PathParameters  map[string]string
	QueryParameters map[string]string
	Body            map[string]interface{}
}

// #region Constructors

// DefaultRequest generates a server request with default values set
func DefaultRequest() request {
	return request{"", "", nil, nil, nil}
}

// NewRequest will create a server request object converted from an http request object received by an api call
func NewRequest(httpRequest *http.Request, routeParts []string) (request, error) {
	if httpRequest == nil {
		panic("httpRequest can not be nil")
	}
	if routeParts == nil {
		panic("routeParts can not be nil")
	}

	pathParams := getPathParameters(httpRequest.URL.Path, routeParts)

	queryParams, err := convertQueryToMap(httpRequest.URL.RawQuery)

	if err != nil {
		return DefaultRequest(), err
	}

	bodyBuffer, err := ioutil.ReadAll(httpRequest.Body)

	if err != nil {
		return DefaultRequest(), err
	}

	// Remove byte order mark
	// Reference: https://stackoverflow.com/questions/31398044/got-error-invalid-character-%C3%AF-looking-for-beginning-of-value-from-json-unmar
	bodyBuffer = bytes.TrimPrefix(bodyBuffer, []byte("\xef\xbb\xbf"))
	body, err := utils.ConvertBytesToMap(bodyBuffer)

	if err != nil {
		return DefaultRequest(), err
	}

	return request{httpRequest.URL.Path, httpRequest.Method, pathParams, queryParams, body}, nil
}

// #endregion

// #region Helpers

func getPathParameters(path string, routeParts []string) map[string]string {
	path = utils.SanitizeRoute(path)
	pathParts := strings.Split(path, "/")

	pathParamters := map[string]string{}
	for index, part := range pathParts {
		if index >= len(routeParts) {
			break
		}

		configRoutePart := routeParts[index]
		if strings.HasPrefix(configRoutePart, "{") && strings.HasSuffix(configRoutePart, "}") {
			pathParamters[configRoutePart] = part
		}
	}

	return pathParamters
}

func convertQueryToMap(query string) (map[string]string, error) {
	if query == "" {
		return map[string]string{}, nil
	}

	queryParts := strings.Split(query, ",")
	queryMap := make(map[string]string)

	for index := range queryParts {
		queryPair := strings.Split(queryParts[index], "=")

		if len(queryPair) != 2 {
			return nil, errors.New("Invalid query parameter pair")
		}

		queryMap[queryPair[0]] = queryPair[1]
	}

	return queryMap, nil
}

// #endregion
