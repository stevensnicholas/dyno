package main

import (
	"context"
	"dyno/internal/logger"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

// Testing SQS message is sent to Lambda with the correct message contents
func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	logger.Warn("Function invoked!")
	if len(sqsEvent.Records) == 0 {
		logger.Warn("No SQS events")
	}
	for _, message := range sqsEvent.Records {
		logger.Warnf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)
	}
	return nil
}
