package tool

import (
	"context"
	"encoding/json"
	"time"
)

type DatetimeTool struct{}

func (*DatetimeTool) GetName() string {
	return "Datetime"
}

func (dt *DatetimeTool) GetDescription() string {
	return `A tool returns current datetime in "{YEAR}-{MONTH}-{DAY} {HOUR}:{MINUTE}:{SECOND}" format`
}

func (*DatetimeTool) GetInputFmt() string {
	return "no input parameter is required"
}

func (*DatetimeTool) Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (*DatetimeTool) Schema() json.RawMessage {
	return json.RawMessage(`{"type":"object","properties":{}}`)
}

func (dt *DatetimeTool) Call(_ context.Context, _ string) (string, error) {
	return dt.Now(), nil
}