package main

import (
	"github.com/google/go-github/v45/github"
)

// TODO Change the process of reading file according to SQS and S3 Buckets
// const bugFile = 6

// Parses the fuzzing files from the bug_buckets folder and creates github issues 
// Inputs: 
//				token is the user token 
//				repoName is the user's repo
//				owner is the owner of the repo
//				file is filepath to the bug_buckets.txt file that stores all the bugs that has occured
func ParseFuzz(token string, repoName string, owner string, file string) {
	
}
// Reads the a bug_bucket file that is specified by the category of the bug found by restler
// Creates the body of the issue in regards to the bug found by the fuzzer with details on the bug and how to fix itInternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug 
// 
// Inputs: 
//				bugFileName is the name of the file that has the logs of the bug 
//        body is the start of the body for the github issue
// Returns: 
// 				body is the body of the issue 
// 				endpoint is the endpoint that has the bug 
func ReadBugFile(location string, bugFileName string, body string) (string, string){
	return "", ""
}

// FuzzBugCheck sorts the bugs found by the fuzzer by there categories and creates a new github issueRequest
// Inputs: 
//				fuzzError is the type of bug that has been found by the fuzzer
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug 
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue 
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo  
// Returns: 
// 				*github.IssueRequest with all the relevant information regarding the certain bug
func FuzzBugCheck(fuzzError string, body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest{
	newIssueRequest := &github.IssueRequest{}
	return newIssueRequest
}
// AddDYNODetails adds the details and visualizer url to the body of the issue request that 
// corresponds to the FuzzError that is received
// Inputs: 
//				fuzzError is the type of bug that has been found by the fuzzer
// Returns: 
// 				details a string with the specified details 
// 				If not details are created leaves an empty string
func AddDYNODetails(fuzzError string) string{
	return ""
}
// InternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug 
// Inputs:
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug 
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue 
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo  
// Returns: 
// 				*github.IssueRequest with all the relevant information regarding the certain bug

func InternalServerErrors(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest{
	newIssueRequest := &github.IssueRequest{}
	return newIssueRequest
}
// InternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug 
// Inputs:
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug 
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue 
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo  
// Returns: 
// 				*github.IssueRequest with all the relevant information regarding the certain bug

func ResourceHierarchyChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest{
	newIssueRequest := &github.IssueRequest{}
	return newIssueRequest
}
// InternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug 
// Inputs:
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug 
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue 
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo  
// Returns: 
// 				*github.IssueRequest with all the relevant information regarding the certain bug

func NameSpaceRuleChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest{
	newIssueRequest := &github.IssueRequest{}
	return newIssueRequest
}
// InternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug 
// Inputs:
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug 
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue 
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo  
// Returns: 
// 				*github.IssueRequest with all the relevant information regarding the certain bug

func UseAfterFreeChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest{
	newIssueRequest := &github.IssueRequest{}
	return newIssueRequest
}
// InternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug 
// Inputs:
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug 
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue 
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo  
// Returns: 
// 				*github.IssueRequest with all the relevant information regarding the certain bug

func LeakageRuleChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest{
	newIssueRequest := &github.IssueRequest{}
	return newIssueRequest
}
// InternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug 
// Inputs:
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug 
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue 
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo  
// Returns: 
// 				*github.IssueRequest with all the relevant information regarding the certain bug

func InvalidDynamicObjectChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest{
	newIssueRequest := &github.IssueRequest{}
	return newIssueRequest
}
// InternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug 
// Inputs:
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug 
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue 
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo  
// Returns: 
// 				*github.IssueRequest with all the relevant information regarding the certain bug

func PayloadBodyChecker(body string, endpoint string, assignee *string, state *string, milestone *int)  *github.IssueRequest{
	newIssueRequest := &github.IssueRequest{}
	return newIssueRequest
}
