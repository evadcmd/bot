package openai

import (
	"context"
	"os"
	"path"
	"testing"

	"github.com/evadcmd/bot/internal/util"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func TestClient(t *testing.T) {
	res, err := ChatCompletion(context.TODO(), GPT3Dot5Turbo, "hi", nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestMain(m *testing.M) {
	godotenv.Load(path.Join(util.RootPath, ".env"))
	client = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	m.Run()
}
