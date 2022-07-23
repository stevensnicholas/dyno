package platform

import (
	"golambda/internal/result"
	"context"
	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

// Creates an Authenticated Client for communitcation with github
// Inputs: token is the user token
// Returns the Client
func CreateClient(ctx context.Context, token *string) *github.Client {
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

func FormatFuzzBody(dynoResult *result.DynoResult) *string {
	body := "# " + *dynoResult.Title + "\n" + *dynoResult.Details 
	if *dynoResult.MethodInformation.ContentType != ""{
		body = body + "\n" + *dynoResult.Method + "\n" + "\n" + "- " + *dynoResult.MethodInformation.AcceptedResponse + "\n" + "- " + *dynoResult.MethodInformation.Host + "\n" + "- " + *dynoResult.MethodInformation.ContentType
	} else {
		body = body + "\n" + *dynoResult.Method + "\n" + "\n" + "- " + *dynoResult.MethodInformation.AcceptedResponse + "\n" + "- " + *dynoResult.MethodInformation.Host
	}
	if *dynoResult.MethodInformation.Request != "" {
		body = body + "\n" + "- " + *dynoResult.MethodInformation.Request 
	}
	body = body + "\n" + "\n" + *dynoResult.TimeDelay + "\n" + *dynoResult.AsyncTime + "\n" + "\n" + *dynoResult.PreviousResponse
	body = body + "\n"
	return &body
}
