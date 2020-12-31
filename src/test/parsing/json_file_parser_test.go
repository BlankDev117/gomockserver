package parsing_test

import (
	"testing"

	"github.com/BlankDev117/gomockserver/src/parsing"
	"github.com/stretchr/testify/assert"
)

// #region SupportedExtensions

func TestSupportedExtensionsReturnsJson(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	parser := parsing.JsonFileParser{}

	// Act
	result := parser.SupportedExtensions()

	// Assert
	assert.NotNil(result)
	assert.Contains(result, ".json")
}

// #endregion

// #region ParseFile

func TestParseFileBadFileReturnsError(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	parser := parsing.JsonFileParser{}

	// Act
	result, err := parser.ParseFile("badfile/path.txt")

	// Assert
	assert.Nil(result)
	assert.NotNil(err)
}

func TestParseFileValidJSONReturnsMap(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	parser := parsing.JsonFileParser{}

	// Act
	result, err := parser.ParseFile("../testdata/testApi.json")

	// Assert
	assert.Nil(err)
	assert.NotNil(result)

	assert.Contains(result, "routes")

	configRoutes := result["routes"].(map[string]interface{})

	assert.Contains(configRoutes, "/url/subUrl")
	assert.Contains(configRoutes, "url/*")
}

// #endregion
