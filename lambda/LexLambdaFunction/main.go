package main

import (
	"context"
	"encoding/json"
	util "github.com/Fallenstedt/lex/packages/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type LexConversationEvent struct {
	Text string `json:"text"`
	SessionId string `json:"sessionId"`
}
 
type LexConversationAnswer struct {
	Message string `json:"message"`
	SessionId string `json:"sessionId"`
}

func HandleLambdaEvent(lambdaCtx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx, cancel := context.WithCancel(lambdaCtx)
	defer cancel()

	event := new(LexConversationEvent)
	err := json.Unmarshal([]byte(request.Body), event)
	if err != nil {
		log.Fatalf("Failed to unmarshal body")
	}

	session := util.NewLex(ctx)
	answer := <- session.RecognizeText(ctx, &event.Text, &event.SessionId)

	if answer.Err != nil {
		log.Fatalf("Failed to recognize text: %v ",answer.Err)
	}

	bodyAnswer, err := json.Marshal(LexConversationAnswer{
		Message: *answer.Output.Messages[0].Content,
		SessionId: *answer.Output.SessionId,
	})

	if err != nil {
		log.Fatalf("Failed to marshal answer: %v ", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "http://localhost:3000",
		},
		MultiValueHeaders: nil,
		Body:              string(bodyAnswer),
		IsBase64Encoded:   false,
	}, nil
}

 
func main() {
        lambda.Start(HandleLambdaEvent)
}