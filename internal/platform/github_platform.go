package platform

import (
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
func GetRepo(ctx context.Context, client *github.Client, owner string, repoName string) *github.Repository {
	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		panic(err)
	}
	return repo
}
