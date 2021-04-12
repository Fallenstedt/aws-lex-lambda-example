package main

import (
	//"context"
	"encoding/json"
	"fmt"

	//util "github.com/Fallenstedt/lex/packages/util"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimev2"
)

func HandleLambdaEvent(event lexruntimev2.PutSessionInput) (lexruntimev2.PutSessionInput, error) {
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	fmt.Println(event.SessionId)
	d, _ := json.Marshal(event)
	fmt.Println(string(d))

	return event, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}