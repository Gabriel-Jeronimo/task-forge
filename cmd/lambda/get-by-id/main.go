package main

import (
	"encoding/json"
	"time"
	"a/cmd/types"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
