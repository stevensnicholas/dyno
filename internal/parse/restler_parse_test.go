package parse_test

import (
	"dyno/internal/parse"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseRestlerFuzzResultsInvalid(t *testing.T) {
	fileContents := "../tests/bug_buckets/"
	fuzzError := strings.Split("example", "example")
	assert.Panics(t, func() { parse.ParseRestlerFuzzResults(fileContents, fuzzError) })
}
func TestParseRestlerFuzzResultsValid(t *testing.T) {
	fileContents := "../tests/bug_buckets/"
	fuzzError := strings.Split("example", "example")
	expectedDynoResults := [11]string{
		"InvalidDynamicObjectChecker Invalid 20x Response",
		"InvalidDynamicObjectChecker Invalid 20x Response",
		"InvalidDynamicObjectChecker Invalid 20x Response",
		"InvalidDynamicObjectChecker Invalid 20x Response",
		"PayloadBodyChecker Invalid 500 Response",
		"PayloadBodyChecker Invalid 500 Response",
		"PayloadBodyChecker Invalid 500 Response",
		"PayloadBodyChecker Invalid 500 Response",
		"UseAfterFreeChecker Invalid 20x Response",
		"UseAfterFreeChecker Invalid 20x Response",
		"UseAfterFreeChecker Invalid 20x Response",
	}
	actualDynoResults := parse.ParseRestlerFuzzResults(fileContents, fuzzError)
	println(*actualDynoResults[0].Title)
	println(*actualDynoResults[0].Endpoint)
	println(*actualDynoResults[0].Method)
	println(*actualDynoResults[0].MethodInformation.AcceptedResponse)
	println(*actualDynoResults[0].MethodInformation.Host)
	println(*actualDynoResults[0].MethodInformation.ContentType)
	println(*actualDynoResults[0].MethodInformation.Request)
	println(*actualDynoResults[0].TimeDelay)
	println(*actualDynoResults[0].AsyncTime)
	println(*actualDynoResults[0].PreviousResponse)
	println(*actualDynoResults[0].ErrorType)
	assert.Equal(t, expectedDynoResults[0], *actualDynoResults[0].Title)
	assert.Equal(t, expectedDynoResults[1], *actualDynoResults[1].Title)
	assert.Equal(t, expectedDynoResults[2], *actualDynoResults[2].Title)
	assert.Equal(t, expectedDynoResults[3], *actualDynoResults[3].Title)
	assert.Equal(t, expectedDynoResults[4], *actualDynoResults[4].Title)
	assert.Equal(t, expectedDynoResults[5], *actualDynoResults[5].Title)
	assert.Equal(t, expectedDynoResults[6], *actualDynoResults[6].Title)
	assert.Equal(t, expectedDynoResults[7], *actualDynoResults[7].Title)
	assert.Equal(t, expectedDynoResults[8], *actualDynoResults[8].Title)
	assert.Equal(t, expectedDynoResults[9], *actualDynoResults[9].Title)
	assert.Equal(t, expectedDynoResults[10], *actualDynoResults[10].Title)
}
