package main

import (
	"context"
	util "github.com/Fallenstedt/lex/packages/util"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type LexConversationEvent struct {
	Text string `json:"text"`
	SessionId string `json:"sessionId"`
}
 
type LexConversationAnswer struct {
	Message string `json:"message:"`
	SessionId string `json:"sessionId"`
}

func HandleLambdaEvent(event LexConversationEvent) (LexConversationAnswer, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	session := util.NewLex(ctx)
	answer := <- session.RecognizeText(ctx, &event.Text, &event.SessionId)

	if answer.Err != nil {
		log.Fatalf("Failed to recognize text: %v ",answer.Err)
	}

	return LexConversationAnswer{
		Message: *answer.Output.Messages[0].Content,
		SessionId: *answer.Output.SessionId,
	}, nil
}

 
func main() {
        lambda.Start(HandleLambdaEvent)
}