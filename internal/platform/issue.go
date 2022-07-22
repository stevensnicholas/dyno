package platform

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"golambda/internal/result"
)

type DynoIssue struct {
	Title *string `json:"title,omitempty"`
	Body *result.DynoResult `json:"body,omitempty"`
	Assignee *string `json:"assignee,omitempty"`
	Labels *[]string `json:"labels,omitempty"`
	State *string `json:"state,omitempty"`
	Milestone *int `json:"milestone,omitempty"`
}

// TODO Change the process of reading file according to SQS and S3 Buckets
const bugFile = 6

// Parses the fuzzing files from the bug_buckets folder and creates github issues
// Inputs:
//				token is the user token
//				repoName is the user's repo
//				owner is the owner of the repo
//				file is filepath to the bug_buckets.txt file that stores all the bugs that has occured
func ParseFuzzIssue(token string, repoName string, owner string, file string) []DynoIssue {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	location := "internal/tests/bug_buckets/"
	scanner := bufio.NewScanner(f)
	dynoIssueSlice := []DynoIssue{}
	dynoIssue := &DynoIssue{}
	dynoResult := &result.DynoResult{}
	// Sending an issue for each error found through fuzz
	for scanner.Scan() {
		line := scanner.Text()
		if line[:1] == "-" {
			scanner.Scan()
			line = scanner.Text()
			bugFileNames := strings.Fields(line)
			if len(bugFileNames) > 5 {
				bugFileName := bugFileNames[bugFile]
				fuzzError := strings.Split(bugFileNames[bugFile], "_")
				title := fmt.Sprintf("# %s Invalid %s Response\n", fuzzError[0], fuzzError[1])
				details := AddDYNODetails(fuzzError[0])
				dynoResult.Title = &title
				dynoResult.Details = &details
				dynoResult = CreatedynoResult(location, bugFileName, dynoResult)
				dynoIssue.Body = dynoResult
				dynoIssue = CreateIssue(fuzzError[0], dynoIssue)
				if dynoIssue != nil {
					dynoIssueSlice = append(dynoIssueSlice, *dynoIssue)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return dynoIssueSlice
}

// CreateIssue sorts the bugs found by the fuzzer by there categories and creates a new github issueRequest
// Inputs:
//				fuzzError is the type of bug that has been found by the fuzzer
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo
// Returns:
// 				*github.IssueRequest with all the relevant information regarding the certain bug
func CreateIssue(fuzzError string, dynoIssue *platform.DynoIssue) *platform.DynoIssue {
	switch fuzzError {
	case "InternalServerErrors":
		dynoIssue = InternalServerErrors(dynoIssue)
	case "UseAfterFreeChecker":
		dynoIssue = UseAfterFreeChecker(dynoIssue)
	case "NameSpaceRuleChecker":
		dynoIssue = NameSpaceRuleChecker(dynoIssue)
	case "ResourceHierarchyChecker":
		dynoIssue = ResourceHierarchyChecker(dynoIssue)
	case "LeakageRuleChecker":
		dynoIssue = LeakageRuleChecker(dynoIssue)
	case "InvalidDynamicObjectChecker":
		dynoIssue = InvalidDynamicObjectChecker(dynoIssue)
	default:
		dynoIssue = PayloadBodyChecker(dynoIssue)
	}
	return dynoIssue
}