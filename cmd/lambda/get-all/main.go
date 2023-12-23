package main

import (
	"a/cmd/types"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	tableName := "TasksTable"
	userId := "a2d8dfaa-80ee-4ec9-b79c-d0b348b7d6f8"
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := dynamodb.New(sess)

	resp, err := client.Query(&dynamodb.QueryInput{
		TableName:              &tableName,
		KeyConditionExpression: aws.String("userId = :val"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":val": {
				S: aws.String(userId),
			},
		},
	})

	if err != nil {
		log.Fatalf("unable to query", err)
	}

	fmt.Printf("%v", resp)

	task := types.Task{
		Name:      "Buy bread",
		CreatedAt: time.Now().String(),
		Done:      false,
	}

	response, _ := json.Marshal(task)

	return events.APIGatewayProxyResponse{
		Body:       string(response),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
