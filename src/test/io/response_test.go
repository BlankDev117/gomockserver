package io_test

import (
	"testing"

	"github.com/BlankDev117/gomockserver/src/io"
	"github.com/stretchr/testify/assert"
)

// #region GetJSONResponse

func TestGetJSONResponseNilBodyReturnsEmptyString(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	response := io.NewResponse(200, nil)

	// Act
	result, err := io.GetJSONResponse(response, nil)

	// Assert
	assert.Nil(err)
	assert.Equal("", result)
}

func TestGetJSONResponseEmptyBodyReturnsEmptyString(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	response := io.NewResponse(200, map[string]interface{}{})

	// Act
	result, err := io.GetJSONResponse(response, nil)

	// Assert
	assert.Nil(err)
	assert.Equal("", result)
}

func TestGetJSONResponseNilUrlParamsReturnsBodyAsJSONString(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	jsonBody := map[string]interface{}{
		"Test": "Hi",
		"Oh":   10,
	}
	response := io.NewResponse(200, jsonBody)

	// Act
	result, err := io.GetJSONResponse(response, nil)

	// Assert
	assert.Nil(err)
	assert.Equal("{\"Oh\":10,\"Test\":\"Hi\"}", result)
}

func TestGetJSONResponseEmptyUrlParamsReturnsBodyAsJSONString(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	jsonBody := map[string]interface{}{
		"Test": "Hi",
		"Oh":   10,
	}
	response := io.NewResponse(200, jsonBody)

	urlParams := map[string]string{}

	// Act
	result, err := io.GetJSONResponse(response, urlParams)

	// Assert
	assert.Nil(err)
	assert.Equal("{\"Oh\":10,\"Test\":\"Hi\"}", result)
}

func TestGetJSONResponseReplacesDesignatedUrlParamsInJSONString(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	jsonBody := map[string]interface{}{
		"Test":      "Hi",
		"Id":        "{{id}}",
		"Oh":        10,
		"Color":     "{{color}}",
		"Secondary": "{color}",
	}
	response := io.NewResponse(200, jsonBody)

	urlParams := map[string]string{
		"{color}": "red",
		"{id}":    "117",
		"{random": "Hi",
	}

	// Act
	result, err := io.GetJSONResponse(response, urlParams)

	// Assert
	assert.Nil(err)
	assert.Equal("{\"Color\":\"red\",\"Id\":117,\"Oh\":10,\"Secondary\":\"{color}\",\"Test\":\"Hi\"}", result)
}

// #endregion
