package openai

import (
	"context"
	"encoding/json"

	"github.com/sashabaranov/go-openai"
)

// Type aliases over go-openai's tool-calling types, exported here so the
// rest of the codebase doesn't need to import github.com/sashabaranov/go-openai
// directly (same "regulate dependencies" motivation as LLMModel in llm_model.go).
type (
	Message      = openai.ChatCompletionMessage
	ToolCall     = openai.ToolCall
	FunctionCall = openai.FunctionCall
	ToolDef      = openai.Tool
	FinishReason = openai.FinishReason
)

const (
	RoleSystem    = openai.ChatMessageRoleSystem
	RoleUser      = openai.ChatMessageRoleUser
	RoleAssistant = openai.ChatMessageRoleAssistant
	RoleTool      = openai.ChatMessageRoleTool

	ToolTypeFunction = openai.ToolTypeFunction

	FinishReasonStop      = openai.FinishReasonStop
	FinishReasonToolCalls = openai.FinishReasonToolCalls
)

// NewFunctionTool builds a tool-calling definition for a function-type tool
// from a JSON Schema describing its parameters.
func NewFunctionTool(name, description string, parameters json.RawMessage) ToolDef {
	return openai.Tool{
		Type: ToolTypeFunction,
		Function: &openai.FunctionDefinition{
			Name:        name,
			Description: description,
			Parameters:  parameters,
		},
	}
}

// ChatCompletionWithTools calls the chat completion API with a full message
// history and an optional list of tool definitions, returning the model's
// message and why it stopped generating (e.g. it wants to call a tool, or
// it's done and the message is the final answer).
func ChatCompletionWithTools(ctx context.Context, model LLMModel, messages []Message, tools []ToolDef) (Message, FinishReason, error) {
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:       string(model),
		Messages:    messages,
		Tools:       tools,
		Temperature: 0.2,
	})
	if err != nil {
		return Message{}, "", err
	}
	choice := resp.Choices[0]
	return choice.Message, choice.FinishReason, nil
}