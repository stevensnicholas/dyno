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
				body, endpoint := ReadBugFile(bugFileName, s)
				client.Issues.Create(ctx, owner, repoName, errorCheck(fuzzError[0], body, endpoint, nil, nil, nil))
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	} 
	
}

func ReadBugFile(bugFileName string, body string) (string, string){
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
	return body, endpoint
}

func errorCheck(fuzzError string, body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest{
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

func InternalServerErrors(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest {
	title := fmt.Sprintf("DYNO Fuzz: InternalServerErrors at Endpoint %s",  endpoint)
	labels := []string{"bug"}
	return CreateIssueRequest(&title, &body, &labels, assignee, state, milestone);
}

func ResourceHierarchyChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest {
	title := fmt.Sprintf("DYNO Fuzz: ResourceHierarchyChecker at Endpoint %s",  endpoint)
	labels := []string{"bug"}
	return CreateIssueRequest(&title, &body, &labels, assignee, state, milestone);
}

func NameSpaceRuleChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest {
	title := fmt.Sprintf("DYNO Fuzz: NameSpaceRuleChecker at Endpoint %s",  endpoint)
	labels := []string{"bug"}
	return CreateIssueRequest(&title, &body, &labels, assignee, state, milestone);
}

func UseAfterFreeChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest {
	title := fmt.Sprintf("DYNO Fuzz: UseAfterFreeChecker at Endpoint %s",  endpoint)
	labels := []string{"bug"}
	return CreateIssueRequest(&title, &body, &labels, assignee, state, milestone);
}

func LeakageRuleChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest {
	title := fmt.Sprintf("DYNO Fuzz: LeakageRuleChecker at Endpoint %s",  endpoint)
	labels := []string{"bug"}
	return CreateIssueRequest(&title, &body, &labels, assignee, state, milestone);
}

func InvalidDynamicObjectChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest {
	title := fmt.Sprintf("DYNO Fuzz: InvalidDynamicObjectChecker at Endpoint %s",  endpoint)
	labels := []string{"bug"}
	return CreateIssueRequest(&title, &body, &labels, assignee, state, milestone);
}

func PayloadBodyChecker(body string, endpoint string, assignee *string, state *string, milestone *int) *github.IssueRequest {
	title := fmt.Sprintf("DYNO Fuzz: PayloadBodyChecker at Endpoint %s",  endpoint)
	labels := []string{"bug"}
	return CreateIssueRequest(&title, &body, &labels, assignee, state, milestone);
}