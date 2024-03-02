package config

import "os"

type ModelConfig struct {
	UserInput string
	ApiKey    string
	AssistUrl string
	ThreadUrl string
	Model     string
}

func NewModelConfigs() map[string]ModelConfig {
	return map[string]ModelConfig{
		"gpt4": {
			UserInput: "your_topic_here",           // Replace with your actual input
			ApiKey:    os.Getenv("OPENAI_API_KEY"), // Ensure you have set your API key in your environment variables
			AssistUrl: "https://api.openai.com/v1/assistants",
			ThreadUrl: "https://api.openai.com/v1/threads",
			Model:     "gpt-4", //"gpt-3.5-turbo",
		},
		"ollama": {
			AssistUrl: "http://localhost:11434/api/generate",
			Model:     "llama2", //"gpt-3.5-turbo",
		},
	}
}
