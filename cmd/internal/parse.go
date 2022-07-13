package main

import (
	"strings"
	"fmt"
	"github.com/google/go-github/v45/github"
	"bufio"
	"os"
)

// TODO Change the process of reading file according to SQS and S3 Buckets
const bugFile = 6

func ParseFuzz(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	// Categorizing Errors
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
				errorCheck(fuzzError[0], body, endpoint, nil, nil, nil)
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
			
			method := strings.Trim(requestSplit[0], "\\r")
			accept := strings.Trim(requestSplit[1], "\\r")
			host := strings.Trim(requestSplit[2], "\\r")
			contentType := strings.Trim(requestSplit[3], "\\r")
			request := ""
			for i := 4; i < len(requestSplit); i++ {
				request = request + requestSplit[i]
			}
			request = strings.Trim(request, "\\r")
			scanner.Scan()
			timeDelay := scanner.Text()
			scanner.Scan()
			asyncTime := scanner.Text()
			scanner.Scan()
			previousResponseText := scanner.Text()
			previousResponseSplit := strings.Split(previousResponseText, "\\n")
			prevresp := ""
			for i := 5; i < len(previousResponseSplit); i++ {
				prevresp = prevresp + previousResponseSplit[i]
			}
			prevresp = strings.Trim(prevresp, "\\r")
			previousResponse := strings.Trim(previousResponseSplit[0], "\\r") + " request:" + prevresp
			if contentType != ""{
				body = body + method + accept + host + contentType
			} else {
				body = body + method + accept + host
			}
			if request != "" {
				body = body + request 
			}
			body = body + timeDelay + asyncTime + previousResponse
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return body, endpoint
}

func errorCheck(s string, body string, endpoint string, assignee *string, state *string, milestone *int) {
	if s == "InternalServerErrors" {
		InternalServerErrors(body, endpoint, assignee, state, milestone)
	}

	if s == "UseAfterFreeChecker" {
		UseAfterFreeChecker(body, endpoint, assignee, state, milestone)
	}

	if s == "NameSpaceRuleChecker" {
		NameSpaceRuleChecker(body, endpoint, assignee, state, milestone)
	}

	if s == "ResourceHierarchyChecker" {
		ResourceHierarchyChecker(body, endpoint, assignee, state, milestone)
	}

	if s == "LeakageRuleChecker" {
		LeakageRuleChecker(body, endpoint, assignee, state, milestone)
	}

	if s == "InvalidDynamicObjectChecker" {
		InvalidDynamicObjectChecker(body, endpoint, assignee, state, milestone)
	}

	if s == "PayloadBodyChecker" {
		PayloadBodyChecker(body, endpoint, assignee, state, milestone)
	}
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