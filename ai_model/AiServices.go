package ai_model


type AIServicer interface {
	ProcessAiMessage(msg string) (string, error)
	CreateAssistant(inst, name, tYpe string) (string, error)
	CreateThread() (string, error)
}