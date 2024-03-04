package ai_model

type AIServicer interface {
	ProcessAiMessage(msg string) (string, error)
}