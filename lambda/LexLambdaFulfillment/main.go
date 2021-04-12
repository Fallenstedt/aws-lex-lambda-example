package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleLambdaEvent(event interface{}) (interface{}, error) {
	fmt.Println("Got Event")
	fmt.Println(event)
	return event, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}