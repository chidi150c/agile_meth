package ai_model

// AIModel represents an interface for interacting with an AI model.
type AIModelServicer interface {
	PromptAI(system, user string) (string, error)
}

// AIConfig represents the configuration for accessing the AI model.
type AIConfig struct {
	APIKey string
	URL    string
}

// RealAIModel represents the real implementation of an AI model.
type RealAIModel struct {
	Config AIConfig
}