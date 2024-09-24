package openai

import (
	"context"
	"os"
	"path"

	"github.com/evadcmd/bot/internal/util"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

var client *openai.Client

func ChatCompletion(ctx context.Context, model LLMModel, userPrompt string, stop []string) (string, error) {
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: string(model),
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userPrompt,
				},
			},
			Temperature: 0.2,
			Stop:        stop,
		})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

func init() {
	godotenv.Load(path.Join(util.RootPath, ".env"))
	client = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
}
