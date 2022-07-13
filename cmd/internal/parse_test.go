package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReadBugFileInvalidNoFile(t *testing.T) {
	testBugFile := ""
	body := ""
	testBody, testEndpoint, err := ReadBugFile(testBugFile, body)
	assert.Equal(t, "", testBody)
	assert.Equal(t, "", testEndpoint)
	assert.Equal(t, "", err)
}

func TestReadBugFileValid(t *testing.T) {
	testBugFile := "InvalidDynamicObjectChecker_20x_1.txt"
	body := ""
	testBody, testEndpoint, err := ReadBugFile(testBugFile, body)
	assert.Equal(t, "", testBody)
	assert.Equal(t, "", testEndpoint)
	assert.Equal(t, "", err)
}