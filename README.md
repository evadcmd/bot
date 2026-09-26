# bot

An experimental MRKL (Modular Reasoning, Knowledge and Language) agent exposed as an HTTP API, built with [gofiber](https://github.com/gofiber/fiber).

The service takes a natural language question, runs a ReAct-style reasoning loop against an LLM (currently OpenAI), lets the model call tools (web search, current datetime) when it needs external information, and returns a final answer.

## How it works

1. A request comes in through `POST /api/v0/mrkl` with the user's question.
2. The question is rendered into a MRKL prompt template (`internal/mrkl/mrkl.tpl`) that lists the available tools and instructs the model to respond with `Thought` / `Action` / `Action Input` / `Final Answer` blocks.
3. `internal/mrkl.Induce` sends the prompt to a cheaper "selector" model to decide the next action, parses the tool call out of the response, executes the tool, and feeds the observation back into the prompt. This repeats for up to 10 steps.
4. Once a `Final Answer` is produced (or the loop ends without one), the full transcript is sent to a stronger "answerer" model to produce the polished final reply.

Available tools live under `internal/tool/`:

- **Datetime** — returns the current date and time.
- **WebSearch** — queries the Google Custom Search API for external information.

> Note: the agent currently relies on regex-parsed text output (the classic 2023-era MRKL prompting style) rather than native LLM tool-calling. A migration to structured tool calls is planned.

## Requirements

- Go 1.23+
- An OpenAI API key
- (Optional, for web search) A Google Custom Search API key and Search Engine ID

## Configuration

The app loads a `.env` file from the project root at startup (via [godotenv](https://github.com/joho/godotenv)). Create one with:

```bash
OPENAI_API_KEY=sk-...
GOOGLESEARCH_BASE_URL=https://www.googleapis.com/customsearch/v1
GOOGLESEARCH_API_KEY=...
GOOGLESEARCH_CSE_ID=...
```

`GOOGLESEARCH_*` variables are only needed if you want the WebSearch tool to work; without them, that tool will simply fail to return results.

## Running the server

```bash
go run ./cmd/server
```

The server listens on `:5252` and exposes:

| Method | Path            | Description                |
|--------|-----------------|-----------------------------|
| GET    | `/healthz`      | Kubernetes readiness probe  |
| POST   | `/api/v0/mrkl`  | Ask the agent a question    |
| GET    | `/swagger/*`    | Swagger UI / API docs       |

### Example request

```bash
curl -X POST http://localhost:5252/api/v0/mrkl \
  -H "Content-Type: application/json" \
  -d '{"content": "What is the weather like in Tokyo right now?"}'
```

Response:

```json
{
  "content": "..."
}
```

## API docs (Swagger)

Swagger annotations live alongside the handlers in `pkg/api/`. To regenerate the docs after changing them:

```bash
# one-time setup
go get github.com/gofiber/contrib/swagger
go install github.com/swaggo/swag/cmd/swag@latest

# regenerate docs/
swag init -g cmd/server/main.go --parseInternal
```

The generated docs are served at `/swagger/index.html` when the server is running.

## Project layout

```
cmd/server/          entry point (fiber app, routes)
pkg/api/             HTTP handlers
pkg/dto/             request/response payloads
internal/mrkl/       MRKL reasoning loop and prompt template
internal/llm/openai/ OpenAI chat completion client
internal/tool/       tools available to the agent (datetime, web search)
internal/util/       shared helpers
docs/                generated Swagger docs
```

## Testing

```bash
go test ./...
```