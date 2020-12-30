package routes

import (
	"github.com/BlankDev117/gomockserver/src/io"
)

// RouteMap contains the user specified route and associated response as defined in the api configuration file
type RouteMap struct {
	Route    Route
	Response io.Response
}

// #region Constructors

// DefaultRouteMap creates a route map with default values set
func DefaultRouteMap() RouteMap {
	return RouteMap{DefaultRoute(), io.DefaultResponse()}
}

// #endregion
