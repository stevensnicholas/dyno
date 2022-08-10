package issue_test

import (
	"dyno/internal/issue"
	"dyno/internal/result"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateIssueIntegratedTestInvalid(t *testing.T) {
	dynoResults := []result.DynoResult{}
	dynoIssues := issue.CreateIssues("RESTler", dynoResults)
	assert.Equal(t, 0, len(dynoIssues))
}

func TestCreateIssueIntegratedTestFuzzerInvalid(t *testing.T) {
	title := "title"
	endpoint := "/endpoint"
	method := "method"
	httpMethod := "METHOD"
	fuzzErrorList := [1]string{
		"InternalServerErrors",
	}
	dynoResults := []result.DynoResult{
		{
			Title:      &title,
			Endpoint:   &endpoint,
			Method:     &method,
			HTTPMethod: &httpMethod,
			ErrorType:  &fuzzErrorList[0],
		},
	}
	assert.Panics(t, func() { issue.CreateIssues("", dynoResults) })

}
func TestCreateIssueIntegratedTestValid(t *testing.T) {
	title := "title"
	endpoint := "/endpoint"
	method := "method"
	httpMethod := "METHOD"
	fuzzErrorList := [7]string{
		"InternalServerErrors",
		"UseAfterFreeChecker",
		"NameSpaceRuleChecker",
		"ResourceHierarchyChecker",
		"LeakageRuleChecker",
		"InvalidDynamicObjectChecker",
		"PayloadBodyChecker",
	}
	dynoResults := []result.DynoResult{
		{
			Title:      &title,
			Endpoint:   &endpoint,
			Method:     &method,
			HTTPMethod: &httpMethod,
			ErrorType:  &fuzzErrorList[0],
		},
		{
			Title:      &title,
			Endpoint:   &endpoint,
			Method:     &method,
			HTTPMethod: &httpMethod,
			ErrorType:  &fuzzErrorList[1],
		},
		{
			Title:      &title,
			Endpoint:   &endpoint,
			Method:     &method,
			HTTPMethod: &httpMethod,
			ErrorType:  &fuzzErrorList[2],
		},
		{
			Title:      &title,
			Endpoint:   &endpoint,
			Method:     &method,
			HTTPMethod: &httpMethod,
			ErrorType:  &fuzzErrorList[3],
		},
		{
			Title:      &title,
			Endpoint:   &endpoint,
			Method:     &method,
			HTTPMethod: &httpMethod,
			ErrorType:  &fuzzErrorList[4],
		},
		{
			Title:      &title,
			Endpoint:   &endpoint,
			Method:     &method,
			HTTPMethod: &httpMethod,
			ErrorType:  &fuzzErrorList[5],
		},
		{
			Title:      &title,
			Endpoint:   &endpoint,
			Method:     &method,
			HTTPMethod: &httpMethod,
			ErrorType:  &fuzzErrorList[6],
		},
	}

	dynoIssues := issue.CreateIssues("RESTler", dynoResults)
	assert.Equal(t, "DYNO Fuzz: InternalServerErrors at Endpoint /endpoint using METHOD Method", *dynoIssues[0].Title)
	assert.Equal(t, "Details: '500 Internal Server' Errors and any other 5xx errors are detected.", *dynoIssues[0].Details)
	assert.Equal(t, "Visualizer: [DYNO](the web url)", *dynoIssues[0].Visualizer)
	assert.Equal(t, []string{"Bug", "Medium"}, *dynoIssues[0].Labels)

	assert.Equal(t, "DYNO Fuzz: UseAfterFreeChecker at Endpoint /endpoint using METHOD Method", *dynoIssues[1].Title)
	assert.Equal(t, "Details: Detects that a deleted resource can still being accessed after deletion.", *dynoIssues[1].Details)
	assert.Equal(t, "Visualizer: [DYNO](the web url)", *dynoIssues[1].Visualizer)
	assert.Equal(t, []string{"Bug", "Medium"}, *dynoIssues[1].Labels)

	assert.Equal(t, "DYNO Fuzz: NameSpaceRuleChecker at Endpoint /endpoint using METHOD Method", *dynoIssues[2].Title)
	assert.Equal(t, "Details: Detects that an unauthorized user can access service resources.", *dynoIssues[2].Details)
	assert.Equal(t, "Visualizer: [DYNO](the web url)", *dynoIssues[2].Visualizer)
	assert.Equal(t, []string{"Bug", "High"}, *dynoIssues[2].Labels)

	assert.Equal(t, "DYNO Fuzz: ResourceHierarchyChecker at Endpoint /endpoint using METHOD Method", *dynoIssues[3].Title)
	assert.Equal(t, "Details: Detects that a child resource can be accessed from a non-parent resource.", *dynoIssues[3].Details)
	assert.Equal(t, "Visualizer: [DYNO](the web url)", *dynoIssues[3].Visualizer)
	assert.Equal(t, []string{"Bug", "High"}, *dynoIssues[3].Labels)

	assert.Equal(t, "DYNO Fuzz: LeakageRuleChecker at Endpoint /endpoint using METHOD Method", *dynoIssues[4].Title)
	assert.Equal(t, "Details: Detects that a failed resource creation leaks data in subsequent requests.", *dynoIssues[4].Details)
	assert.Equal(t, "Visualizer: [DYNO](the web url)", *dynoIssues[4].Visualizer)
	assert.Equal(t, []string{"Bug", "High"}, *dynoIssues[4].Labels)

	assert.Equal(t, "DYNO Fuzz: InvalidDynamicObjectChecker at Endpoint /endpoint using METHOD Method", *dynoIssues[5].Title)
	assert.Equal(t, "Details: Detects 500 errors or unexpected success status codes when invalid dynamic objects are sent in requests.", *dynoIssues[5].Details)
	assert.Equal(t, "Visualizer: [DYNO](the web url)", *dynoIssues[5].Visualizer)
	assert.Equal(t, []string{"Bug", "Medium"}, *dynoIssues[5].Labels)

	assert.Equal(t, "DYNO Fuzz: PayloadBodyChecker at Endpoint /endpoint using METHOD Method", *dynoIssues[6].Title)
	assert.Equal(t, "Details: Detects 500 errors when fuzzing the JSON bodies of requests.", *dynoIssues[6].Details)
	assert.Equal(t, "Visualizer: [DYNO](the web url)", *dynoIssues[6].Visualizer)
	assert.Equal(t, []string{"Bug", "Low"}, *dynoIssues[6].Labels)
}
