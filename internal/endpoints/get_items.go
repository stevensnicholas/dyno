package endpoints

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	clientID   string
	Title      string
	Details    string
	Visualizer string
	Body       string
	Assignee   string
	Labels     string
	State      string
	Milestone  string
}

func GetDBItem(svc *dynamodb.DynamoDB) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("dyno_table"),
		Key: map[string]*dynamodb.AttributeValue{
			"clientID": {
				S: aws.String("1"),
			},
			"Title": {
				S: aws.String("Fuzzing Error 1"),
			},
			"Details": {
				S: aws.String("Weak Sql statement"),
			},
			"Visualizer": {
				S: aws.String("Visualizer Test1"),
			},
			"Body": {
				S: aws.String("Visualizer Test1"),
			},
			"Assignee": {
				S: aws.String("Joe"),
			},
			"Labels": {
				S: aws.String("1"),
			},
			"State": {
				S: aws.String("Test State"),
			},
			"Milestone": {
				S: aws.String("Test Milestone"),
			},
		},
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}
	item := Item{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	fmt.Println("Found item:")
	fmt.Println("Year:  ", item.clientID)
}
