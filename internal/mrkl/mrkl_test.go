package mrkl

import (
	"bytes"
	"os"
	"testing"
)

func TestExecTemplate(t *testing.T) {
	if err := mrklTemplate.Execute(
		os.Stdout,
		mrklParam{
			Tools: tools,
			Input: "content"}); err != nil {
		t.Error(err)
	}
}

func TestRegexGroup(t *testing.T) {
	var buffer bytes.Buffer
	if err := mrklTemplate.Execute(
		&buffer,
		mrklParam{
			Tools: tools,
			Input: "content"}); err != nil {
		t.Error(err)
	}
	tpl := buffer.String()
	t.Log(tpl)
	match := actionRegex.FindStringSubmatch(tpl)
	groupNames := actionRegex.SubexpNames()
	groupIdx := actionRegex.SubexpIndex("action")
	t.Logf("%+v %+v %+v", match, groupNames, groupIdx)

	idx := actionRegex.FindStringSubmatchIndex(tpl)
	t.Logf(`%+v\n%+v\n%+v`, tpl[idx[0]:idx[1]], tpl[idx[2]:idx[3]], tpl[idx[4]:idx[5]])
}
