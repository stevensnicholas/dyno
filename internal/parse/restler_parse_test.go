package parse_test

import (
	"os"
	"dyno/internal/parse"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseInvalidFuzzError(t *testing.T) {
	assert.Panics(t, func() { parse.GetFuzzError("something") })
}
func TestParseRestlerFuzzResultsValid(t *testing.T) {
	file, _ := os.ReadFile("../tests/bug_buckets/InvalidDynamicObjectChecker_20x_1.txt")
	fileContents := string(file)
	fuzzError := parse.GetFuzzError("InvalidDynamicObjectChecker_20x_1.txt")
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
	assert.Equal(t, expectedDynoResults[0], *actualDynoResults[0].Title)
	assert.Equal(t, expectedDynoResults[1], *actualDynoResults[1].Title)

	file, _ = os.ReadFile("../tests/bug_buckets/InvalidDynamicObjectChecker_20x_2.txt")
	fileContents = string(file)
	fuzzError = parse.GetFuzzError("InvalidDynamicObjectChecker_20x_2.txt")
	actualDynoResults = parse.ParseRestlerFuzzResults(fileContents, fuzzError)
	assert.Equal(t, expectedDynoResults[2], *actualDynoResults[0].Title)
	assert.Equal(t, expectedDynoResults[3], *actualDynoResults[1].Title)

	file, _ = os.ReadFile("../tests/bug_buckets/PayloadBodyChecker_500_1.txt")
	fileContents = string(file)
	fuzzError = parse.GetFuzzError("PayloadBodyChecker_500_1.txt")
	actualDynoResults = parse.ParseRestlerFuzzResults(fileContents, fuzzError)
	assert.Equal(t, expectedDynoResults[4], *actualDynoResults[0].Title)
	assert.Equal(t, expectedDynoResults[5], *actualDynoResults[1].Title)

	file, _ = os.ReadFile("../tests/bug_buckets/PayloadBodyChecker_500_2.txt")
	fileContents = string(file)
	fuzzError = parse.GetFuzzError("PayloadBodyChecker_500_2.txt")
	actualDynoResults = parse.ParseRestlerFuzzResults(fileContents, fuzzError)
	assert.Equal(t, expectedDynoResults[6], *actualDynoResults[0].Title)
	assert.Equal(t, expectedDynoResults[7], *actualDynoResults[1].Title)
	
	file, _ = os.ReadFile("../tests/bug_buckets/UseAfterFreeChecker_20x_1.txt")
	fileContents = string(file)
	fuzzError = parse.GetFuzzError("UseAfterFreeChecker_20x_1.txt")
	actualDynoResults = parse.ParseRestlerFuzzResults(fileContents, fuzzError)
	assert.Equal(t, expectedDynoResults[8], *actualDynoResults[0].Title)
	assert.Equal(t, expectedDynoResults[9], *actualDynoResults[1].Title)
	assert.Equal(t, expectedDynoResults[10], *actualDynoResults[2].Title)
}
