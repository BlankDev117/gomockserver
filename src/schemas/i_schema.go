package schemas

import "github.com/BlankDev117/gomockserver/src/routes"

// ISchema is able to convert a generic schema map into a manageable set of routes
type ISchema interface {
	CanReadMap(schemaMap map[string]interface{}) bool
	ReadMap(schemaMap map[string]interface{}) ([]routes.RouteMap, error)
}
