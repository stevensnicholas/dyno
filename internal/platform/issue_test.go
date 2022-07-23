package platform_test

import (
	"golambda/internal/parse"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateIssueInvalid(t *testing.T) {
	fuzzError := "InternalServerErrors"
	newIssueRequest := parse.FuzzBugCheck(fuzzError, "body", "/endpoint", nil, nil, nil)
	assert.Equal(t, "DYNO Fuzz: InternalServerErrors at Endpoint /endpoint", *newIssueRequest.Title)
	assert.Equal(t, "body", *newIssueRequest.Body)
	assert.Nil(t, newIssueRequest.Assignee)
	assert.Nil(t, newIssueRequest.State)
	assert.Nil(t, newIssueRequest.Milestone)
}

func TestCreateIssueValid(t *testing.T) {
	body := "body"
	endpoint := "/endpoint"
	assignee := "fishua"
	state := "state"
	milestone := 1

	fuzzErrorList := [7]string{
		"InternalServerErrors",
		"UseAfterFreeChecker",
		"NameSpaceRuleChecker",
		"ResourceHierarchyChecker",
		"LeakageRuleChecker",
		"InvalidDynamicObjectChecker",
		"PayloadBodyChecker",
	}

	newIssueRequest := parse.FuzzBugCheck(fuzzErrorList[0], body, endpoint, &assignee, &state, &milestone)
	assert.Equal(t, "DYNO Fuzz: InternalServerErrors at Endpoint /endpoint", *newIssueRequest.Title)
	assert.Equal(t, "body", *newIssueRequest.Body)
	assert.Equal(t, "fishua", *newIssueRequest.Assignee)
	assert.Equal(t, "state", *newIssueRequest.State)
	assert.Equal(t, 1, *newIssueRequest.Milestone)

	newIssueRequest = parse.FuzzBugCheck(fuzzErrorList[1], body, endpoint, &assignee, &state, &milestone)
	assert.Equal(t, "DYNO Fuzz: UseAfterFreeChecker at Endpoint /endpoint", *newIssueRequest.Title)
	assert.Equal(t, "body", *newIssueRequest.Body)
	assert.Equal(t, "fishua", *newIssueRequest.Assignee)
	assert.Equal(t, "state", *newIssueRequest.State)
	assert.Equal(t, 1, *newIssueRequest.Milestone)

	newIssueRequest = parse.FuzzBugCheck(fuzzErrorList[2], body, endpoint, &assignee, &state, &milestone)
	assert.Equal(t, "DYNO Fuzz: NameSpaceRuleChecker at Endpoint /endpoint", *newIssueRequest.Title)
	assert.Equal(t, "body", *newIssueRequest.Body)
	assert.Equal(t, "fishua", *newIssueRequest.Assignee)
	assert.Equal(t, "state", *newIssueRequest.State)
	assert.Equal(t, 1, *newIssueRequest.Milestone)

	newIssueRequest = parse.FuzzBugCheck(fuzzErrorList[3], body, endpoint, &assignee, &state, &milestone)
	assert.Equal(t, "DYNO Fuzz: ResourceHierarchyChecker at Endpoint /endpoint", *newIssueRequest.Title)
	assert.Equal(t, "body", *newIssueRequest.Body)
	assert.Equal(t, "fishua", *newIssueRequest.Assignee)
	assert.Equal(t, "state", *newIssueRequest.State)
	assert.Equal(t, 1, *newIssueRequest.Milestone)

	newIssueRequest = parse.FuzzBugCheck(fuzzErrorList[4], body, endpoint, &assignee, &state, &milestone)
	assert.Equal(t, "DYNO Fuzz: LeakageRuleChecker at Endpoint /endpoint", *newIssueRequest.Title)
	assert.Equal(t, "body", *newIssueRequest.Body)
	assert.Equal(t, "fishua", *newIssueRequest.Assignee)
	assert.Equal(t, "state", *newIssueRequest.State)
	assert.Equal(t, 1, *newIssueRequest.Milestone)

	newIssueRequest = parse.FuzzBugCheck(fuzzErrorList[5], body, endpoint, &assignee, &state, &milestone)
	assert.Equal(t, "DYNO Fuzz: InvalidDynamicObjectChecker at Endpoint /endpoint", *newIssueRequest.Title)
	assert.Equal(t, "body", *newIssueRequest.Body)
	assert.Equal(t, "fishua", *newIssueRequest.Assignee)
	assert.Equal(t, "state", *newIssueRequest.State)
	assert.Equal(t, 1, *newIssueRequest.Milestone)

	newIssueRequest = parse.FuzzBugCheck(fuzzErrorList[6], body, endpoint, &assignee, &state, &milestone)
	assert.Equal(t, "DYNO Fuzz: PayloadBodyChecker at Endpoint /endpoint", *newIssueRequest.Title)
	assert.Equal(t, "body", *newIssueRequest.Body)
	assert.Equal(t, "fishua", *newIssueRequest.Assignee)
	assert.Equal(t, "state", *newIssueRequest.State)
	assert.Equal(t, 1, *newIssueRequest.Milestone)
}