package schemas

import (
	"errors"
)

// SchemaProvider assists with getting an appropriate schema for a generic map, if one is available
type SchemaProvider struct {
	Schemas []ISchema
}

// #region Constructors

// DefaultSchemaProvider creates a default schema provider
func DefaultSchemaProvider() SchemaProvider {
	return SchemaProvider{[]ISchema{DefaultSchemaV1()}}
}

// #endregion

// #region helpers

// GetSchema retrieves a schema that is capable of parsing the provided generic schema map
func (provider SchemaProvider) GetSchema(schemaMap map[string]interface{}) (ISchema, error) {

	if schemaMap == nil {
		panic("Schema map can not be nil")
	}

	for _, schema := range provider.Schemas {
		if schema.CanReadMap(schemaMap) {
			return schema, nil
		}
	}

	return nil, errors.New("No schemas were found that could support the provided configuration map")
}

// #endregion
