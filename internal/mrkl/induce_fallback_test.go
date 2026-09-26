package mrkl

import (
	"context"
	"strings"
	"testing"

	"github.com/evadcmd/bot/internal/llm/openai"
)

// Regression test for a bug where Induce's fallback re-asked the raw
// question (q) instead of the accumulated prompt (template scaffolding +
// any tool observations gathered so far), silently discarding context.
func TestInduceFallbackUsesAccumulatedPrompt(t *testing.T) {
	original := chatCompletion
	defer func() { chatCompletion = original }()

	type call struct {
		model      openai.LLMModel
		userPrompt string
	}
	var calls []call
	chatCompletion = func(ctx context.Context, model openai.LLMModel, userPrompt string, stop []string) (string, error) {
		calls = append(calls, call{model: model, userPrompt: userPrompt})
		// Matches neither finishRegex nor actionRegex, forcing Induce to
		// fall through to the final fallback after a single iteration.
		return "I cannot determine the next step.", nil
	}

	q := "what is the weather today"
	if _, err := Induce(context.Background(), q); err != nil {
		t.Fatalf("Induce returned an unexpected error: %v", err)
	}

	if len(calls) != 2 {
		t.Fatalf("expected the selector call plus one fallback call, got %d calls", len(calls))
	}

	fallbackPrompt := calls[1].userPrompt
	if fallbackPrompt == q {
		t.Fatal("fallback discarded the accumulated MRKL context and reused the raw question")
	}
	if !strings.Contains(fallbackPrompt, "Rigorously adhere") || !strings.Contains(fallbackPrompt, q) {
		t.Fatalf("fallback prompt is missing the MRKL scaffolding built up during Induce: %q", fallbackPrompt)
	}
}