package parsing_test

import (
	"testing"

	"github.com/BlankDev117/gomockserver/src/parsing"
	"github.com/stretchr/testify/assert"
)

// #region Helpers

type testParser struct {
}

func (p testParser) SupportedExtensions() []string {
	return []string{".test"}
}

func (p testParser) ParseFile(filePath string) (map[string]interface{}, error) {
	panic("Not Implemented")
}

// #endregion

// #region GetParser

func TestGetParserNoParserForExtensionReturnsError(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	provider := parsing.FileParserProvider{Parsers: []parsing.IFileParser{}}

	// Act
	result, err := provider.GetParser("./helloWorld.txt")

	// Assert
	assert.Nil(result)
	assert.NotNil(err)
}

func TestGetParserHasParserForExtensionReturnsParser(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	p := testParser{}

	provider := parsing.FileParserProvider{Parsers: []parsing.IFileParser{p}}

	// Act
	result, err := provider.GetParser("./helloWorld.test")

	// Assert
	assert.Nil(err)
	assert.NotNil(result)
	assert.Equal(p, result)
}

// #endregion
