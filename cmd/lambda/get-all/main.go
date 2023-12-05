package main

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Task struct {
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	Done      bool   `json:"done"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	task := Task{
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
