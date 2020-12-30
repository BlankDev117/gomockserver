package schemas_test

import (
	"testing"

	"github.com/BlankDev117/gomockserver/src/routes"
	"github.com/BlankDev117/gomockserver/src/schemas"
	"github.com/stretchr/testify/assert"
)

// #region Helpers

type testSchema struct {
}

func (s testSchema) CanReadMap(schemaMap map[string]interface{}) bool {
	return true
}

func (s testSchema) ReadMap(schemaMap map[string]interface{}) ([]routes.RouteMap, error) {
	panic("Not Implemented")
}

// #endregion

// #region GetSchema

func TestGetSchemaNilMapPanics(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	provider := schemas.SchemaProvider{}

	// Act/Assert
	assert.Panics(func() { provider.GetSchema(nil) })
}

func TestGetSchemaNoSchemaSupportsMapReturnsError(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	provider := schemas.SchemaProvider{}
	schemaMap := map[string]interface{}{}

	// Act
	result, err := provider.GetSchema(schemaMap)

	// Assert
	assert.Nil(result)
	assert.NotNil(err)
}

func TestGetSchemaSchemaSupportsMapReturnsSchema(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	schema := testSchema{}
	provider := schemas.SchemaProvider{Schemas: []schemas.ISchema{schema}}
	schemaMap := map[string]interface{}{}

	// Act
	result, err := provider.GetSchema(schemaMap)

	// Assert
	assert.Nil(err)
	assert.NotNil(result)
	assert.Equal(schema, result)
}

// #endregion
