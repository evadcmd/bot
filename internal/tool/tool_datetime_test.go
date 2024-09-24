package tool

import (
	"fmt"
	"testing"
)

func TestDatetimeTool(t *testing.T) {
	var x Tool = &DatetimeTool{}
	dt, ok := x.(*DatetimeTool)
	if !ok {
		t.Error("failed to cast tool to DatetimeTool")
	}
	fmt.Println(dt.Now())
}
