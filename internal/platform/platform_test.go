package platform_test

import (
	"github.com/stretchr/testify/assert"
	"dyno/internal/platform"
	"dyno/internal/issue"
	"dyno/internal/result"
	"testing"
)

func TestFormatFuzzBody( t *testing.T) {
	title := "InvalidDynamicObjectChecker Invalid 20x Response"
	endpoint := "/api/blog/posts"
	method := "-> POST /api/blog/posts HTTP/1.1"
	acceptedResponse := "Accept: application/json"
	host := "Host: localhost:8888"
	contentType := "Content-Type: application/json"
	request := "Request: {    \"id\":99,    \"body\":\"my first blog post\"}"
	timeDelay := "! producer_timing_delay 0"
	asyncTime := "! max_async_wait_time 20"
	previousResponse := "PREVIOUS RESPONSE: 'HTTP/1.1 201 Created response:{\"id\":10,\"body\":\"my first blog post\"}'"
	errorType := "InvalidDynamicObjectChecker"
	methodInformation := result.DynoMethodInformation {
		AcceptedResponse: &acceptedResponse,
		Host: &host, 
		ContentType: &contentType, 
		Request: &request,
	}
	
	dynoResult := result.DynoResult{
		Title: &title, 
		Endpoint: &endpoint,
		Method: &method, 
		MethodInformation: &methodInformation, 
		TimeDelay: &timeDelay, 
		AsyncTime: &asyncTime, 
		PreviousResponse: &previousResponse, 
		ErrorType: &errorType,
	}

	issueTitle := "DYNO Fuzz: InvalidDynamicObjectChecker at Endpoint /api/blog/posts"
	details := "Details: Detects 500 errors or unexpected success status codes when invalid dynamic objects are sent in requests."
	visualizer := "Visualizer: [DYNO](the web url)"
	labels := []string{"bug"}
	assignee := "fishua"
	state := "state"
	milestone := 1
	issue := issue.DynoIssue{
		Title: &issueTitle,
		Details: &details, 
		Visualizer: &visualizer, 
		Body: &dynoResult,
		Labels: &labels, 
		Assignee: &assignee, 
		State: &state, 
		Milestone: &milestone, 
	}

	expectedBody := "# InvalidDynamicObjectChecker Invalid 20x Response\n" + 
									"\nDetails: Detects 500 errors or unexpected success status codes when invalid dynamic objects are sent in requests.\n" + 
									"\nVisualizer: [DYNO](the web url)\n" + 
									"\n-> POST /api/blog/posts HTTP/1.1\n\n" + 
									"- Accept: application/json\n" + 
									"- Host: localhost:8888\n" + 
									"- Content-Type: application/json\n" + 
									"- Request: {    \"id\":99,    \"body\":\"my first blog post\"}\n\n" + 
									"! producer_timing_delay 0\n" + 
									"! max_async_wait_time 20\n" + 
									"\nPREVIOUS RESPONSE: 'HTTP/1.1 201 Created response:{\"id\":10,\"body\":\"my first blog post\"}'\n"
								
	assert.Equal(t, expectedBody, *platform.FormatFuzzBody(&issue))

	methodInformation.ContentType = nil 
	expectedBody = "# InvalidDynamicObjectChecker Invalid 20x Response\n" + 
	"\nDetails: Detects 500 errors or unexpected success status codes when invalid dynamic objects are sent in requests.\n" + 
	"\nVisualizer: [DYNO](the web url)\n" + 
	"\n-> POST /api/blog/posts HTTP/1.1\n\n" + 
	"- Accept: application/json\n" + 
	"- Host: localhost:8888\n" +
	"- Request: {    \"id\":99,    \"body\":\"my first blog post\"}\n\n" + 
	"! producer_timing_delay 0\n" + 
	"! max_async_wait_time 20\n" + 
	"\nPREVIOUS RESPONSE: 'HTTP/1.1 201 Created response:{\"id\":10,\"body\":\"my first blog post\"}'\n"
	assert.Equal(t, expectedBody, *platform.FormatFuzzBody(&issue))
} 