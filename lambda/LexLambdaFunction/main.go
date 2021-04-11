package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimev2"
	"github.com/google/uuid"
	"log"
)

type LexConversationEvent struct {
	Text string `json:"text"`
}
 
type MyResponse struct {
        Message string `json:"Answer:"`
}
 
func HandleLambdaEvent(event LexConversationEvent) (MyResponse, error) {
    fmt.Println("creating context")
	ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Failed to load default configuration: %v", err.Error())
	}

	lex := lexruntimev2.NewFromConfig(cfg)

	botAliasid := ""
	botId := ""
	localeId := "en_US"
	sessionId := uuid.NewString()

	resp, err := lex.RecognizeText(ctx, &lexruntimev2.RecognizeTextInput{
		BotAliasId:        &botAliasid,
		BotId:             &botId,
		LocaleId:          &localeId,
		SessionId:         &sessionId,
		Text:              &event.Text,
	})
	if err != nil {
		log.Fatalf("Failed to recognize text: %v ",err)
	}

	return MyResponse{
		Message: *resp.Messages[0].Content,
	}, nil


}
 
func main() {
        lambda.Start(HandleLambdaEvent)
}