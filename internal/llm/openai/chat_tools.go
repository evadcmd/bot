package openai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

// Type aliases over github.com/sashabaranov/go-openai's tool-calling types.
// Kept in this package for the same reason LLMModel wraps model name
// constants: it keeps the SDK dependency contained to this package instead
// of leaking into every caller.
type (
	Message            = openai.ChatCompletionMessage
	ToolCall           = openai.ToolCall
	ToolDefinition     = openai.Tool
	FunctionDefinition = openai.FunctionDefinition
	FinishReason       = openai.FinishReason
)

const (
	RoleSystem    = openai.ChatMessageRoleSystem
	RoleUser      = openai.ChatMessageRoleUser
	RoleAssistant = openai.ChatMessageRoleAssistant
	RoleTool      = openai.ChatMessageRoleTool
)

const (
	FinishReasonStop      = openai.FinishReasonStop
	FinishReasonToolCalls = openai.FinishReasonToolCalls
	FinishReasonLength    = openai.FinishReasonLength
)

const ToolTypeFunction = openai.ToolTypeFunction

// ChatCompletionWithTools sends a multi-turn conversation together with a
// set of callable tool definitions and returns the assistant's reply
// message along with the reason the model stopped generating. A
// FinishReason of FinishReasonToolCalls means the returned message's
// ToolCalls should be executed and fed back as tool-role messages in the
// next call.
func ChatCompletionWithTools(ctx context.Context, model LLMModel, messages []Message, tools []ToolDefinition) (Message, FinishReason, error) {
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:       string(model),
		Messages:    messages,
		Tools:       tools,
		Temperature: 0.2,
	})
	if err != nil {
		return Message{}, "", fmt.Errorf("failed to send a tool-calling request to OpenAI API server: %w", err)
	}
	choice := resp.Choices[0]
	return choice.Message, choice.FinishReason, nil
}