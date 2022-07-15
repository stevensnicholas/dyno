package github_test

import (
	"testing"
	"github.com/joho/godotenv"
	"os"
)
func getEnv() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	return os.Getenv("GITHUB_API")
}

func TestValidClient(t *testing.T) {
	
}

func TestInvalidClient() {

}

func TestValidGetRepo() {

}

func TestInvalidGetRepo() {

}

func TestValidCreateIssueRequest() {

}

func TestInValidCreateIssueRequest() {
	
}