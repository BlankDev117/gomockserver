package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/BlankDev117/gomockserver/src/io"
)

type routeManager struct {
	routes []RouteMap
}

// #region Constructors

// NewRouteManager creates a route manager by providing an array of route maps
func NewRouteManager(routes []RouteMap) routeManager {
	return routeManager{routes}
}

// #endregion

// #region Helpers

func (routeManager routeManager) ProcessAPIRequest(writer http.ResponseWriter, request *http.Request) {
	urlPath := request.URL.Path

	log.Printf("Web request received on url path %s. Method: %s, Query Parameters: %s", urlPath, request.Method, request.URL.RawQuery)

	matchedRoute := DefaultRouteMap()
	matchType := NoMatch
	for _, route := range routeManager.routes {
		pathMatch := route.Route.DoesPathMatch(request.URL.Path)

		if !strings.EqualFold(route.Route.Method, request.Method) {
			continue
		}
		// Wildcard matching should match for any suburl
		if pathMatch == PartialMatch {
			matchType = PartialMatch
			matchedRoute = route
		}

		if pathMatch == FullMatch {
			matchType = FullMatch
			matchedRoute = route
			break
		}
	}

	log.Printf("Request processed. %s found.", matchType)

	statusCode := 404
	responseBody := "{\"Error\": \"No route was found that matched the request.\"}"
	headers := map[string]string{
		"content-type": "application/json",
	}

	if matchType != NoMatch {
		serverRequest, err := io.NewRequest(request, matchedRoute.Route.RouteParts)

		if err != nil {
			panic(err)
		}

		statusCode, responseBody, headers, err = getResponse(matchedRoute, serverRequest.PathParameters)

		if err != nil {
			panic(err)
		}
	}

	for headerKey, headerValue := range headers {
		writer.Header().Add(headerKey, headerValue)
	}
	writer.WriteHeader(statusCode)
	_, err := writer.Write([]byte(responseBody))

	if err != nil {
		fmt.Printf("Failed to write response, err %s", err)
	}
}

func getResponse(route RouteMap, pathParameters map[string]string) (int, string, map[string]string, error) {
	if route.Route.RawPath == "" {
		return 500, "Error: No route was found for the given request", map[string]string{
			"content-type": "application/json",
		}, nil
	}

	logPathParams(pathParameters)
	jsonResponse, err := io.GetJSONResponse(route.Response, pathParameters)

	if err != nil {
		return 500, "Error parsing response body", map[string]string{
			"content-type": "application/json",
		}, err
	}

	return route.Response.StatusCode, jsonResponse, route.Response.Headers, nil
}

func logPathParams(m map[string]string) {
	b := strings.Builder{}

	length := len(m)
	index := 0
	for key, value := range m {
		b.WriteString(key + ": " + value)

		if index < length {
			b.WriteString(", ")
		}

		index++
	}

	log.Printf("Path parameters: %s", b.String())
}

// #endregion
