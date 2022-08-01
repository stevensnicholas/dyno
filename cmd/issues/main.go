package main

import (
	"github.com/aws/aws-lambda-go/events"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

// Testing SQS message is sent to Lambda with the correct message contents 
func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	fmt.Println("Function invoked!")
	if len(sqsEvent.Records) == 0 {
		fmt.Println("No SQS events")
	}
	for _, message := range sqsEvent.Records {
		fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)
	}
	return nil
}
