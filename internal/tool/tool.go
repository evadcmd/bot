package tool

import "encoding/json"

// Tool describes something the MRKL agent can call via OpenAI's native
// tool-calling API.
type Tool interface {
	GetName() string
	GetDescription() string
	// GetParameters returns the JSON Schema for the tool's input, in the
	// shape OpenAI's tool-calling API expects for a function definition.
	GetParameters() json.RawMessage
}