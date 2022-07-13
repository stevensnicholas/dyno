package main

import (
	"context"
	"strings"
	"fmt"
	"github.com/google/go-github/v45/github"
	"bufio"
	"os"
)

// TODO Change the process of reading file according to SQS and S3 Buckets
const bugFile = 6

// Parses the fuzzing files from the bug_buckets folder and creates github issues 
// Inputs: 
//				token is the user token 
//				repoName is the user's repo
//				owner is the owner of the repo
//				file is filepath to the bug_buckets.txt file that stores all the bugs that has occured
func ParseFuzz(token string, repoName string, owner string, file string) {
	ctx := context.Background()
	client := CreateClient(ctx, token)
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	// Sending an issue for each error found through fuzz 
	for scanner.Scan() {
		line := scanner.Text()
		if line[:1]== "-" {
			scanner.Scan()
			line = scanner.Text()
			bugFileNames := strings.Fields(line)
			if len(bugFileNames) > 5 {
				bugFileName := bugFileNames[bugFile]
				fuzzError := strings.Split(bugFileNames[bugFile], "_")
				s := fmt.Sprintf("# %s Invalid %s Response\n", fuzzError[0], fuzzError[1])
				body, endpoint, err := ReadBugFile(bugFileName, s)
				if err != nil {
					panic(err)
				}
				client.Issues.Create(ctx, owner, repoName, FuzzBugCheck(fuzzError[0], body, endpoint, nil, nil, nil))
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	} 
	
}
// Reads the a bug_bucket file that is specified by the category of the bug found by restler
// Creates the body of the issue in regards to the bug found by the fuzzer with details on the bug and how to fix itInternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug 

// Inputs: 
//				bugFileName is the name of the file that has the logs of the bug 
//        body is the start of the body for the github issue
// Returns: 
// 				body is the body of the issue 
// 				endpoint is the endpoint that has the bug 
// 				error is if there is an error with the function 
func ReadBugFile(bugFileName string, body string) (string, string, error){
	endpoint := ""
	f, err := os.Open(fmt.Sprintf("cmd/internal/tests/bug_buckets/%s", bugFileName))
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[:1]== "-" {
			requestSplit := strings.Split(line, "\\n")
			endpoint = strings.Split(requestSplit[0], " ")[2]
			method := strings.Trim(requestSplit[0], "\\r")
			accept := strings.Trim(requestSplit[1], "\\r")
			host := strings.Trim(requestSplit[2], "\\r")
			contentType := strings.Trim(requestSplit[3], "\\r")
			request := ""
			for i := 4; i < len(requestSplit); i++ {
				request = request + requestSplit[i]
			}
			if request != "" {
				request = "Request: " + strings.Trim(request, "\\r")
			}
			scanner.Scan()
			timeDelay := scanner.Text()
			scanner.Scan()
			asyncTime := scanner.Text()
			scanner.Scan()
			previousResponseText := scanner.Text()
			previousResponseSplit := strings.Split(previousResponseText, "\\n")
			prevrequest := ""
			for i := 5; i < len(previousResponseSplit); i++ {
				prevrequest = prevrequest + previousResponseSplit[i]
			}
			if prevrequest != "" {
				prevrequest = " request:" + strings.Trim(prevrequest, "\\r")
			}
			previousResponse := strings.Trim(previousResponseSplit[0], "\\r") + prevrequest
			if contentType != ""{
				body = body + "\n" + method + "\n" + "\n" + "- " + accept + "\n" + "- " + host + "\n" + "- " + contentType
			} else {
				body = body + "\n" + method + "\n" + "\n" + "- " + accept + "\n" + "- " + host
			}
			if request != "" {
				body = body + "\n" + "- " + request 
			}
			body = body + "\n" + "\n" + timeDelay + "\n" + asyncTime + "\n" + "\n" + previousResponse
		}
		body = body + "\n"
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	
	return body, endpoint, err
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
	if fuzzError == "InternalServerErrors" {
		newIssueRequest = InternalServerErrors(body, endpoint, assignee, state, milestone)
	}

	if fuzzError == "UseAfterFreeChecker" {
		newIssueRequest = UseAfterFreeChecker(body, endpoint, assignee, state, milestone)
	}

	if fuzzError == "NameSpaceRuleChecker" {
		newIssueRequest = NameSpaceRuleChecker(body, endpoint, assignee, state, milestone)
	}

	if fuzzError == "ResourceHierarchyChecker" {
		newIssueRequest = ResourceHierarchyChecker(body, endpoint, assignee, state, milestone)
	}

	if fuzzError == "LeakageRuleChecker" {
		newIssueRequest = LeakageRuleChecker(body, endpoint, assignee, state, milestone)
	}

	if fuzzError == "InvalidDynamicObjectChecker" {
		newIssueRequest = InvalidDynamicObjectChecker(body, endpoint, assignee, state, milestone)
	}

	if fuzzError == "PayloadBodyChecker" {
		newIssueRequest = PayloadBodyChecker(body, endpoint, assignee, state, milestone)
	}

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

func InternalServerErrors(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest {
	title := fmt.Sprintf("DYNO Fuzz: InternalServerErrors at Endpoint %s",  endpoint)
	labels := []string{"bug"}
	return CreateIssueRequest(&title, &body, &labels, assignee, state, milestone);
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

func ResourceHierarchyChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest {
	title := fmt.Sprintf("DYNO Fuzz: ResourceHierarchyChecker at Endpoint %s",  endpoint)
	labels := []string{"bug"}
	return CreateIssueRequest(&title, &body, &labels, assignee, state, milestone);
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

func NameSpaceRuleChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest {
	title := fmt.Sprintf("DYNO Fuzz: NameSpaceRuleChecker at Endpoint %s",  endpoint)
	labels := []string{"bug"}
	return CreateIssueRequest(&title, &body, &labels, assignee, state, milestone);
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

func UseAfterFreeChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest {
	title := fmt.Sprintf("DYNO Fuzz: UseAfterFreeChecker at Endpoint %s",  endpoint)
	labels := []string{"bug"}
	return CreateIssueRequest(&title, &body, &labels, assignee, state, milestone);
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

func LeakageRuleChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest {
	title := fmt.Sprintf("DYNO Fuzz: LeakageRuleChecker at Endpoint %s",  endpoint)
	labels := []string{"bug"}
	return CreateIssueRequest(&title, &body, &labels, assignee, state, milestone);
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

func InvalidDynamicObjectChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest {
	title := fmt.Sprintf("DYNO Fuzz: InvalidDynamicObjectChecker at Endpoint %s",  endpoint)
	labels := []string{"bug"}
	return CreateIssueRequest(&title, &body, &labels, assignee, state, milestone);
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

func PayloadBodyChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest {
	title := fmt.Sprintf("DYNO Fuzz: PayloadBodyChecker at Endpoint %s",  endpoint)
	labels := []string{"bug"}
	return CreateIssueRequest(&title, &body, &labels, assignee, state, milestone);
}