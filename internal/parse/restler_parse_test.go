package parse_test

//TODO Talk about what to do when the file is wrong

import (
	"golambda/internal/parse"
	"strings"
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestParseRestlerFuzzResultsInvalid(t *testing.T) {
	location := "../tests/bug_buckets/"
	file := ""
	assert.Panics(t, func() {parse.ParseRestlerFuzzResults(location, file)})
}
func TestParseRestlerFuzzResultsValid(t *testing.T) {
	location := "../tests/bug_buckets/"
	file := "../tests/bug_buckets/bug_buckets.txt"
	expectedDynoResults := [5]string{
		"InvalidDynamicObjectChecker Invalid 20x Response",
		"InvalidDynamicObjectChecker Invalid 20x Response",
		"PayloadBodyChecker Invalid 500 Response",
		"PayloadBodyChecker Invalid 500 Response",
		"UseAfterFreeChecker Invalid 20x Response",
	}
	actualDynoResults := parse.ParseRestlerFuzzResults(location, file)
	assert.Equal(t, expectedDynoResults[0], *actualDynoResults[0][0].Title)
	assert.Equal(t, expectedDynoResults[0], *actualDynoResults[0][1].Title)
	assert.Equal(t, expectedDynoResults[1], *actualDynoResults[1][0].Title)
	assert.Equal(t, expectedDynoResults[1], *actualDynoResults[1][1].Title)
	assert.Equal(t, expectedDynoResults[2], *actualDynoResults[2][0].Title)
	assert.Equal(t, expectedDynoResults[2], *actualDynoResults[2][1].Title)
	assert.Equal(t, expectedDynoResults[3], *actualDynoResults[3][0].Title)
	assert.Equal(t, expectedDynoResults[3], *actualDynoResults[3][1].Title)
	assert.Equal(t, expectedDynoResults[4], *actualDynoResults[4][0].Title)
	assert.Equal(t, expectedDynoResults[4], *actualDynoResults[4][1].Title)
}

func TestCreateResultInvalidNoFile(t *testing.T) {
	location := ""
	testBugFile := ""
	var fuzzError []string
	assert.Panics(t, func() { parse.CreateResults(location, testBugFile, fuzzError) }, "Panics as there is no file")
}

func TestCreateResultsInvalidDynamicObjectChecker_1(t *testing.T) {
	// Setup
	expectedTitle := ""
	expectedMethod := ""
	expectedAcceptedResponse := ""
	expectedHost := ""
	expectedContentType := ""
	expectedRequest := ""
	expectedTimeDelay := ""
	expectedAsyncTime := ""
	expectedPreviousResponse := ""
	location := "../tests/bug_buckets/"
	testBugFile := "InvalidDynamicObjectChecker_20x_1.txt"
	fuzzError := strings.Split(testBugFile, "_")
	// Testing function
	actualResult := parse.CreateResults(location, testBugFile, fuzzError)
	expectedTitle = "InvalidDynamicObjectChecker Invalid 20x Response"
	assert.Equal(t, expectedTitle, string(*actualResult[0].Title))
	expectedMethod = "-> POST /api/blog/posts HTTP/1.1"
	assert.Equal(t, expectedMethod, string(*actualResult[0].Method))
	expectedAcceptedResponse = "Accept: application/json"
	assert.Equal(t, expectedAcceptedResponse, string(*actualResult[0].MethodInformation.AcceptedResponse))
	expectedHost = "Host: localhost:8888"
	assert.Equal(t, expectedHost, string(*actualResult[0].MethodInformation.Host))
	expectedContentType = "Content-Type: application/json"
	assert.Equal(t, expectedContentType, string(*actualResult[0].MethodInformation.ContentType))
	expectedRequest = "Request: {    \"id\":99,    \"body\":\"my first blog post\"}"
	assert.Equal(t, expectedRequest, string(*actualResult[0].MethodInformation.Request))
	expectedTimeDelay = "! producer_timing_delay 0"
	assert.Equal(t, expectedTimeDelay, string(*actualResult[0].TimeDelay))
	expectedAsyncTime = "! max_async_wait_time 20"
	assert.Equal(t, expectedAsyncTime, string(*actualResult[0].AsyncTime))
	expectedPreviousResponse = "PREVIOUS RESPONSE: 'HTTP/1.1 201 Created response:{\"id\":10,\"body\":\"my first blog post\"}'"
	assert.Equal(t, expectedPreviousResponse, string(*actualResult[0].PreviousResponse))
	expectedMethod = "-> GET /api/blog/posts/10?injected_query_string=123 HTTP/1.1"
	assert.Equal(t, expectedMethod, string(*actualResult[1].Method))
	expectedAcceptedResponse = "Accept: application/json"
	assert.Equal(t, expectedAcceptedResponse, string(*actualResult[1].MethodInformation.AcceptedResponse))
	expectedHost = "Host: localhost:8888"
	assert.Equal(t, expectedHost, string(*actualResult[1].MethodInformation.Host))
	expectedTimeDelay = "! producer_timing_delay 0"
	assert.Equal(t, expectedTimeDelay, string(*actualResult[1].TimeDelay))
	expectedAsyncTime = "! max_async_wait_time 0"
	assert.Equal(t, expectedAsyncTime, string(*actualResult[1].AsyncTime))
	expectedPreviousResponse = "PREVIOUS RESPONSE: 'HTTP/1.1 200 OK response:{\"id\":10,\"body\":\"my first blog post\"}'"
	assert.Equal(t, expectedPreviousResponse, string(*actualResult[1].PreviousResponse))
	expectedEndpoint := "/api/blog/posts/10?injected_query_string=123"
	assert.Equal(t, expectedEndpoint, string(*actualResult[1].Endpoint))
}
func TestCreateResultsValidInvalidDynamicObjectChecker_2(t *testing.T) {
	// Setup
	expectedTitle := ""
	expectedMethod := ""
	expectedAcceptedResponse := ""
	expectedHost := ""
	expectedContentType := ""
	expectedRequest := ""
	expectedTimeDelay := ""
	expectedAsyncTime := ""
	expectedPreviousResponse := ""
	location := "../tests/bug_buckets/"
	testBugFile := "InvalidDynamicObjectChecker_20x_2.txt"
	fuzzError := strings.Split(testBugFile, "_")
	// Testing function
	actualResult := parse.CreateResults(location, testBugFile, fuzzError)
	expectedTitle = "InvalidDynamicObjectChecker Invalid 20x Response"
	assert.Equal(t, expectedTitle, string(*actualResult[0].Title))
	expectedMethod = "-> POST /api/blog/posts HTTP/1.1"
	assert.Equal(t, expectedMethod, string(*actualResult[0].Method))
	expectedAcceptedResponse = "Accept: application/json"
	assert.Equal(t, expectedAcceptedResponse, string(*actualResult[0].MethodInformation.AcceptedResponse))
	expectedHost = "Host: localhost:8888"
	assert.Equal(t, expectedHost, string(*actualResult[0].MethodInformation.Host))
	expectedContentType = "Content-Type: application/json"
	assert.Equal(t, expectedContentType, string(*actualResult[0].MethodInformation.ContentType))
	expectedRequest = "Request: {    \"id\":99,    \"body\":\"my first blog post\"}"
	assert.Equal(t, expectedRequest, string(*actualResult[0].MethodInformation.Request))
	expectedTimeDelay = "! producer_timing_delay 0"
	assert.Equal(t, expectedTimeDelay, string(*actualResult[0].TimeDelay))
	expectedAsyncTime = "! max_async_wait_time 20"
	assert.Equal(t, expectedAsyncTime, string(*actualResult[0].AsyncTime))
	expectedPreviousResponse = "PREVIOUS RESPONSE: 'HTTP/1.1 201 Created response:{\"id\":13,\"body\":\"my first blog post\"}'"
	assert.Equal(t, expectedPreviousResponse, string(*actualResult[0].PreviousResponse))
	expectedMethod = "-> PUT /api/blog/posts/13?injected_query_string=123 HTTP/1.1"
	assert.Equal(t, expectedMethod, string(*actualResult[1].Method))
	expectedAcceptedResponse = "Accept: application/json"
	assert.Equal(t, expectedAcceptedResponse, string(*actualResult[1].MethodInformation.AcceptedResponse))
	expectedHost = "Host: localhost:8888"
	assert.Equal(t, expectedHost, string(*actualResult[1].MethodInformation.Host))
	expectedContentType = "Content-Type: application/json"
	assert.Equal(t, expectedContentType, string(*actualResult[1].MethodInformation.ContentType))
	expectedRequest = "Request: {    \"id\":13,    \"body\":\"my first blog post?injected_query_string=123\",    \"checksum\":\"abcde\"}"
	assert.Equal(t, expectedRequest, string(*actualResult[1].MethodInformation.Request))
	expectedTimeDelay = "! producer_timing_delay 0"
	assert.Equal(t, expectedTimeDelay, string(*actualResult[1].TimeDelay))
	expectedAsyncTime = "! max_async_wait_time 0"
	assert.Equal(t, expectedAsyncTime, string(*actualResult[1].AsyncTime))
	expectedPreviousResponse = "PREVIOUS RESPONSE: 'HTTP/1.1 204 No Content"
	assert.Equal(t, expectedPreviousResponse, string(*actualResult[1].PreviousResponse))
	expectedEndpoint := "/api/blog/posts/13?injected_query_string=123"
	assert.Equal(t, expectedEndpoint, string(*actualResult[1].Endpoint))
}

func TestCreateResultsValidPayloadBodyChecker_1(t *testing.T) {
	// Setup
	expectedTitle := ""
	expectedMethod := ""
	expectedAcceptedResponse := ""
	expectedHost := ""
	expectedContentType := ""
	expectedRequest := ""
	expectedTimeDelay := ""
	expectedAsyncTime := ""
	expectedPreviousResponse := ""
	location := "../tests/bug_buckets/"
	testBugFile := "PayloadBodyChecker_500_1.txt"
	fuzzError := strings.Split(testBugFile, "_")
	// Testing function
	actualResult := parse.CreateResults(location, testBugFile, fuzzError)
	expectedTitle = "PayloadBodyChecker Invalid 500 Response"
	assert.Equal(t, expectedTitle, string(*actualResult[0].Title))
	expectedMethod = "-> POST /api/blog/posts HTTP/1.1"
	assert.Equal(t, expectedMethod, string(*actualResult[0].Method))
	expectedAcceptedResponse = "Accept: application/json"
	assert.Equal(t, expectedAcceptedResponse, string(*actualResult[0].MethodInformation.AcceptedResponse))
	expectedHost = "Host: localhost:8888"
	assert.Equal(t, expectedHost, string(*actualResult[0].MethodInformation.Host))
	expectedContentType = "Content-Type: application/json"
	assert.Equal(t, expectedContentType, string(*actualResult[0].MethodInformation.ContentType))
	expectedRequest = "Request: {    \"id\":99,    \"body\":\"my first blog post\"}"
	assert.Equal(t, expectedRequest, string(*actualResult[0].MethodInformation.Request))
	expectedTimeDelay = "! producer_timing_delay 0"
	assert.Equal(t, expectedTimeDelay, string(*actualResult[0].TimeDelay))
	expectedAsyncTime = "! max_async_wait_time 0"
	assert.Equal(t, expectedAsyncTime, string(*actualResult[0].AsyncTime))
	expectedPreviousResponse = "PREVIOUS RESPONSE: 'HTTP/1.1 201 Created response:{\"id\":12,\"body\":\"my first blog post\"}'"
	assert.Equal(t, expectedPreviousResponse, string(*actualResult[0].PreviousResponse))
	expectedMethod = "-> PUT /api/blog/posts/14 HTTP/1.1"
	assert.Equal(t, expectedMethod, string(*actualResult[1].Method))
	expectedAcceptedResponse = "Accept: application/json"
	assert.Equal(t, expectedAcceptedResponse, string(*actualResult[1].MethodInformation.AcceptedResponse))
	expectedHost = "Host: localhost:8888"
	assert.Equal(t, expectedHost, string(*actualResult[1].MethodInformation.Host))
	expectedContentType = "Content-Type: application/json"
	assert.Equal(t, expectedContentType, string(*actualResult[1].MethodInformation.ContentType))
	expectedRequest = "Request: {\"body\":\"my first blog post\"}"
	assert.Equal(t, expectedRequest, string(*actualResult[1].MethodInformation.Request))
	expectedTimeDelay = "! producer_timing_delay 0"
	assert.Equal(t, expectedTimeDelay, string(*actualResult[1].TimeDelay))
	expectedAsyncTime = "! max_async_wait_time 0"
	assert.Equal(t, expectedAsyncTime, string(*actualResult[1].AsyncTime))
	expectedPreviousResponse = "PREVIOUS RESPONSE: 'HTTP/1.1 500 Internal Server Error response:{\"detail\":\"ID was not specified.\"}'"
	assert.Equal(t, expectedPreviousResponse, string(*actualResult[1].PreviousResponse))
	expectedEndpoint := "/api/blog/posts/14"
	assert.Equal(t, expectedEndpoint, string(*actualResult[1].Endpoint))
}

func TestCreateResultsValidPayloadBodyChecker_2(t *testing.T) {
	// Setup
	expectedTitle := ""
	expectedMethod := ""
	expectedAcceptedResponse := ""
	expectedHost := ""
	expectedContentType := ""
	expectedRequest := ""
	expectedTimeDelay := ""
	expectedAsyncTime := ""
	expectedPreviousResponse := ""
	location := "../tests/bug_buckets/"
	testBugFile := "PayloadBodyChecker_500_2.txt"
	fuzzError := strings.Split(testBugFile, "_")
	// Testing function
	actualResult := parse.CreateResults(location, testBugFile, fuzzError)
	expectedTitle = "PayloadBodyChecker Invalid 500 Response"
	assert.Equal(t, expectedTitle, string(*actualResult[0].Title))
	expectedMethod = "-> POST /api/blog/posts HTTP/1.1"
	assert.Equal(t, expectedMethod, string(*actualResult[0].Method))
	expectedAcceptedResponse = "Accept: application/json"
	assert.Equal(t, expectedAcceptedResponse, string(*actualResult[0].MethodInformation.AcceptedResponse))
	expectedHost = "Host: localhost:8888"
	assert.Equal(t, expectedHost, string(*actualResult[0].MethodInformation.Host))
	expectedContentType = "Content-Type: application/json"
	assert.Equal(t, expectedContentType, string(*actualResult[0].MethodInformation.ContentType))
	expectedRequest = "Request: {    \"id\":99,    \"body\":\"my first blog post\"}"
	assert.Equal(t, expectedRequest, string(*actualResult[0].MethodInformation.Request))
	expectedTimeDelay = "! producer_timing_delay 0"
	assert.Equal(t, expectedTimeDelay, string(*actualResult[0].TimeDelay))
	expectedAsyncTime = "! max_async_wait_time 0"
	assert.Equal(t, expectedAsyncTime, string(*actualResult[0].AsyncTime))
	expectedPreviousResponse = "PREVIOUS RESPONSE: 'HTTP/1.1 201 Created response:{\"id\":12,\"body\":\"my first blog post\"}'"
	assert.Equal(t, expectedPreviousResponse, string(*actualResult[0].PreviousResponse))
	expectedMethod = "-> PUT /api/blog/posts/16 HTTP/1.1"
	assert.Equal(t, expectedMethod, string(*actualResult[1].Method))
	expectedAcceptedResponse = "Accept: application/json"
	assert.Equal(t, expectedAcceptedResponse, string(*actualResult[1].MethodInformation.AcceptedResponse))
	expectedHost = "Host: localhost:8888"
	assert.Equal(t, expectedHost, string(*actualResult[1].MethodInformation.Host))
	expectedContentType = "Content-Type: application/json"
	assert.Equal(t, expectedContentType, string(*actualResult[1].MethodInformation.ContentType))
	expectedRequest = "Request: {\"body\":0}"
	assert.Equal(t, expectedRequest, string(*actualResult[1].MethodInformation.Request))
	expectedTimeDelay = "! producer_timing_delay 0"
	assert.Equal(t, expectedTimeDelay, string(*actualResult[1].TimeDelay))
	expectedAsyncTime = "! max_async_wait_time 0"
	assert.Equal(t, expectedAsyncTime, string(*actualResult[1].AsyncTime))
	expectedPreviousResponse = "PREVIOUS RESPONSE: 'HTTP/1.1 500 Internal Server Error response:{\"detail\":\"ID was not specified.\"}'"
	assert.Equal(t, expectedPreviousResponse, string(*actualResult[1].PreviousResponse))
	expectedEndpoint := "/api/blog/posts/16"
	assert.Equal(t, expectedEndpoint, string(*actualResult[1].Endpoint))
}

func TestCreateResultsValidUseAfterFreeChecker(t *testing.T) {

	// Setup
	expectedTitle := ""
	expectedMethod := ""
	expectedAcceptedResponse := ""
	expectedHost := ""
	expectedContentType := ""
	expectedRequest := ""
	expectedTimeDelay := ""
	expectedAsyncTime := ""
	expectedPreviousResponse := ""
	location := "../tests/bug_buckets/"
	testBugFile := "UseAfterFreeChecker_20x_1.txt"
	fuzzError := strings.Split(testBugFile, "_")
	// Testing function
	actualResult := parse.CreateResults(location, testBugFile, fuzzError)
	expectedTitle = "UseAfterFreeChecker Invalid 20x Response"
	assert.Equal(t, expectedTitle, string(*actualResult[0].Title))
	expectedMethod = "-> POST /api/blog/posts HTTP/1.1"
	assert.Equal(t, expectedMethod, string(*actualResult[0].Method))
	expectedAcceptedResponse = "Accept: application/json"
	assert.Equal(t, expectedAcceptedResponse, string(*actualResult[0].MethodInformation.AcceptedResponse))
	expectedHost = "Host: localhost:8888"
	assert.Equal(t, expectedHost, string(*actualResult[0].MethodInformation.Host))
	expectedContentType = "Content-Type: application/json"
	assert.Equal(t, expectedContentType, string(*actualResult[0].MethodInformation.ContentType))
	expectedRequest = "Request: {    \"id\":99,    \"body\":\"my first blog post\"}"
	assert.Equal(t, expectedRequest, string(*actualResult[0].MethodInformation.Request))
	expectedTimeDelay = "! producer_timing_delay 0"
	assert.Equal(t, expectedTimeDelay, string(*actualResult[0].TimeDelay))
	expectedAsyncTime = "! max_async_wait_time 0"
	assert.Equal(t, expectedAsyncTime, string(*actualResult[0].AsyncTime))
	expectedPreviousResponse = "PREVIOUS RESPONSE: 'HTTP/1.1 201 Created response:{\"id\":20,\"body\":\"my first blog post\"}'"
	assert.Equal(t, expectedPreviousResponse, string(*actualResult[0].PreviousResponse))
	expectedMethod = "-> DELETE /api/blog/posts/20 HTTP/1.1"
	assert.Equal(t, expectedMethod, string(*actualResult[1].Method))
	expectedAcceptedResponse = "Accept: application/json"
	assert.Equal(t, expectedAcceptedResponse, string(*actualResult[1].MethodInformation.AcceptedResponse))
	expectedHost = "Host: localhost:8888"
	assert.Equal(t, expectedHost, string(*actualResult[1].MethodInformation.Host))
	expectedTimeDelay = "! producer_timing_delay 0"
	assert.Equal(t, expectedTimeDelay, string(*actualResult[1].TimeDelay))
	expectedAsyncTime = "! max_async_wait_time 20"
	assert.Equal(t, expectedAsyncTime, string(*actualResult[1].AsyncTime))
	expectedPreviousResponse = "PREVIOUS RESPONSE: 'HTTP/1.1 204 No Content"
	assert.Equal(t, expectedPreviousResponse, string(*actualResult[1].PreviousResponse))
	expectedMethod = "-> GET /api/blog/posts/20 HTTP/1.1"
	assert.Equal(t, expectedMethod, string(*actualResult[2].Method))
	expectedAcceptedResponse = "Accept: application/json"
	assert.Equal(t, expectedAcceptedResponse, string(*actualResult[2].MethodInformation.AcceptedResponse))
	expectedHost = "Host: localhost:8888"
	assert.Equal(t, expectedHost, string(*actualResult[2].MethodInformation.Host))
	expectedTimeDelay = "! producer_timing_delay 0"
	assert.Equal(t, expectedTimeDelay, string(*actualResult[2].TimeDelay))
	expectedAsyncTime = "! max_async_wait_time 0"
	assert.Equal(t, expectedAsyncTime, string(*actualResult[2].AsyncTime))
	expectedPreviousResponse = "PREVIOUS RESPONSE: 'HTTP/1.1 200 OK response:null'"
	assert.Equal(t, expectedPreviousResponse, string(*actualResult[2].PreviousResponse))
	expectedEndpoint := "/api/blog/posts/20"
	assert.Equal(t, expectedEndpoint, string(*actualResult[2].Endpoint))
}

func TestDYNODetailsValid(t *testing.T) {
	fuzzErrorList := [7]string{"InternalServerErrors",
		"UseAfterFreeChecker",
		"NameSpaceRuleChecker",
		"ResourceHierarchyChecker",
		"LeakageRuleChecker",
		"InvalidDynamicObjectChecker",
		"PayloadBodyChecker",
	}
	actualDetails := parse.AddDYNODetails(fuzzErrorList[0])
	assert.Equal(t, "\nDetails: '500 Internal Server' Errors and any other 5xx errors are detected.\n\nVisualizer: [DYNO](the web url)\n", actualDetails)
	actualDetails = parse.AddDYNODetails(fuzzErrorList[1])
	assert.Equal(t, "\nDetails: Detects that a deleted resource can still being accessed after deletion.\n\nVisualizer: [DYNO](the web url)\n", actualDetails)
	actualDetails = parse.AddDYNODetails(fuzzErrorList[2])
	assert.Equal(t, "\nDetails: Detects that an unauthorized user can access service resources.\n\nVisualizer: [DYNO](the web url)\n", actualDetails)
	actualDetails = parse.AddDYNODetails(fuzzErrorList[3])
	assert.Equal(t, "\nDetails: Detects that a child resource can be accessed from a non-parent resource.\n\nVisualizer: [DYNO](the web url)\n", actualDetails)
	actualDetails = parse.AddDYNODetails(fuzzErrorList[4])
	assert.Equal(t, "\nDetails: Detects that a failed resource creation leaks data in subsequent requests.\n\nVisualizer: [DYNO](the web url)\n", actualDetails)
	actualDetails = parse.AddDYNODetails(fuzzErrorList[5])
	assert.Equal(t, "\nDetails: Detects 500 errors or unexpected success status codes when invalid dynamic objects are sent in requests.\n\nVisualizer: [DYNO](the web url)\n", actualDetails)
	actualDetails = parse.AddDYNODetails(fuzzErrorList[6])
	assert.Equal(t, "\nDetails: Detects 500 errors when fuzzing the JSON bodies of requests.\n\nVisualizer: [DYNO](the web url)\n", actualDetails)
}

func TestDYNODetailsInvalid(t *testing.T) {
	actualDetails := parse.AddDYNODetails("")
	assert.Equal(t, "", actualDetails)
}
