package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Define a struct to hold the request payload

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type requestMessage struct{
	Model string `json:"model"`
	Stream bool `json:"stream"`
	Messages []Message `json:"messages"`
	Prompt string `json:"prompt"`
}
type Response struct {
    Model      string    `json:"model"`
    CreatedAt  time.Time `json:"created_at"`
    Response   string    `json:"response"` 
    Done       bool      `json:"done"`
}
// Ollama holds configuration for making requests to the Ollama API.
// ApiKey is the API key to use for authentication.
// Url is the base URL of the Ollama API.
// Model is the name of the AI model to use.
// Messages is a slice of Message structs to send to the API.
type Ollama struct {
	ApiKey   string
	Url      string
	Model    string 
	SystemMessage string
	Messages []Message
}
func NewOllama(url, model string)*Ollama{
	return &Ollama{
		Url: url, 
		Model: model,
	}
}

func (a *Ollama)ProcessAiMessage(msg string) (string, error){
	messages := requestMessage{
		Model: "llama2",
		Prompt: a.SystemMessage+" "+msg,
		Stream: false,
	}
	resp, err := apiFetch(messages, a.Model, a.Url, a.ApiKey)
	if err != nil {
		log.Fatal(err)
	}
	// // Parse the JSON response
	var response Response
	if err := json.Unmarshal(resp, &response); err != nil {
		return "", fmt.Errorf("error parsing the JSON response: %v", err)
	}
	return response.Response, nil
}
func (a *Ollama)CreateAssistant(inst, name, tYpe string) (string, error){
	return "", nil
}
func (a *Ollama)CreateThread() (string, error){
	return "", nil
}


func apiFetch(msg interface{}, Model, Url, ApiKey string)([]byte, error){
	// Marshal the request body to JSON
	jsonBody, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("error marshalling the request body: %v", err)
	}
	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
