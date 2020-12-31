package routes_test

import (
	"testing"

	"github.com/BlankDev117/gomockserver/src/routes"
	"github.com/stretchr/testify/assert"
)

// #region DoesPathMatch

func TestDoesPathMatchEmptyPathReturnsNoMatch(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	route := routes.NewRoute("hello/world", "Post")

	// Act
	result := route.DoesPathMatch("")

	// Assert
	assert.Equal(routes.NoMatch, result)
}

func TestDoesPathMatchDifferentPathReturnsNoMatch(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	route := routes.NewRoute("hello/world", "Post")

	// Act
	result := route.DoesPathMatch("hello/planet")

	// Assert
	assert.Equal(routes.NoMatch, result)
}

func TestDoesPathMatchWildCardPathReturnsPartialMatch(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	route := routes.NewRoute("hello/*", "Post")

	// Act
	result := route.DoesPathMatch("hello/world/planet")

	// Assert
	assert.Equal(routes.PartialMatch, result)
}

func TestDoesPathMatchExactPathReturnsFullMatch(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	route := routes.NewRoute("hello/world/planet", "Post")

	// Act
	result := route.DoesPathMatch("hello/world/planet")

	// Assert
	assert.Equal(routes.FullMatch, result)
}

func TestDoesPathMatchPathWithPlaceholdersReturnsFullMatch(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	route := routes.NewRoute("hello/{world}/planet/{id}", "Post")

	// Act
	result := route.DoesPathMatch("hello/earth/planet/1")

	// Assert
	assert.Equal(routes.FullMatch, result)
}

func TestDoesPathMatchPathTrailingAndsReturnsFullMatch(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	route := routes.NewRoute("hello/{world}/planet/{id}", "Post")

	// Act
	result := route.DoesPathMatch("hello/earth/planet/1")

	// Assert
	assert.Equal(routes.FullMatch, result)
}

func TestDoesPathMatchPathRouteLengthDifferentThanUrlNoMatch(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	route := routes.NewRoute("hello/{world}/planet/{id}", "Post")

	// Act
	result := route.DoesPathMatch("hello/earth/planet")

	// Assert
	assert.Equal(routes.NoMatch, result)
}

// #endregion
