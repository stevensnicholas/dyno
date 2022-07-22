package main

import (
	"golambda/internal/parse"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	repoName := "Demo-Server"
	owner := "thefishua"
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	token := os.Getenv("GITHUB_API")
	println(token)
	file := "internal/tests/bug_buckets/bug_buckets.txt"
	dynoIssues := parse.ParseFuzz(token, repoName, owner, file)
	for i,issue := range dynoIssues {
		println(i, *issue.Title)
	}
	// slice := []int{}
	// for i := 0; i < 5; i++ {
	// 	slice = append(slice, i)
	// }

	// for _, e := range slice {
	// 	println(e)
	// }
}
