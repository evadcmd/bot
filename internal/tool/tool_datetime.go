package tool

import (
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

func (*DatetimeTool) GetParameters() json.RawMessage {
	return json.RawMessage(`{"type":"object","properties":{}}`)
}

func (*DatetimeTool) Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}