package utils_test

import (
	"fmt"
	"testing"

	"github.com/BlankDev117/gomockserver/src/utils"
	"github.com/stretchr/testify/assert"
)

// #region ConvertBytesToMap

func TestConvertBytesToMapNilByteArrayReturnsError(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act
	result, err := utils.ConvertBytesToMap(nil)

	// Assert
	assert.Nil(result)
	assert.NotNil(err)
}

func TestConvertBytesToMapValidByteArrayReturnsBytesAsMap(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	testKey := "test"
	testValue := "amazing"
	jsonStr := fmt.Sprintf("{\"%s\": \"%s\"}", testKey, testValue)

	// Act
	result, err := utils.ConvertBytesToMap([]byte(jsonStr))

	// Assert
	assert.Nil(err)
	assert.NotNil(result)

	assert.Contains(result, testKey)

	resultValue := result[testKey]

	assert.Equal(testValue, resultValue)
}

// #endregion
