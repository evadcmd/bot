package mrkl

import "testing"

func TestToolDefsMatchTools(t *testing.T) {
	if len(toolDefs) != len(tools) {
		t.Fatalf("expected %d tool definitions, got %d", len(tools), len(toolDefs))
	}
	for i, td := range toolDefs {
		if td.Function == nil {
			t.Fatalf("tool definition %d has no function definition", i)
		}
		if td.Function.Name != tools[i].GetName() {
			t.Errorf("expected tool definition %d name %q, got %q", i, tools[i].GetName(), td.Function.Name)
		}
		if len(td.Function.Parameters) == 0 {
			t.Errorf("tool definition %q has empty parameters schema", td.Function.Name)
		}
	}
}

func TestNameToToolIsPopulated(t *testing.T) {
	for _, tl := range tools {
		if _, ok := nameToTool[tl.GetName()]; !ok {
			t.Errorf("nameToTool is missing an entry for %q", tl.GetName())
		}
	}
}