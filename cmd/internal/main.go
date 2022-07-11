package main

import (
	// "context"
)

func main() {
	// repoName := "Demo-Server"
	// owner := "thefishua"
	// token := "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" // TODO Remove this before push to branch
	// ctx := context.Background()
	// client := CreateClient(ctx, token)
	// title := "First issue"
	// body := "This is a test"
	// client.Issues.Create(ctx, owner, repoName, CreateIssueRequest(&title, &body, nil, nil, nil, nil))	
	ParseFuzz("cmd/internal/tests/bug_buckets.txt")
}
