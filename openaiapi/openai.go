package openaiapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Define structs to match the response structure
type OpenAIResponse struct {
	ID       string  `json:"id"`
	Object   string  `json:"object"`
	Created  int64   `json:"created"`
	Model    string  `json:"model"`
	Choices  []Choice `json:"choices"`
	Usage    Usage   `json:"usage"`
}

type Choice struct {
	Index         int    `json:"index"`
	Message       Message `json:"message"`
	Logprobs      interface{} `json:"logprobs"`
	FinishReason  string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens    int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens     int `json:"total_tokens"`
}

// Define a struct to hold the request payload
type CompletionRequest struct {
	Model    string `json:"model"`
	Messages []Message `json:"messages"`
}
type RequestBody struct{
	Model string `json:"model"`  //ID of the model to use. You can use the List models API to see all of your available models, or see our Model overview for descriptions of them.
	Name string `json:"name"`  //The name of the assistant. The maximum length is 256 characters.
	Description string `json:"description"` //The description of the assistant. The maximum length is 512 characters.
	Instructions string `json:"instructions"` //The system instructions that the assistant uses. The maximum length is 32768 characters.
	Tools []Tool `json:"tootls"`  //A list of tool enabled on the assistant. There can be a maximum of 128 tools per assistant. Tools can be of types code_interpreter, retrieval, or function.
	File_ids []string `json:"file_ids"`//A list of file IDs attached to this assistant. There can be a maximum of 20 files attached to the assistant. Files are ordered by their creation date in ascending order.
	Metadata map[string]string `json:"metadata"`//Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
}
type Tool struct{
	CodeInterpreter string `json:"code_interpreter"`
	Retrieval string `json:"retrieval"`
	Function string `json:"function"`
}
type AsistantsResponse struct{
	Id string `json:"id"`	//The identifier, which can be referenced in API endpoints.
	Object string `json:"object"` //The object type, which is always assistant.
	Created_at int64 `json:"created_at"` //The Unix timestamp (in seconds) for when the assistant was created.
	Name string `json:"name"` //The name of the assistant. The maximum length is 256 characters.
	Description string `json:"description"` //The description of the assistant. The maximum length is 512 characters.
	Model string `json:"model"` //ID of the model to use. You can use the List models API to see all of your available models, or see our Model overview for descriptions of them.
	Instructions string `json:"instruction"` //The system instructions that the assistant uses. The maximum length is 32768 characters.
	Tools []Tool `json:"tools"` //A list of tool enabled on the assistant. There can be a maximum of 128 tools per assistant. Tools can be of types code_interpreter, retrieval, or function.
	File_ids string `json:"file_ids"`//A list of file IDs attached to this assistant. There can be a maximum of 20 files attached to the assistant. Files are ordered by their creation date in ascending order.
	Metadata map[string]string `json:"metadata"`//Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
}
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAI struct{
	ApiKey string
	Url string
	Model string  
}

func NewOpenAI(apiKey, url, model string)OpenAI{
	return OpenAI{
		ApiKey: apiKey,
		Url: url, 
		Model: model,
	}
}

func (o *OpenAI)ApiFetch(Messages []Message)(string, error){
	requestBody := CompletionRequest{
		Model:    o.Model,
		Messages: Messages,
	}

	// Marshal the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshalling the request body: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", o.Url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("error creating the request: %v", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+o.ApiKey)
	req.Header.Set("OpenAI-Beta", "assistants=v1")
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making the request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading the response body: %v", err)
	}
	// fmt.Println("yes=", string(responseBody))
	// Parse the JSON response
	var response OpenAIResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return "", fmt.Errorf("error parsing the JSON response: %v", err)
	}

	// Assuming you're interested in the content of the first choice
	if len(response.Choices) > 0 {
		generatedContent := response.Choices[0].Message.Content
		return generatedContent, nil
	} else {
		return "", fmt.Errorf("no content generated")
	}
}
