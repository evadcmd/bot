package tool

import (
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
