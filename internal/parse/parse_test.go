package parse_test

//TODO Talk about what to do when the file is wrong

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"golambda/internal/parse"
)

func TestReadBugFileInvalidNoFile(t *testing.T) {
	location := ""
	testBugFile := ""
	body := ""
	assert.Panics(t, func() {ReadBugFile(location, testBugFile, body)}, "Panics as there is no file")
}

func TestReadBugFileValidInvalidDynamicObjectChecker_1(t *testing.T) {
	location := "tests/bug_buckets/"
	testBugFile := "InvalidDynamicObjectChecker_20x_1.txt"
	body := "# InvalidDynamicObjectChecker Invalid 200 Response"
	bodyCheck, endpointCheck := ReadBugFile(location, testBugFile, body)

	expectedBody := "# InvalidDynamicObjectChecker Invalid 200 Response\n" + 
									"-> POST /api/blog/posts HTTP/1.1\n\n" + 
									"- Accept: application/json\n" +
									"- Host: localhost:8888\n" +
									"- Content-Type: application/json\n" +
									"- Request: {    \"id\":99,    \"body\":\"my first blog post\"}\n\n" +
									"! producer_timing_delay 0\n" +
									"! max_async_wait_time 20\n\n" +
									"PREVIOUS RESPONSE: 'HTTP/1.1 201 Created request:{\"id\":10,\"body\":\"my first blog post\"}'\n\n" + 
									"-> GET /api/blog/posts/10?injected_query_string=123 HTTP/1.1\n\n" +
									"- Accept: application/json\n" +
									"- Host: localhost:8888\n\n" +
									"! producer_timing_delay 0\n" +
									"! max_async_wait_time 0\n\n" +
									"PREVIOUS RESPONSE: 'HTTP/1.1 200 OK request:{\"id\":10,\"body\":\"my first blog post\"}'\n"
	expectedEndpoint := "/api/blog/posts/10?injected_query_string=123"
	assert.Equal(t, expectedEndpoint, string(endpointCheck))
	assert.Equal(t, expectedBody, string(bodyCheck))
}
func TestReadBugFileValidInvalidDynamicObjectChecker_2(t *testing.T) {
	location := "tests/bug_buckets/"
	testBugFile := "InvalidDynamicObjectChecker_20x_2.txt"
	body := "# InvalidDynamicObjectChecker Invalid 200 Response"
	bodyCheck, endpointCheck := ReadBugFile(location, testBugFile, body)

	expectedBody := "# InvalidDynamicObjectChecker Invalid 200 Response\n" + 
									"-> POST /api/blog/posts HTTP/1.1\n\n" + 
									"- Accept: application/json\n" +
									"- Host: localhost:8888\n" +
									"- Content-Type: application/json\n" +
									"- Request: {    \"id\":99,    \"body\":\"my first blog post\"}\n\n" +
									"! producer_timing_delay 0\n" +
									"! max_async_wait_time 20\n\n" +
									"PREVIOUS RESPONSE: 'HTTP/1.1 201 Created request:{\"id\":13,\"body\":\"my first blog post\"}'\n\n" + 
									"-> PUT /api/blog/posts/13?injected_query_string=123 HTTP/1.1\n\n" +
									"- Accept: application/json\n" +
									"- Host: localhost:8888\n\n" + 
									"- Content-Type: application/json\n" +
									"- Request: {    \"id\":13,    \"body\":\"my first blog post?injected_query_string=123\",    \"checksum\":\"abcde\"}\n\n" +
									"! producer_timing_delay 0\n" +
									"! max_async_wait_time 0\n\n" +
									"PREVIOUS RESPONSE: 'HTTP/1.1 204 No Content\n"
	expectedEndpoint := "/api/blog/posts/13?injected_query_string=123"
	assert.Equal(t, expectedEndpoint, string(endpointCheck))
	assert.Equal(t, expectedBody, string(bodyCheck))
}

func TestReadBugFileValidPayloadBodyChecker_1(t *testing.T) {
	location := "tests/bug_buckets/"
	testBugFile := "PayloadBodyChecker_500_1.txt"
	body := "# PayloadBodyChecker Invalid 500 Response"
	bodyCheck, endpointCheck := ReadBugFile(location, testBugFile, body)

	expectedBody := "# PayloadBodyChecker Invalid 500 Response\n" + 
									"-> POST /api/blog/posts HTTP/1.1\n\n" + 
									"- Accept: application/json\n" +
									"- Host: localhost:8888\n" +
									"- Content-Type: application/json\n" +
									"- Request: {    \"id\":99,    \"body\":\"my first blog post\"}\n\n" +
									"! producer_timing_delay 0\n" +
									"! max_async_wait_time 0\n\n" +
									"PREVIOUS RESPONSE: 'HTTP/1.1 201 Created request:{\"id\":12,\"body\":\"my first blog post\"}'\n\n" + 
									"-> PUT /api/blog/posts/14 HTTP/1.1\n\n" +
									"- Accept: application/json\n" +
									"- Host: localhost:8888\n\n" +
									"- Content-Type: application/json\n" +
									"- Request: {\"body\":\"my first blog post\"}\n\n" +
									"! producer_timing_delay 0\n" +
									"! max_async_wait_time 0\n\n" +
									"PREVIOUS RESPONSE: 'HTTP/1.1 500 Internal Server Error request:{\"detail\":\"ID was not specified.\"}'\n"
	expectedEndpoint := "/api/blog/posts"
	assert.Equal(t, expectedEndpoint, string(endpointCheck))
	assert.Equal(t, expectedBody, string(bodyCheck))
}

func TestReadBugFileValidPayloadBodyChecker_2(t *testing.T) {
	location := "tests/bug_buckets/"
	testBugFile := "PayloadBodyChecker_500_2.txt"
	body := "# PayloadBodyChecker Invalid 500 Response"
	bodyCheck, endpointCheck := ReadBugFile(location, testBugFile, body)

	expectedBody := "# PayloadBodyChecker Invalid 500 Response\n" + 
									"-> POST /api/blog/posts HTTP/1.1\n\n" + 
									"- Accept: application/json\n" +
									"- Host: localhost:8888\n" +
									"- Content-Type: application/json\n" +
									"- Request: {    \"id\":99,    \"body\":\"my first blog post\"}\n\n" +
									"! producer_timing_delay 0\n" +
									"! max_async_wait_time 0\n\n" +
									"PREVIOUS RESPONSE: 'HTTP/1.1 201 Created request:{\"id\":12,\"body\":\"my first blog post\"}'\n\n" + 
									"-> PUT /api/blog/posts/16 HTTP/1.1\n\n" +
									"- Accept: application/json\n" +
									"- Host: localhost:8888\n\n" +
									"- Content-Type: application/json\n" +
									"- Request: {\"body\":0}\n\n" +
									"! producer_timing_delay 0\n" +
									"! max_async_wait_time 0\n\n" +
									"PREVIOUS RESPONSE: 'HTTP/1.1 500 Internal Server Error request:{\"detail\":\"ID was not specified.\"}'\n"
	expectedEndpoint := "/api/blog/posts/16"
	assert.Equal(t, expectedEndpoint, string(endpointCheck))
	assert.Equal(t, expectedBody, string(bodyCheck))
}

func TestReadBugFileValidUseAfterFreeChecker(t *testing.T) {
	location := "tests/bug_buckets/"
	testBugFile := "UseAfterFreeChecker_20x_1.txt"
	body := "# UseAfterFreeChecker Invalid 200 Response"
	bodyCheck, endpointCheck := ReadBugFile(location, testBugFile, body)

	expectedBody := "# UseAfterFreeChecker Invalid 200 Response\n" + 
									"-> POST /api/blog/posts HTTP/1.1\n\n" + 
									"- Accept: application/json\n" +
									"- Host: localhost:8888\n" +
									"- Content-Type: application/json\n" +
									"- Request: {    \"id\":99,    \"body\":\"my first blog post\"}\n\n" +
									"! producer_timing_delay 0\n" +
									"! max_async_wait_time 0\n\n" +
									"PREVIOUS RESPONSE: 'HTTP/1.1 201 Created request:{\"id\":20,\"body\":\"my first blog post\"}'\n\n" + 
									"-> DELETE /api/blog/posts/20 HTTP/1.1\n\n" +
									"- Accept: application/json\n" +
									"- Host: localhost:8888\n\n" +
									"! producer_timing_delay 0\n" +
									"! max_async_wait_time 20\n\n" +
									"PREVIOUS RESPONSE: 'HTTP/1.1 204 No Content\n" +
									"-> GET /api/blog/posts/20 HTTP/1.1\n\n" +
									"- Accept: application/json\n" +
									"- Host: localhost:8888\n\n" +
									"! producer_timing_delay 0\n" +
									"! max_async_wait_time 0\n\n" +
									"PREVIOUS RESPONSE: 'HTTP/1.1 200 OK request:null'\n"
	expectedEndpoint := "/api/blog/posts/20"
	assert.Equal(t, expectedEndpoint, string(endpointCheck))
	assert.Equal(t, expectedBody, string(bodyCheck))
}

func TestFuzzBugCheckInvalid(t *testing.T) {
	fuzzError := "InternalServerErrors"
	newIssueRequest := FuzzBugCheck(fuzzError, "body", "/endpoint", nil, nil, nil)
	assert.Equal(t, "DYNO Fuzz: InternalServerErrors at Endpoint /endpoint", *newIssueRequest.Title)
	assert.Equal(t, "body", *newIssueRequest.Body)
	assert.Equal(t, nil, newIssueRequest.Assignee)
	assert.Equal(t, nil, newIssueRequest.State)
	assert.Equal(t, nil, newIssueRequest.Milestone)
} 

func TestFuzzBugCheckValid(t *testing.T) {
	body := "body"
	endpoint := "/endpoint"
	assignee := "fishua"
	state := "state"
	milestone := 1
	

	fuzzErrorList := [7]string{"InternalServerErrors", 
								"UseAfterFreeChecker",
								"NameSpaceRuleChecker",
								"ResourceHierarchyChecker",
								"LeakageRuleChecker",
								"InvalidDynamicObjectChecker",
								"PayloadBodyChecker",
							}
	for _, fuzzError := range fuzzErrorList {
		if fuzzError == "InternalServerErrors" {
			newIssueRequest := FuzzBugCheck(fuzzError, body, endpoint, &assignee, &state, &milestone)
			assert.Equal(t, "DYNO Fuzz: InternalServerErrors at Endpoint /endpoint", *newIssueRequest.Title)
			assert.Equal(t, "body", *newIssueRequest.Body)
			assert.Equal(t, "fishua", *newIssueRequest.Assignee)
			assert.Equal(t, "state", *newIssueRequest.State)
			assert.Equal(t, 1, *newIssueRequest.Milestone)
		}
	
		if fuzzError == "UseAfterFreeChecker" {
			newIssueRequest := FuzzBugCheck(fuzzError, body, endpoint, &assignee, &state, &milestone)
			assert.Equal(t, "DYNO Fuzz: UseAfterFreeChecker at Endpoint /endpoint", *newIssueRequest.Title)
			assert.Equal(t, "body", *newIssueRequest.Body)
			assert.Equal(t, "fishua", *newIssueRequest.Assignee)
			assert.Equal(t, "state", *newIssueRequest.State)
			assert.Equal(t, 1, *newIssueRequest.Milestone)
		}
	
		if fuzzError == "NameSpaceRuleChecker" {
			newIssueRequest := FuzzBugCheck(fuzzError, body, endpoint, &assignee, &state, &milestone)
			assert.Equal(t, "DYNO Fuzz: NameSpaceRuleChecker at Endpoint /endpoint", *newIssueRequest.Title)
			assert.Equal(t, "body", *newIssueRequest.Body)
			assert.Equal(t, "fishua", *newIssueRequest.Assignee)
			assert.Equal(t, "state", *newIssueRequest.State)
			assert.Equal(t, 1, *newIssueRequest.Milestone)
		}
	
		if fuzzError == "ResourceHierarchyChecker" {
			newIssueRequest := FuzzBugCheck(fuzzError, body, endpoint, &assignee, &state, &milestone)
			assert.Equal(t, "DYNO Fuzz: ResourceHierarchyChecker at Endpoint /endpoint", *newIssueRequest.Title)
			assert.Equal(t, "body", *newIssueRequest.Body)
			assert.Equal(t, "fishua", *newIssueRequest.Assignee)
			assert.Equal(t, "state", *newIssueRequest.State)
			assert.Equal(t, 1, *newIssueRequest.Milestone)
		}
	
		if fuzzError == "LeakageRuleChecker" {
			newIssueRequest := FuzzBugCheck(fuzzError, body, endpoint, &assignee, &state, &milestone)
			assert.Equal(t, "DYNO Fuzz: LeakageRuleChecker at Endpoint /endpoint", *newIssueRequest.Title)
			assert.Equal(t, "body", *newIssueRequest.Body)
			assert.Equal(t, "fishua", *newIssueRequest.Assignee)
			assert.Equal(t, "state", *newIssueRequest.State)
			assert.Equal(t, 1, *newIssueRequest.Milestone)
		}
	
		if fuzzError == "InvalidDynamicObjectChecker" {
			newIssueRequest := FuzzBugCheck(fuzzError, body, endpoint, &assignee, &state, &milestone)
			assert.Equal(t, "DYNO Fuzz: InvalidDynamicObjectChecker at Endpoint /endpoint", *newIssueRequest.Title)
			assert.Equal(t, "body", *newIssueRequest.Body)
			assert.Equal(t, "fishua", *newIssueRequest.Assignee)
			assert.Equal(t, "state", *newIssueRequest.State)
			assert.Equal(t, 1, *newIssueRequest.Milestone)
		}
	
		if fuzzError == "PayloadBodyChecker" {
			newIssueRequest := FuzzBugCheck(fuzzError, body, endpoint, &assignee, &state, &milestone)
			assert.Equal(t, "DYNO Fuzz: PayloadBodyChecker at Endpoint /endpoint", *newIssueRequest.Title)
			assert.Equal(t, "body", *newIssueRequest.Body)
			assert.Equal(t, "fishua", *newIssueRequest.Assignee)
			assert.Equal(t, "state", *newIssueRequest.State)
			assert.Equal(t, 1, *newIssueRequest.Milestone)
		}
	}
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
	actualDetails := ""
	for _, fuzzError := range fuzzErrorList {
		actualDetails = AddDYNODetails(fuzzError)
		if fuzzError == "InternalServerErrors" {
			assert.Equal(t, "\nDetails: '500 Internal Server' Errors and any other 5xx errors are detected.\n\nVisualizer: [DYNO](the web url)\n", actualDetails)
		}
	
		if fuzzError == "UseAfterFreeChecker" {
			assert.Equal(t, "\nDetails: Detects that a deleted resource can still being accessed after deletion.\n\nVisualizer: [DYNO](the web url)\n", actualDetails)
		}
	
		if fuzzError == "NameSpaceRuleChecker" {
			assert.Equal(t, "\nDetails: Detects that an unauthorized user can access service resources.\n\nVisualizer: [DYNO](the web url)\n", actualDetails)
		}
	
		if fuzzError == "ResourceHierarchyChecker" {
			assert.Equal(t, "\nDetails: Detects that a child resource can be accessed from a non-parent resource.\n\nVisualizer: [DYNO](the web url)\n", actualDetails)
		}
	
		if fuzzError == "LeakageRuleChecker" {
			assert.Equal(t, "\nDetails: Detects that a failed resource creation leaks data in subsequent requests.\n\nVisualizer: [DYNO](the web url)\n", actualDetails)
		}
	
		if fuzzError == "InvalidDynamicObjectChecker" {
			assert.Equal(t, "\nDetails: Detects 500 errors or unexpected success status codes when invalid dynamic objects are sent in requests.\n\nVisualizer: [DYNO](the web url)\n", actualDetails)
		}
	
		if fuzzError == "PayloadBodyChecker" {
			assert.Equal(t, "\nDetails: Detects 500 errors when fuzzing the JSON bodies of requests.\n\nVisualizer: [DYNO](the web url)\n", actualDetails)
		}
	}
}

func TestDYNODetailsInvalid(t *testing.T) {
	actualDetails := AddDYNODetails("")
	assert.Equal(t, "", actualDetails)

}