package main

import (
	"LayeredStructured/infrastructure"
	"LayeredStructured/initializers"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest() (events.APIGatewayProxyResponse, error) {

	initializers.LoadEnvVariables()
	cfg := infrastructure.NewAWSConfig()
	out := infrastructure.NewDynamoDBClient(cfg)
	fmt.Println(out)

	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string("successfully"),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil

}

func main() {
	lambda.Start(handleRequest)
}

func HelloWorld() {
	fmt.Println("Say Hello")
}
