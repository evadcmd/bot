package tool

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

type closeTrackingBody struct {
	io.Reader
	closed *bool
}

func (b *closeTrackingBody) Close() error {
	*b.closed = true
	return nil
}

type stubTransport struct {
	statusCode int
	closed     *bool
}

func (s *stubTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: s.statusCode,
		Body:       &closeTrackingBody{Reader: strings.NewReader(`{"error":"quota exceeded"}`), closed: s.closed},
		Header:     make(http.Header),
	}, nil
}

// Regression test for a bug where the response body was only closed on
// the happy path. http.DefaultClient has no Transport set, so it falls
// back to http.DefaultTransport, which this test swaps out to observe
// whether Search() closes the body on a non-200 response.
func TestSearchClosesResponseBodyOnErrorStatus(t *testing.T) {
	closed := false
	originalTransport := http.DefaultTransport
	http.DefaultTransport = &stubTransport{statusCode: http.StatusTooManyRequests, closed: &closed}
	defer func() { http.DefaultTransport = originalTransport }()

	ws := NewWebSearch()
	if _, err := ws.Search(context.Background(), "ai"); err == nil {
		t.Fatal("expected Search to return an error for a non-200 response")
	}
	if !closed {
		t.Error("Search did not close the response body on a non-200 status, leaking the underlying connection")
	}
}