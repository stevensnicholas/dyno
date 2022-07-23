package issue_test

import (
	"golambda/internal/result"
	"golambda/internal/issue"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateIssueIntegratedTestInvalid(t *testing.T) {
	dynoResults := []result.DynoResult{}
	dynoIssues := issue.CreateIssues(dynoResults)
	assert.Equal(t, 0, len(dynoIssues))
}

func TestCreateIssueIntegratedTestValid(t *testing.T) {
	title := "title"
	endpoint := "/endpoint"
	method := "method"
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
		result.DynoResult{
			Title:    &title,
			Endpoint:	&endpoint, 
			Method:		&method, 
			ErrorType: &fuzzErrorList[0], 	
		}, 
		result.DynoResult{
			Title:    &title,
			Endpoint:	&endpoint, 
			Method:		&method, 
			ErrorType: &fuzzErrorList[1], 	
		}, 
		result.DynoResult{
			Title:    &title,
			Endpoint:	&endpoint, 
			Method:		&method, 
			ErrorType: &fuzzErrorList[2], 	
		}, 
		result.DynoResult{
			Title:    &title,
			Endpoint:	&endpoint, 
			Method:		&method, 
			ErrorType: &fuzzErrorList[3], 	
		}, 
		result.DynoResult{
			Title:    &title,
			Endpoint:	&endpoint, 
			Method:		&method, 
			ErrorType: &fuzzErrorList[4], 	
		}, 
		result.DynoResult{
			Title:    &title,
			Endpoint:	&endpoint, 
			Method:		&method, 
			ErrorType: &fuzzErrorList[5], 	
		}, 
		result.DynoResult{
			Title:    &title,
			Endpoint:	&endpoint, 
			Method:		&method, 
			ErrorType: &fuzzErrorList[6], 	
		}, 
	}

	dynoIssues := issue.CreateIssues(dynoResults)
	assert.Equal(t, "DYNO Fuzz: InternalServerErrors at Endpoint /endpoint", *dynoIssues[0].Title)
	assert.Equal(t, "\nDetails: '500 Internal Server' Errors and any other 5xx errors are detected.\n", *dynoIssues[0].Details)
	assert.Equal(t, "\nVisualizer: [DYNO](the web url)\n", *dynoIssues[0].Visualizer)
	assert.Equal(t, []string{"bug"}, *dynoIssues[0].Labels)

	assert.Equal(t, "DYNO Fuzz: UseAfterFreeChecker at Endpoint /endpoint", *dynoIssues[1].Title)
	assert.Equal(t, "\nDetails: Detects that a deleted resource can still being accessed after deletion.\n", *dynoIssues[1].Details)
	assert.Equal(t, "\nVisualizer: [DYNO](the web url)\n", *dynoIssues[1].Visualizer)
	assert.Equal(t, []string{"bug"}, *dynoIssues[1].Labels)

	assert.Equal(t, "DYNO Fuzz: NameSpaceRuleChecker at Endpoint /endpoint", *dynoIssues[2].Title)
	assert.Equal(t, "\nDetails: Detects that an unauthorized user can access service resources.\n", *dynoIssues[2].Details)
	assert.Equal(t, "\nVisualizer: [DYNO](the web url)\n", *dynoIssues[2].Visualizer)
	assert.Equal(t, []string{"bug"}, *dynoIssues[2].Labels)

	assert.Equal(t, "DYNO Fuzz: ResourceHierarchyChecker at Endpoint /endpoint", *dynoIssues[3].Title)
	assert.Equal(t, "\nDetails: Detects that a child resource can be accessed from a non-parent resource.\n", *dynoIssues[3].Details)
	assert.Equal(t, "\nVisualizer: [DYNO](the web url)\n", *dynoIssues[3].Visualizer)
	assert.Equal(t, []string{"bug"}, *dynoIssues[3].Labels)

	assert.Equal(t, "DYNO Fuzz: LeakageRuleChecker at Endpoint /endpoint", *dynoIssues[4].Title)
	assert.Equal(t, "\nDetails: Detects that a failed resource creation leaks data in subsequent requests.\n", *dynoIssues[4].Details)
	assert.Equal(t, "\nVisualizer: [DYNO](the web url)\n", *dynoIssues[4].Visualizer)
	assert.Equal(t, []string{"bug"}, *dynoIssues[4].Labels)

	assert.Equal(t, "DYNO Fuzz: InvalidDynamicObjectChecker at Endpoint /endpoint", *dynoIssues[5].Title)
	assert.Equal(t, "\nDetails: Detects 500 errors or unexpected success status codes when invalid dynamic objects are sent in requests.\n", *dynoIssues[5].Details)
	assert.Equal(t, "\nVisualizer: [DYNO](the web url)\n", *dynoIssues[5].Visualizer)
	assert.Equal(t, []string{"bug"}, *dynoIssues[5].Labels)

	assert.Equal(t, "DYNO Fuzz: PayloadBodyChecker at Endpoint /endpoint", *dynoIssues[6].Title)
	assert.Equal(t, "\nDetails: Detects 500 errors when fuzzing the JSON bodies of requests.\n", *dynoIssues[6].Details)
	assert.Equal(t, "\nVisualizer: [DYNO](the web url)\n", *dynoIssues[6].Visualizer)
	assert.Equal(t, []string{"bug"}, *dynoIssues[6].Labels)
}