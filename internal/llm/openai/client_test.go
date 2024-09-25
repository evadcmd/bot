package openai

import (
	"context"
	"testing"
)

func TestClient(t *testing.T) {
	res, err := ChatCompletion(context.TODO(), GPT3Dot5Turbo, "hi", nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
