package mrkl

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"path"
	"regexp"

	"text/template"

	"github.com/evadcmd/bot/internal/llm/openai"
	"github.com/evadcmd/bot/internal/tool"
	"github.com/evadcmd/bot/internal/util"
	"github.com/gofiber/fiber/v2/log"
)

var tools = []tool.Tool{&tool.DatetimeTool{}, tool.NewWebSearch()}
var nameToTool map[string]tool.Tool

var finishRegex = regexp.MustCompile(`Final Answer\s*:\s*`)
var actionRegex = regexp.MustCompile(`Action\s*:\s*(?P<action>.*)\s*Action\s*Input\s*:\s*(?P<actionInput>.*)`)
var mrklTemplate = template.Must(template.ParseFiles(path.Join(util.RootPath, "/internal/mrkl/mrkl.tpl")))
var stopFlags = []string{"Observation"}

var selector = openai.GPT3Dot5Turbo1106
var answerer = openai.GPT4

type mrklParam struct {
	Tools []tool.Tool
	Input string
}

func init() {
	nameToTool = make(map[string]tool.Tool)
	for _, tool := range tools {
		nameToTool[tool.GetName()] = tool
	}
}

func Induce(ctx context.Context, q string) (string, error) {
	var mrklTplBytes bytes.Buffer
	if err := mrklTemplate.Execute(&mrklTplBytes, mrklParam{Tools: tools, Input: q}); err != nil {
		return "", fmt.Errorf("failed to execute the MRKL template: %w", err)
	}
	prompt := mrklTplBytes.String()
	for range 10 {
		res, err := openai.ChatCompletion(ctx, selector, prompt, stopFlags)
		if err != nil {
			return "", fmt.Errorf("failed to send a request to OpenAI API server: %w", err)
		}
		slog.Info(res)

		if idx := finishRegex.FindStringSubmatchIndex(res); len(idx) != -1 {
			prompt += res[:idx[1]]
			slog.Info(prompt)
			return openai.ChatCompletion(ctx, answerer, prompt, nil)
		} else if idx = actionRegex.FindStringSubmatchIndex(res); len(idx) != -1 {
			name, input := res[idx[2]:idx[3]], res[idx[4]:idx[5]]
			slog.Info("use tool", name, input)
			tl, ok := nameToTool[name]
			if !ok {
				return "", fmt.Errorf("failed to get the tool name and input from LLM's action text block: %s %s", name, input)
			}
			var observation string
			switch tool := tl.(type) {
			case *tool.DatetimeTool:
				observation = tool.Now()
			case *tool.WebSearchTool:
				if observation, err = tool.Search(ctx, input); err != nil {
					return "", fmt.Errorf("failed to execute WebSearch tool: %w", err)
				}
			}
			slog.Info(observation)
			prompt += (res + "Observation: " + observation + "\n")
		} else {
			log.Warn("failed to parse the MKRL template")
			break
		}
	}
	return openai.ChatCompletion(ctx, answerer, q, nil)
}
