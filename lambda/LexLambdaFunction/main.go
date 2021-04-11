package main

import (
	"context"
	"encoding/json"
	"github.com/Fallenstedt/lex/lambda/LexLambdaFunction/lex"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type LexConversationEvent struct {
	Text string `json:"text"`
	SessionId string `json:"sessionId"`
}
 
type MyResponse struct {
        Message string `json:"Answer:"`
}

func HandleLambdaEvent(event LexConversationEvent) (MyResponse, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	session := lex.NewLex(ctx)
	answer := <- session.RecognizeText(ctx, &event.Text, &event.SessionId)

	if answer.Err != nil {
		log.Fatalf("Failed to recognize text: %v ",answer.Err)
	}

	dump, _ := json.Marshal(answer.Output)
	return MyResponse{
		Message: string(dump),
	}, nil

}

 
func main() {
        lambda.Start(HandleLambdaEvent)
}