package main

import (
	"golambda/internal/platform"
	"golambda/internal/parse"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// repoName := "Demo-Server"
	// owner := "thefishua"
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	token := os.Getenv("GITHUB_API")
	println(token)
	file := "internal/tests/bug_buckets/bug_buckets.txt"
	dynoResults := parse.ParseRestlerFuzzResults(file)
	platform.CreateIssues(dynoResults)
	println("done")
}
