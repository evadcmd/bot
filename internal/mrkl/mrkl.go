package mrkl

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/evadcmd/bot/internal/llm/openai"
	"github.com/evadcmd/bot/internal/tool"
)

const maxIterations = 10

const systemPrompt = "You are a versatile assistant with access to tools. " +
	"Use them when they help answer the question. " +
	"Respond in Japanese regardless of the language of tool outputs or intermediate reasoning."

var tools = []tool.Tool{&tool.DatetimeTool{}, tool.NewWebSearch()}
var nameToTool map[string]tool.Tool
var toolDefs []openai.ToolDef

var selector = openai.GPT3Dot5Turbo1106
var answerer = openai.GPT4

// chatCompletionWithTools is a seam over openai.ChatCompletionWithTools so tests can stub it.
var chatCompletionWithTools = openai.ChatCompletionWithTools

func init() {
	nameToTool = make(map[string]tool.Tool)
	for _, t := range tools {
		nameToTool[t.GetName()] = t
		toolDefs = append(toolDefs, openai.NewFunctionTool(t.GetName(), t.GetDescription(), t.GetParameters()))
	}
}

// Induce runs a tool-calling agent loop: the selector model decides whether
// to answer directly or call a tool, tool results are fed back as message
// history, and once the selector is done the answerer model produces the
// final response from the full accumulated conversation.
func Induce(ctx context.Context, q string) (string, error) {
	messages := []openai.Message{
		{Role: openai.RoleSystem, Content: systemPrompt},
		{Role: openai.RoleUser, Content: q},
	}

	for range maxIterations {
		msg, finishReason, err := chatCompletionWithTools(ctx, selector, messages, toolDefs)
		if err != nil {
			return "", fmt.Errorf("failed to send a request to OpenAI API server: %w", err)
		}
		messages = append(messages, msg)

		if finishReason != openai.FinishReasonToolCalls {
			return finalize(ctx, messages)
		}

		for _, tc := range msg.ToolCalls {
			observation, err := runTool(ctx, tc)
			if err != nil {
				return "", err
			}
			slog.Info("tool call", "name", tc.Function.Name, "input", tc.Function.Arguments, "observation", observation)
			messages = append(messages, openai.Message{
				Role:       openai.RoleTool,
				ToolCallID: tc.ID,
				Content:    observation,
			})
		}
	}

	slog.Warn("MRKL reached the maximum number of tool-call iterations")
	return finalize(ctx, messages)
}

// finalize hands the full accumulated conversation, including every tool
// call and result gathered so far, to the answerer model for a final
// response — nothing gathered during the loop is discarded.
func finalize(ctx context.Context, messages []openai.Message) (string, error) {
	msg, _, err := chatCompletionWithTools(ctx, answerer, messages, nil)
	if err != nil {
		return "", fmt.Errorf("failed to send a request to OpenAI API server: %w", err)
	}
	return msg.Content, nil
}

func runTool(ctx context.Context, tc openai.ToolCall) (string, error) {
	tl, ok := nameToTool[tc.Function.Name]
	if !ok {
		return "", fmt.Errorf("unknown tool requested by the model: %q", tc.Function.Name)
	}
	switch t := tl.(type) {
	case *tool.DatetimeTool:
		return t.Now(), nil
	case *tool.WebSearchTool:
		var args struct {
			Query string `json:"query"`
		}
		if err := json.Unmarshal([]byte(tc.Function.Arguments), &args); err != nil {
			return "", fmt.Errorf("failed to parse arguments for tool %q: %w", tc.Function.Name, err)
		}
		observation, err := t.Search(ctx, args.Query)
		if err != nil {
			return "", fmt.Errorf("failed to execute WebSearch tool: %w", err)
		}
		return observation, nil
	default:
		return "", fmt.Errorf("unknown tool type for tool %q: %T", tc.Function.Name, t)
	}
}