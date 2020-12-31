package io_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	serverio "github.com/BlankDev117/gomockserver/src/io"
	"github.com/stretchr/testify/assert"
)

// #region NewRequest

func TestNewRequestNilRequestPanics(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	// Act/Assert
	assert.Panics(func() { serverio.NewRequest(nil, []string{}) })
}

func TestNewRequestNilRoutePartsPanics(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	request := http.Request{}

	// Act/Assert
	assert.Panics(func() { serverio.NewRequest(&request, nil) })
}

func TestNewRequestValidRequestReturnsNewRequest(t *testing.T) {
	// Arrange
	assert := assert.New(t)

	jsonStr := "{\"Name\": \"House On Main Street\", \"Cost\": 1000}"
	requestPath := "realestate/88/houses/wood/214"
	query := "color=red,petsAllowed=true"
	configuredRoute := "realestate/{tenantId}/houses/wood/{id}"
	routeParts := strings.Split(configuredRoute, "/")
	method := "Post"
	request := http.Request{URL: &url.URL{Path: requestPath, RawQuery: query}, Method: method, Body: ioutil.NopCloser(strings.NewReader(jsonStr))}

	// Act
	result, err := serverio.NewRequest(&request, routeParts)

	// Assert
	assert.Nil(err)
	assert.NotNil(result)

	assert.Equal(requestPath, result.Path)
	assert.Equal(method, result.Method)

	assert.NotNil(result.PathParameters)
	assert.Len(result.PathParameters, 2)
	assert.Contains(result.PathParameters, "{tenantId}")
	assert.Equal("88", result.PathParameters["{tenantId}"])
	assert.Contains(result.PathParameters, "{id}")
	assert.Equal("214", result.PathParameters["{id}"])

	assert.NotNil(result.QueryParameters)
	assert.Len(result.QueryParameters, 2)
	assert.Contains(result.QueryParameters, "color")
	assert.Equal("red", result.QueryParameters["color"])
	assert.Contains(result.QueryParameters, "petsAllowed")
	assert.Equal("true", result.QueryParameters["petsAllowed"])

	assert.NotNil(result.Body)
	assert.Len(result.Body, 2)
	assert.Contains(result.Body, "Name")
	assert.Equal("House On Main Street", result.Body["Name"])
	assert.Contains(result.Body, "Cost")
	assert.Equal(float64(1000), result.Body["Cost"])
}

// #endregion
