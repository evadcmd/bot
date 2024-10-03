package tool

import (
	"context"
	"testing"
)

func TestWebSearch(t *testing.T) {
	ws := NewWebSearch()
	res, err := ws.Search(context.TODO(), "ai")
	if err != nil {
		t.Errorf("failed to fetch info: %s", err)
	}
	t.Log(res)
}
