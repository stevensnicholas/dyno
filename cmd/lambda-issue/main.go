package main

import (
	"github.com/google/go-github/v45/github"
	"context"
	"golambda/internal/platform"
	"golambda/internal/parse"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	repoName := ""
	owner := ""
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	token := os.Getenv("GITHUB_API")
	ctx := context.Background()
	client := platform.CreateClient(ctx, &token)
	file := "internal/tests/bug_buckets/bug_buckets.txt"
	dynoResults := parse.ParseRestlerFuzzResults(file)
	dynoIssues := platform.CreateIssues(dynoResults)
	for _, issue := range dynoIssues {
		client.Issues.Create(ctx, owner, repoName, &github.IssueRequest{
			Title:     issue.Title,
			Body:      platform.FormatFuzzBody(issue.Body),
			Labels:    issue.Labels,
			Assignee:  issue.Assignee,
			State:     issue.State,
			Milestone: issue.Milestone,
		})
	}
	println("done")
}
