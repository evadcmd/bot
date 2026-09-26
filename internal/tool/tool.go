package tool

import (
	"context"
	"encoding/json"
)

type Tool interface {
	GetName() string
	GetDescription() string
	GetInputFmt() string
}

// Callable is implemented by tools that can be invoked through a model's
// native tool-calling API. Schema returns the JSON Schema describing the
// call arguments; Call executes the tool given the model-supplied JSON
// arguments and returns the observation to feed back to the model.
type Callable interface {
	Tool
	Schema() json.RawMessage
	Call(ctx context.Context, argumentsJSON string) (string, error)
}