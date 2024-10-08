package tool

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/evadcmd/bot/internal/util"
	"github.com/joho/godotenv"
)

type cseErr struct {
	Domain  string
	Message string
}

type cseErrs struct {
	Errors []cseErr
}

type item struct {
	Title   string
	Link    string
	Snippet string
}

type cseResp struct {
	Items  []item
	Errors *cseErrs
}

type WebSearchTool struct {
	url     string
	api_key string
	cse_id  string
	safe    bool
}

func (*WebSearchTool) GetName() string {
	return "WebSearch"
}

func (ws *WebSearchTool) GetDescription() string {
	return "A tool for obtaining information from the internet"
}

func (*WebSearchTool) GetInputFmt() string {
	return "a string format search query"
}

func (ws *WebSearchTool) Search(ctx context.Context, query string) (string, error) {
	safe := "active"
	if !ws.safe {
		safe = "off"
	}
	url := fmt.Sprintf("%s?key=%s&cx=%s&q=%s&safe=%s", ws.url, ws.api_key, ws.cse_id, url.QueryEscape(query), safe)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to build a cse request: %w", err)
	}
	slog.Info("perform an google custom search: " + query)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send a cse request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		data, err := io.ReadAll(res.Body)
		if err == nil {
			err = errors.New(string(data))
		}
		return "", fmt.Errorf("get an error response from google custom search: %w", err)
	}
	defer res.Body.Close()

	var resp cseResp
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return "", fmt.Errorf("failed to unmarshal the cse response: %w", err)
	}
	if resp.Errors != nil {
		return "", fmt.Errorf("get an error response from google custom search: %+v", resp.Errors)
	}
	var items []string
	for i, item := range resp.Items {
		items = append(items, fmt.Sprintf("(%d) %s %s %s", i, item.Title, item.Snippet, item.Link))
	}
	return strings.Join(items, "	"), nil
}

func NewWebSearch() *WebSearchTool {
	return &WebSearchTool{
		url:     os.Getenv("GOOGLESEARCH_BASE_URL"),
		api_key: os.Getenv("GOOGLESEARCH_API_KEY"),
		cse_id:  os.Getenv("GOOGLESEARCH_CSE_ID"),
		safe:    true,
	}
}

func init() {
	godotenv.Load(path.Join(util.RootPath, ".env"))
}
