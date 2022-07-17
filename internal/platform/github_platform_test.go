package platform_test


import (
	"context"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"golambda/internal/platform"
	"os"
	"testing"
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
	actualClient := platform.CreateClient(ctx, &token)
	assert.Equal(t, "api.github.com", actualClient.BaseURL.Host)
}

func TestInvalidClient(t *testing.T) {
	var token *string = nil
	ctx := context.Background()
	assert.Panics(t, func() { platform.CreateClient(ctx, token) }, "Nil passed as token")
}

func TestValidGetRepo(t *testing.T) {
	token := getEnv()
	ctx := context.Background()
	client := platform.CreateClient(ctx, &token)
	actualRepo := platform.GetRepo(ctx, client, "thefishua", "Demo-Server")
	assert.Equal(t, "thefishua/Demo-Server", *actualRepo.FullName)
}

func TestInvalidGetRepo(t *testing.T) {
	token := getEnv()
	ctx := context.Background()
	client := platform.CreateClient(ctx, &token)
	actualRepo := platform.GetRepo(ctx, client, "thefishua", "Demo-Server")
	assert.Equal(t, "thefishua/Demo-Server", *actualRepo.FullName)
}
