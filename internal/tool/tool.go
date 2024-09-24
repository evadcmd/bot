package tool

type Tool interface {
	GetName() string
	GetDescription() string
	GetInputFmt() string
}
