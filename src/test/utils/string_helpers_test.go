package utils_test

import (
	"testing"

	"github.com/BlankDev117/gomockserver/src/utils"
	"github.com/stretchr/testify/assert"
)

// #region SubstrLength

func TestSubstrLengthNegativeStartIndexPanics(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	startIndex := -1
	str := "Hello World"

	// Act/Assert
	assert.Panics(func() { utils.SubstrLength(str, startIndex, 1) })
}

func TestSubstrLengthStartGreaterThanStrReturnsEmptyStr(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	startIndex := 100
	str := "Hello World"

	// Act/Assert
	assert.Panics(func() { utils.SubstrLength(str, startIndex, 1) })
}

func TestSubstrLengthLengthGreaterThanStringReturnsStringFromStartIndex(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	startIndex := 0
	str := "Hello World"

	// Act
	result := utils.SubstrLength(str, startIndex, 100)

	// Assert
	assert.NotNil(result)
	assert.Equal(str, result)
}

func TestSubstrLengthLengthLessThanStringReturnsSubstringFromStartIndexToLength(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	startIndex := 6
	str := "Hello World"

	// Act
	result := utils.SubstrLength(str, startIndex, 5)

	// Assert
	assert.NotNil(result)
	assert.Equal("World", result)
}

// #endregion
