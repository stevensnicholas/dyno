package main

import (
	// "context"
)

func main() {
	// repoName := "Demo-Server"
	// owner := "thefishua"
	// token := "token" // TODO Remove this before push to branch
	// ctx := context.Background()
	// client := CreateClient(ctx, token)
	// client.Issues.Create(ctx, owner, repoName, InternalServerError(nil, nil, nil))
	ParseFuzz("cmd/internal/tests/bug_buckets/bug_buckets.txt")
	println("done")
}
