package openai

import "github.com/sashabaranov/go-openai"

type LLMModel string

// copy from github.com/sashabaranov/go-openai to regulate dependencies
// should be maintained while version of github.com/sashabaranov/go-openai is upgraded
const (
	O1Mini                = LLMModel(openai.O1Mini)
	O1Mini20240912        = LLMModel(openai.O1Mini20240912)
	O1Preview             = LLMModel(openai.O1Preview)
	O1Preview20240912     = LLMModel(openai.O1Preview20240912)
	GPT432K0613           = LLMModel(openai.GPT432K0613)
	GPT432K0314           = LLMModel(openai.GPT432K0314)
	GPT432K               = LLMModel(openai.GPT432K)
	GPT40613              = LLMModel(openai.GPT40613)
	GPT40314              = LLMModel(openai.GPT40314)
	GPT4o                 = LLMModel(openai.GPT4o)
	GPT4o20240513         = LLMModel(openai.GPT4o20240513)
	GPT4o20240806         = LLMModel(openai.GPT4o20240806)
	GPT4oLatest           = LLMModel(openai.GPT4oLatest)
	GPT4oMini             = LLMModel(openai.GPT4oMini)
	GPT4oMini20240718     = LLMModel(openai.GPT4oMini20240718)
	GPT4Turbo             = LLMModel(openai.GPT4Turbo)
	GPT4Turbo20240409     = LLMModel(openai.GPT4Turbo20240409)
	GPT4Turbo0125         = LLMModel(openai.GPT4Turbo0125)
	GPT4Turbo1106         = LLMModel(openai.GPT4Turbo1106)
	GPT4TurboPreview      = LLMModel(openai.GPT4TurboPreview)
	GPT4VisionPreview     = LLMModel(openai.GPT4VisionPreview)
	GPT4                  = LLMModel(openai.GPT4)
	GPT3Dot5Turbo0125     = LLMModel(openai.GPT3Dot5Turbo0125)
	GPT3Dot5Turbo1106     = LLMModel(openai.GPT3Dot5Turbo1106)
	GPT3Dot5Turbo0613     = LLMModel(openai.GPT3Dot5Turbo0613)
	GPT3Dot5Turbo0301     = LLMModel(openai.GPT3Dot5Turbo0301)
	GPT3Dot5Turbo16K      = LLMModel(openai.GPT3Dot5Turbo16K)
	GPT3Dot5Turbo16K0613  = LLMModel(openai.GPT3Dot5Turbo16K0613)
	GPT3Dot5Turbo         = LLMModel(openai.GPT3Dot5Turbo)
	GPT3Dot5TurboInstruct = LLMModel(openai.GPT3Dot5TurboInstruct)
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextDavinci003 = LLMModel(openai.GPT3TextDavinci003)
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextDavinci002 = LLMModel(openai.GPT3TextDavinci002)
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextCurie001 = LLMModel(openai.GPT3TextCurie001)
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextBabbage001 = LLMModel(openai.GPT3TextBabbage001)
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextAda001 = LLMModel(openai.GPT3TextAda001)
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3TextDavinci001 = LLMModel(openai.GPT3TextDavinci001)
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3DavinciInstructBeta = LLMModel(openai.GPT3DavinciInstructBeta)
	// Deprecated: Model is shutdown. Use davinci-002 instead.
	GPT3Davinci    = LLMModel(openai.GPT3Davinci)
	GPT3Davinci002 = LLMModel(openai.GPT3Davinci002)
	// Deprecated: Model is shutdown. Use gpt-3.5-turbo-instruct instead.
	GPT3CurieInstructBeta = LLMModel(openai.GPT3CurieInstructBeta)
	GPT3Curie             = LLMModel(openai.GPT3Curie)
	GPT3Curie002          = LLMModel(openai.GPT3Curie002)
	// Deprecated: Model is shutdown. Use babbage-002 instead.
	GPT3Ada    = LLMModel(openai.GPT3Ada)
	GPT3Ada002 = LLMModel(openai.GPT3Ada002)
	// Deprecated: Model is shutdown. Use babbage-002 instead.
	GPT3Babbage    = LLMModel(openai.GPT3Babbage)
	GPT3Babbage002 = LLMModel(openai.GPT3Babbage002)
)
