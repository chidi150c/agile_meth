package openaiapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Tool struct {
	Type        string      `json:"type"`                  // The type of tool being defined: code_interpreter, retrieval, or function
	Description string      `json:"description,omitempty"` // A description of what the function does (only for function type)
	Name        string      `json:"name,omitempty"`        // The name of the function to be called (only for function type)
	Parameters  interface{} `json:"parameters,omitempty"`  // The parameters the function accepts (only for function type)
}

type RequestBody struct {
	Model        string      `json:"model"`              // ID of the model to use
	Name         string      `json:"name"`               // The name of the assistant
	Description  string      `json:"description"`        // The description of the assistant
	Instructions string      `json:"instructions"`       // The system instructions that the assistant uses
	Tools        []Tool      `json:"tools,omitempty"`    // A list of tools enabled on the assistant
	FileIDs      []string    `json:"file_ids,omitempty"` // A list of file IDs attached to this assistant
	Metadata     interface{} `json:"metadata,omitempty"` // Set of key-value pairs attached to the assistant
}

type AsistantsResponse struct {
	Id           string   `json:"id"`          //The identifier, which can be referenced in API endpoints.
	Object       string   `json:"object"`      //The object type, which is always assistant.
	Created_at   int64    `json:"created_at"`  //The Unix timestamp (in seconds) for when the assistant was created.
	Name         string   `json:"name"`        //The name of the assistant. The maximum length is 256 characters.
	Description  string   `json:"description"` //The description of the assistant. The maximum length is 512 characters.
	Model        string   `json:"model"`       //ID of the model to use. You can use the List models API to see all of your available models, or see our Model overview for descriptions of them.
	Instructions string   `json:"instruction"` //The system instructions that the assistant uses. The maximum length is 32768 characters.
	Tools        []Tool   `json:"tools"`       //A list of tool enabled on the assistant. There can be a maximum of 128 tools per assistant. Tools can be of types code_interpreter, retrieval, or function.
	File_ids     []string `json:"file_ids"`    //A list of file IDs attached to this assistant. There can be a maximum of 20 files attached to the assistant. Files are ordered by their creation date in ascending order.
	Metadata     struct{} `json:"metadata"`    //Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long.
}

type OpenAI struct {
	ApiKey    string
	AssistUrl string
	ThreadUrl string
	RunUrl    string
	Model     string
}

func NewOpenAI(apiKey, assistUrl, threadUrl, runUrl, model string) OpenAI {
	return OpenAI{
		ApiKey:    apiKey,
		AssistUrl: assistUrl,
		Model:     model,
	}
}
func (o *OpenAI) Assistants(inst, name, tYpe string) (string, error) {
	requestBody := RequestBody{
		Instructions: inst,
		Name:         name,
		Tools: []Tool{{
			Type: "code_interpreter",
		}},
		Model: "gpt-4",
	}
	// Marshal the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshalling the request body: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", o.AssistUrl, bytes.NewBuffer(jsonBody))
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
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("assistant Error: Unexpected status code: %v", resp.StatusCode)
	}

	// Read the response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading the response body: %v", err)
	}
	// Parse the JSON response
	var response AsistantsResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return "", fmt.Errorf("error parsing the JSON response: %v", err)
	}
	fmt.Printf("Assistant_id = %v", response.Id)
	return response.Id, nil
}

type Thread struct {
	ID        string            `json:"id"`
	Object    string            `json:"object"`
	CreatedAt int64             `json:"created_at"`
	Metadata  map[string]string `json:"metadata"`
}

// ThreadResponse represents the structure of the response body
type ThreadResponse struct {
	ID        string            `json:"id"`
	Object    string            `json:"object"`
	CreatedAt int64             `json:"created_at"`
	Metadata  map[string]string `json:"metadata"`
}

// RequestBody represents the request body structure
type ThreadRequest struct {
	Messages []string          `json:"messages,omitempty"` // A list of messages to start the thread with
	Metadata map[string]string `json:"metadata,omitempty"` // Set of key-value pairs for additional information
}

func (o *OpenAI) CreateThread() (string, error) {
	// Define the request body
	requestBody := []byte("")

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/threads", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("in Thread Error creating request: %v", err)
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
	req.Header.Set("OpenAI-Beta", "assistants=v1")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}

	// Print response body
	responseBody := new(bytes.Buffer)
	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body: %v", err)
	}
	fmt.Println("Thread Response body:", responseBody.String())
	return responseBody.String(), nil
}

// func retrieveThread(threadID string) (Thread, error) {
// 	url := fmt.Sprintf("https://api.openai.com/v1/threads/%s", threadID)
// 	headers := map[string]string{
// 		"Authorization": "Bearer YOUR_OPENAI_API_KEY",
// 		"OpenAI-Beta":   "assistants=v1",
// 	}
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return Thread{}, err
// 	}
// 	for key, value := range headers {
// 		req.Header.Set(key, value)
// 	}
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return Thread{}, err
// 	}
// 	defer resp.Body.Close()
// 	var thread Thread
// 	if err := json.NewDecoder(resp.Body).Decode(&thread); err != nil {
// 		return Thread{}, err
// 	}
// 	return thread, nil
// }
