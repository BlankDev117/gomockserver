package parsing_test

import (
	"testing"

	"github.com/BlankDev117/gomockserver/src/parsing"
	"github.com/stretchr/testify/assert"
)

// #region CanParse

func TestCanParseEmptyFilePath(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	parser := testParser{}

	// Act
	result := parsing.CanParse(parser, "")

	// Assert
	assert.False(result)
}

func TestCanParseExtensionNotSupportedReturnsFalse(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	parser := testParser{}

	// Act
	result := parsing.CanParse(parser, "testFile.json")

	// Assert
	assert.False(result)
}

func TestCanParseExtensionSupportedReturnsTrue(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	parser := testParser{}

	// Act
	result := parsing.CanParse(parser, "testFile.test")

	// Assert
	assert.True(result)
}

// #endregion
