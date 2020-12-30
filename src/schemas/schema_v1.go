package schemas

import (
	"errors"
	"strings"

	"github.com/BlankDev117/gomockserver/src/io"
	"github.com/BlankDev117/gomockserver/src/routes"
	"github.com/BlankDev117/gomockserver/src/utils"
)

// SchemaV1 is the standard schema for v1
type SchemaV1 struct {
}

// #region Variables

const schemaVersionKey = "schema"
const supportedSchemaVersion = "v1"
const routesKey = "routes"
const statusCodeKey = "statusCode"
const bodyKey = "body"

// #endregion

// #region Constructors

// DefaultSchemaV1 creates a default schema that can read version 1 schemas
func DefaultSchemaV1() SchemaV1 {
	return SchemaV1{}
}

// #endregion

// #region ischemaparser

// CanReadMap Determines whether this schema can read the provided map
func (schema SchemaV1) CanReadMap(schemaMap map[string]interface{}) bool {
	if schemaMap == nil {
		return false
	}

	if version, containsKey := schemaMap[schemaVersionKey]; containsKey {
		return version.(string) == supportedSchemaVersion
	}

	_, containsKey := schemaMap[routesKey]

	return containsKey
}

// ReadMap Converts the provided schema map into manageable route maps
func (schema SchemaV1) ReadMap(schemaMap map[string]interface{}) ([]routes.RouteMap, error) {
	if !schema.CanReadMap(schemaMap) {
		return nil, errors.New("Schema could not read the schema map")
	}

	routeMap := []routes.RouteMap{}

	routes := schemaMap[routesKey].(map[string]interface{})

	for routeURL, routeConfiguration := range routes {
		routeConfig, err := parseRoutes(routeURL, routeConfiguration.(map[string]interface{}))

		if err != nil {
			return nil, err
		}

		for _, route := range routeConfig {
			routeMap = append(routeMap, route)
		}
	}

	return routeMap, nil
}

// #endregion

// #region Helpers

func parseRoutes(url string, routeConfig map[string]interface{}) ([]routes.RouteMap, error) {
	routeMaps := []routes.RouteMap{}
	for method, configResponse := range routeConfig {

		serverRoute, err := createServerRoute(method, url)

		if err != nil {
			return nil, err
		}

		serverResponse, err := createServerResponse(configResponse.(map[string]interface{}))

		if err != nil {
			return nil, err
		}

		route := routes.RouteMap{Route: serverRoute, Response: serverResponse}

		routeMaps = append(routeMaps, route)
	}

	return routeMaps, nil
}

func createServerRoute(method string, url string) (routes.Route, error) {
	url = strings.Trim(url, "/ ")
	method = strings.TrimSpace(method)

	// method should not include { } param characters
	if index := strings.Index(method, "{"); index >= 0 {
		method = utils.SubstrLength(method, 0, index)
	}

	return routes.NewRoute(url, method), nil
}

func createServerResponse(responseBody map[string]interface{}) (io.Response, error) {

	if _, containsKey := responseBody[statusCodeKey]; !containsKey {
		return io.DefaultResponse(), errors.New("Provided response body does not fit desired schema. Key '" + statusCodeKey + "' is missing")
	}

	statusCode, err := utils.GetIntFromGeneric(responseBody[statusCodeKey])

	if err != nil {
		return io.DefaultResponse(), err
	}

	body := map[string]interface{}{}

	if responseBody[bodyKey] == nil {
		body = nil
	} else {
		body = responseBody[bodyKey].(map[string]interface{})
	}

	return io.NewResponse(statusCode, body), nil
}

// #endregion
