package main

func main() {
	repoName := "Demo-Server"
	owner := "thefishua"
	token := "token"
	file := "cmd/internal/tests/bug_buckets/bug_buckets.txt"
	ParseFuzz(token, repoName, owner, file)
	println("done")
}
