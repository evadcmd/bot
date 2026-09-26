package mrkl

import (
	"context"
	"testing"

	"github.com/evadcmd/bot/internal/llm/openai"
)

// Regression test: Induce must hand the answerer model the full accumulated
// message history (including tool call/result turns), not just the
// original question, when finalizing. This is the tool-calling equivalent
// of a bug this loop used to have when it was driven by regex-parsed text:
// the fallback path silently discarded everything gathered so far.
func TestInduceFinalizeIncludesToolHistory(t *testing.T) {
	original := chatCompletionWithTools
	defer func() { chatCompletionWithTools = original }()

	type call struct {
		model    openai.LLMModel
		messages []openai.Message
	}
	var calls []call
	first := true
	chatCompletionWithTools = func(ctx context.Context, model openai.LLMModel, messages []openai.Message, tools []openai.ToolDef) (openai.Message, openai.FinishReason, error) {
		calls = append(calls, call{model: model, messages: messages})

		if first {
			first = false
			return openai.Message{
				Role: openai.RoleAssistant,
				ToolCalls: []openai.ToolCall{
					{
						ID:   "call_1",
						Type: openai.ToolTypeFunction,
						Function: openai.FunctionCall{
							Name:      "Datetime",
							Arguments: "{}",
						},
					},
				},
			}, openai.FinishReasonToolCalls, nil
		}
		return openai.Message{Role: openai.RoleAssistant, Content: "the answer"}, openai.FinishReasonStop, nil
	}

	got, err := Induce(context.Background(), "what time is it")
	if err != nil {
		t.Fatalf("Induce returned an unexpected error: %v", err)
	}
	if got != "the answer" {
		t.Fatalf("expected final content %q, got %q", "the answer", got)
	}
	if len(calls) != 2 {
		t.Fatalf("expected a tool-call round plus a finalize call, got %d calls", len(calls))
	}

	finalizeCall := calls[1]
	if finalizeCall.model != answerer {
		t.Errorf("finalize should use the answerer model, got %v", finalizeCall.model)
	}
	var sawToolResult bool
	for _, m := range finalizeCall.messages {
		if m.Role == openai.RoleTool {
			sawToolResult = true
		}
	}
	if !sawToolResult {
		t.Fatal("finalize call is missing the tool result gathered during the loop — Induce is discarding accumulated context")
	}
}