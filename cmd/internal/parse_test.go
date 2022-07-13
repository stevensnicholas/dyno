package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReadBugFileInvalidNoFile(t *testing.T) {
	testBugFile := ""
	body := ""
	testBody, testEndpoint := ReadBugFile(testBugFile, body)
	assert.Equal(t, "", testBody)
	assert.Equal(t, "", testEndpoint)
}