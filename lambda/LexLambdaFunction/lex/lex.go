package lex

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimev2"
	"github.com/google/uuid"
	"log"
)

type (
	ILex interface {
		RecognizeText(ctx context.Context, text *string, sessionId *string)  <-chan RecognizeTextResult
	}

	Lex struct {
		session *lexruntimev2.Client
	}

	RecognizeTextResult struct {
		Output *lexruntimev2.RecognizeTextOutput
		Err error
	}
)

func NewLex(parentCtx context.Context) ILex {
	ctx, cancel := context.WithCancel(parentCtx)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Printf("Failed to load default configuration: %v", err.Error())
		cancel()
	}

	return &Lex{session: lexruntimev2.NewFromConfig(cfg)}
}

func (l *Lex) RecognizeText(ctx context.Context, text *string, sessionId *string) <-chan RecognizeTextResult {
	result := make(chan RecognizeTextResult)

	go func() {
		defer close(result)

		findResultLoop:
			for {
				select {
					case <- ctx.Done():
						result <- RecognizeTextResult{
							Output: nil,
							Err: ctx.Err(),
						}
						break findResultLoop
					default:
						textInput := l.buildRecognizeTextInput(text)(sessionId)
						resp, err := l.session.RecognizeText(ctx, textInput)
						result <- RecognizeTextResult{
							Output: resp,
							Err: err,
						}
						break findResultLoop
				}
			}
	}()

	return result
}


func (l *Lex) buildRecognizeTextInput(text *string) func(sessionId *string) *lexruntimev2.RecognizeTextInput {
	botAliasid := ""
	botId := ""
	localeId := "en_US"

	return func(sessionId *string) *lexruntimev2.RecognizeTextInput {
		var id *string
		if *sessionId == "" {
			newId := uuid.NewString()
			id = &newId
		} else {
			id = sessionId
		}

		return &lexruntimev2.RecognizeTextInput{
			BotAliasId:        &botAliasid,
			BotId:             &botId,
			LocaleId:          &localeId,
			SessionId:         id,
			Text:              text,
		}
	}
}