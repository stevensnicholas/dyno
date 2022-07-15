package github_test

import (
	"context"
	"testing"
	"github.com/joho/godotenv"
	"os"
	"golambda/internal/github"
	"github.com/stretchr/testify/assert"
)
func getEnv() string {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	return os.Getenv("GITHUB_API")
}

func TestValidClient(t *testing.T) {
	token := getEnv()
	ctx := context.Background()
	actualClient := github.CreateClient(ctx, &token)
	assert.Equal(t, "api.github.com", actualClient.BaseURL.Host)
}

func TestInvalidClient(t *testing.T) {
	var token *string = nil
	ctx := context.Background()
	assert.Panics(t, func() {github.CreateClient(ctx, token)}, "Nil passed as token")
}


func TestValidGetRepo(t *testing.T) {
	token := getEnv()
	ctx := context.Background()
	client := github.CreateClient(ctx, &token)
	actualRepo := github.GetRepo(ctx, client, "thefishua", "Demo-Server")
	assert.Equal(t, "thefishua/Demo-Server", *actualRepo.FullName)
}

func TestInvalidGetRepo(t *testing.T) {
	token := getEnv()
	ctx := context.Background()
	client := github.CreateClient(ctx, &token)
	actualRepo := github.GetRepo(ctx, client, "thefishua", "Demo-Server")
	assert.Equal(t, "thefishua/Demo-Server", *actualRepo.FullName)
}