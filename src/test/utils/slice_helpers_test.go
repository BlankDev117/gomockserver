package utils_test

import (
	"testing"

	"github.com/BlankDev117/gomockserver/src/utils"
	"github.com/stretchr/testify/assert"
)

// #region StringInSlice

func TestStringInSliceNilSliceReturnsFalse(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	searchStr := "hello"

	// Act
	result := utils.StringInSlice(searchStr, nil)

	// Assert
	assert.False(result)
}

func TestStringInSliceSliceDoesNotContainStringReturnsFalse(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	searchStr := "hello"
	list := []string{"hi", "too late", "no way"}

	// Act
	result := utils.StringInSlice(searchStr, list)

	// Assert
	assert.False(result)
}

func TestStringInSliceSliceContainsStringReturnsTrue(t *testing.T) {
	// Arrange
	assert := assert.New(t)
	searchStr := "hello"
	list := []string{"hi", "too late", searchStr, "no way"}

	// Act
	result := utils.StringInSlice(searchStr, list)

	// Assert
	assert.True(result)
}

// #endregion
