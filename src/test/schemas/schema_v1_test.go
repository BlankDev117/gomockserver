package schemas_test

import (
	"testing"

	"github.com/BlankDev117/gomockserver/src/schemas"
	"github.com/stretchr/testify/assert"
)

// #region CanReadMap

func TestCanReadMapNilMapReturnsFalse(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	schema := schemas.DefaultSchemaV1()

	// Act
	result := schema.CanReadMap(nil)

	// Assert
	assert.False(result)
}

func TestCanReadMapVersionNotSupportedReturnsFalse(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	schema := schemas.DefaultSchemaV1()

	schemaMap := map[string]interface{}{
		"version": "v2",
	}

	// Act
	result := schema.CanReadMap(schemaMap)

	// Assert
	assert.False(result)
}

func TestCanReadMapRoutesMissingsReturnsFalse(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	schema := schemas.DefaultSchemaV1()

	schemaMap := map[string]interface{}{
		"version": "v1",
	}

	// Act
	result := schema.CanReadMap(schemaMap)

	// Assert
	assert.False(result)
}

func TestCanReadMapMissingVersionReturnsTrue(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	schema := schemas.DefaultSchemaV1()

	schemaMap := map[string]interface{}{
		"routes": map[string]interface{}{},
	}

	// Act
	result := schema.CanReadMap(schemaMap)

	// Assert
	assert.True(result)
}

func TestCanReadMapReturnsTrue(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	schema := schemas.DefaultSchemaV1()

	schemaMap := map[string]interface{}{
		"version": "v1",
		"routes":  map[string]interface{}{},
	}

	// Act
	result := schema.CanReadMap(schemaMap)

	// Assert
	assert.True(result)
}

// #endregion

// #region ReadMap

func TestReadMapInvalidSchemaMapReturnsError(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	schema := schemas.DefaultSchemaV1()

	// Act
	result, err := schema.ReadMap(nil)

	// Assert
	assert.Nil(result)
	assert.NotNil(err)
}

func TestReadMapMissingResponsePropertyReturnsError(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	schema := schemas.DefaultSchemaV1()

	schemaMap := map[string]interface{}{
		"routes": map[string]interface{}{
			"url/test": map[string]interface{}{
				"post": map[string]interface{}{
					"body": map[string]interface{}{},
				},
			},
		},
	}

	// Act
	result, err := schema.ReadMap(schemaMap)

	// Assert
	assert.Nil(result)
	assert.NotNil(err)
}

func TestReadMapValidSchemaReturnsExpectedRouteMaps(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	schema := schemas.DefaultSchemaV1()

	schemaMap := map[string]interface{}{
		"routes": map[string]interface{}{
			"url/test": map[string]interface{}{
				"post": map[string]interface{}{
					"statusCode": 200,
				},
				"put{id}": map[string]interface{}{
					"statusCode": 202,
					"body": map[string]interface{}{
						"Hello":   "there",
						"One":     1,
						"listStr": []string{"hi", "bye"},
						"listInt": []int{1, 2},
						"sub": map[string]interface{}{
							"on": "off",
						},
					},
				},
			},
			"url/pets/*": map[string]interface{}{
				"get": map[string]interface{}{
					"statusCode": 404,
				},
			},
		},
	}

	// Act
	result, err := schema.ReadMap(schemaMap)

	// Assert
	assert.Nil(err)
	assert.NotNil(result)

	assert.Len(result, 3)
	route := result[0]

	assert.NotNil(route)
	assert.Equal("url/test", route.Route.RawPath)
	assert.Equal("post", route.Route.Method)
	assert.Equal(200, route.Response.StatusCode)
	assert.Nil(route.Response.Body)

	route = result[1]
	assert.NotNil(route)
	assert.Equal("url/test", route.Route.RawPath)
	assert.Equal("put", route.Route.Method)
	assert.Equal(202, route.Response.StatusCode)
	assert.NotNil(route.Response.Body)
	assert.Len(route.Response.Body, 5)
	assert.Contains(route.Response.Body, "Hello")
	assert.Contains(route.Response.Body, "One")
	assert.Contains(route.Response.Body, "listStr")
	assert.Contains(route.Response.Body, "listInt")
	assert.Contains(route.Response.Body, "sub")
	assert.Len(route.Response.Body["sub"], 1)
	assert.Contains(route.Response.Body["sub"], "on")

	route = result[2]
	assert.NotNil(route)
	assert.Equal("url/pets/*", route.Route.RawPath)
	assert.Equal("get", route.Route.Method)
	assert.Equal(404, route.Response.StatusCode)
	assert.Nil(route.Response.Body)
}

// #endregion
