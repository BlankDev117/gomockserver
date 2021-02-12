package routes

import (
	"strings"

	"github.com/BlankDev117/gomockserver/src/utils"
)

// Route is a parsed configuration route from a settings file
type Route struct {
	RawPath    string
	RouteParts []string
	Method     string
}

// #region Constructors

// DefaultRoute creates a royte with default values set
func DefaultRoute() Route {
	return Route{}
}

// NewRoute creates a route using the specified path and method
func NewRoute(rawPath string, method string) Route {
	rawPath = utils.SanitizeRoute(rawPath)
	return Route{rawPath, strings.Split(rawPath, "/"), method}
}

// #endregion

// #region Helpers

// DoesPathMatch determines whether a specified route url partially (wild card '*') matches, fully (exact) matches or does not match the current route
func (route Route) DoesPathMatch(path string) RouteMatch {
	path = utils.SanitizeRoute(path)
	pathParts := strings.Split(path, "/")

	for index, part := range pathParts {
		if index >= len(route.RouteParts) {
			return NoMatch
		}

		routePart := route.RouteParts[index]
		if routePart == "*" {
			return PartialMatch
		}

		partsEqual := false
		if strings.HasPrefix(routePart, "{") && strings.HasSuffix(routePart, "}") {
			partsEqual = true
		} else {
			partsEqual = strings.EqualFold(routePart, part)
		}

		if index == len(pathParts)-1 && partsEqual && len(pathParts) == len(route.RouteParts) {
			return FullMatch
		}
		if !partsEqual {
			return NoMatch
		}
	}

	return NoMatch
}

// #endregion
