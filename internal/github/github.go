package github

import (
	"context"
	"golang.org/x/oauth2"
	"github.com/google/go-github/v45/github"
)

// Creates an Authenticated Client 
// Inputs: token is the user token 
// Returns the Client 
func CreateClient(ctx context.Context, token string) *github.Client {
	authToken := token
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: authToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	
	return client
}
// Gets the specific repo provided by the owner and given repo name.
// Inputs: owner is the owner of the repo, repoName is the name of the repo.
// Returns: the specified repo as a *github.Repository 
func GetRepo(ctx context.Context, client *github.Client, owner string, repoName string) *github.Repository {
	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		panic(err)
	}
	return repo
}
// Creates an IssueRequest that provides the body of the vulnerabilities within the software.
// Parses the files from fuzzing into an IssueRequest.
// Returns an IssueRequest.
// I think this function is redundant and actually dont need it
func CreateIssueRequest(title *string, body *string, labels *[]string, assignee *string, state *string, milestone *int) *github.IssueRequest {
	newIssueRequest := &github.IssueRequest{}
	
	if title != nil {
		newIssueRequest.Title = title
	}
	if body != nil {
		newIssueRequest.Body = body
	}
	if labels != nil {
		newIssueRequest.Labels = labels
	}
	if assignee != nil {
		newIssueRequest.Assignee = assignee
	}
	if state != nil {
		newIssueRequest.State = state
	}
	if milestone != nil {
		newIssueRequest.Milestone = milestone
	}
	return newIssueRequest
}