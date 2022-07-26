package platform

import (
	"context"
	"dyno/internal/issue"
	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

// Creates an Authenticated Client for communitcation with github
// Inputs: token is the user token
// Returns the Client
func CreateGithubClient(ctx context.Context, token *string) *github.Client {
	authToken := token
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *authToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return client
}

// Gets the specific repo provided by the owner and given repo name.
// Inputs: owner is the owner of the repo, repoName is the name of the repo.
// Returns: the specified repo as a *github.Repository
func GetRepo(ctx context.Context, client *github.Client, owner *string, repoName *string) *github.Repository {
	repo, _, err := client.Repositories.Get(ctx, *owner, *repoName)
	if err != nil {
		panic(err)
	}
	return repo
}

// FormatFuzzBody formats the structs of DynoResults and creates a string to be used within
// a github issue and be presented to the user
// Inputted is a DynoResult and outputted is a string used for the github issue
func FormatFuzzBody(dynoIssue *issue.DynoIssue) *string {
	body := "# " + *dynoIssue.Body.Title + "\n\n" + *dynoIssue.Details + "\n\n" + *dynoIssue.Visualizer + "\n"
	if dynoIssue.Body.MethodInformation.ContentType != nil && *dynoIssue.Body.MethodInformation.ContentType != "" {
		body = body + "\n" + *dynoIssue.Body.Method + "\n" + "\n" + "- " + *dynoIssue.Body.MethodInformation.AcceptedResponse + "\n" + "- " + *dynoIssue.Body.MethodInformation.Host + "\n" + "- " + *dynoIssue.Body.MethodInformation.ContentType
	} else {
		body = body + "\n" + *dynoIssue.Body.Method + "\n" + "\n" + "- " + *dynoIssue.Body.MethodInformation.AcceptedResponse + "\n" + "- " + *dynoIssue.Body.MethodInformation.Host
	}
	if dynoIssue.Body.MethodInformation.Request != nil && *dynoIssue.Body.MethodInformation.Request != "" {
		body = body + "\n" + "- " + *dynoIssue.Body.MethodInformation.Request
	}
	body = body + "\n" + "\n" + *dynoIssue.Body.TimeDelay + "\n" + *dynoIssue.Body.AsyncTime + "\n" + "\n" + *dynoIssue.Body.PreviousResponse
	body = body + "\n"
	return &body
}
