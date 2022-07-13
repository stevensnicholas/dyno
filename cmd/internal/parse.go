package main

import (
	"fmt"
	"github.com/google/go-github/v45/github"
	"bufio"
	"os"
)

func ParseFuzz(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	// Categorizing Errors
	for scanner.Scan() {
		println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	} 
	
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