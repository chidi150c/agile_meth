package assistants

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type OpenAI struct {
	ApiKey      string
	AssistantID string
	ThreadID    string
	MessageID   string
	RunID       string
	AssistUrl   string
	ThreadUrl   string
	BaseUrl     string
	Model       string
}

func NewOpenAI(apiKey, assistUrl, threadUrl, model string) *OpenAI {
	o := &OpenAI{
		ApiKey:      apiKey,
		AssistUrl:   assistUrl,
		ThreadUrl:   threadUrl,
		Model:       model,
	}
	o.BaseUrl = o.ThreadUrl + "/" + o.ThreadID
	return o
}
func (o *OpenAI) CreateAssistant(inst, name, tYpe string) (string, error) {
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
	o.AssistantID = response.Id
	return response.Id, nil
}
func (o *OpenAI) CreateThread() (string, error) {
	// Define the request body
	requestBody := []byte("")
	// Create a new HTTP request
	req, err := http.NewRequest("POST", o.ThreadUrl, bytes.NewBuffer(requestBody))
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
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Unmarshal the response body into the ThreadResponse struct
	var threadResponse ThreadResponse
	if err := json.NewDecoder(resp.Body).Decode(&threadResponse); err != nil {
		return "", fmt.Errorf("in Thread Unexpected status code: %d", resp.StatusCode)
	}
	// Print the response data
	fmt.Println("Thread ID:", threadResponse.ID)
	o.ThreadID = threadResponse.ID
	o.BaseUrl = o.ThreadUrl + "/" + o.ThreadID
	return threadResponse.ID, nil
}

// CodeUserStory simulates the developer coding the given user story
func (o *OpenAI) ProcessAiMessage(msg string) (string, error) {
	// Initialize a map to store the parsed information or actions
	// Create the request body
	requestBody := MessageRequestBody{
		Role:    "user",
		Content: msg,
		// Add file IDs and metadata as needed
	}
	url := fmt.Sprintf("%s/messages", o.BaseUrl)
	//Creating user message
	data, err := o.APICall(url, "ProcessAiMessage", "POST", requestBody)
	if err != nil {
		log.Fatalln(err)
	} else if v, ok := data.(string); ok {
		o.MessageID = v
		fmt.Println("Message created successfully!")
	}
	o.Run()
	for {
		if status := o.GetRunStatus(); status == "completed" {
			fmt.Printf("Status = %s \n\n", status)
			break
		}
	}
	message, err := o.GetMessages()
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(message)
	// Initialize a map to store the parsed information or actions
	return message, nil
}
func (o *OpenAI) Run() {
	//Creating Run
	runReq := RunRequest{
		AssistantID: o.AssistantID,
		// You can add other fields here as needed
	}
	// fmt.Println("AssistID = ", o.AssistantID)
	//Creating Run continues
	data, err := o.APICall(o.BaseUrl+"/runs", "Run", "POST", runReq)
	if err != nil {
		log.Fatalln(err)
	} else if v, ok := data.(string); ok {
		o.RunID = v
	}
}
func (o *OpenAI) GetMessages() (string, error) {
	data, err := o.APICall(fmt.Sprintf("%s/messages/%s", o.BaseUrl, o.MessageID), "GetMessages", "GET", nil)
	if err != nil {
		log.Fatalln(err)
	} else if v, ok := data.(string); ok {
		return v, nil
	}
	return "", fmt.Errorf("GetMessage Error Something went wrong")
}
func (o *OpenAI) GetRunStatus() string {
	data, err := o.APICall(o.BaseUrl+"/runs/"+o.RunID, "GetRunStatus", "GET", nil)
	if err != nil {
		log.Fatalln(err)
	} else if v, ok := data.(string); ok {
		return v
	}
	return "GetMessage Error Something went wrong"
}
func (o *OpenAI) APICall(Url, name, method string, requestBody interface{}) (interface{}, error) {
	// Convert the request body to JSON
	var (
		req *http.Request
		err error
	)
	if method == "POST" {
		requestBodyJSON, err := json.Marshal(requestBody)
		if err != nil {
			return nil, fmt.Errorf("error at %s marshaling request body: %v", name, err)
		}
		// Send the POST request to create the message
		req, err = http.NewRequest("POST", Url, bytes.NewBuffer(requestBodyJSON))
		if err != nil {
			return nil, fmt.Errorf("error at %s creating request: %v", name, err)
		}
	} else {
		req, err = http.NewRequest("GET", Url, nil)
		if err != nil {
			return nil, fmt.Errorf("error at %s creating request: %v", name, err)
		}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
	req.Header.Set("OpenAI-Beta", "assistants=v1")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error at %s making request: %v ", name, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error at %s Unexpected status code: %v url: %s", name, resp.StatusCode, Url)
	}
	// Decode the response body into the MessageResponse struct
	if name == "CodeUserStory" {
		var messageResponse MessageResponse
		if err := json.NewDecoder(resp.Body).Decode(&messageResponse); err != nil {
			return nil, fmt.Errorf("error at %s decoding response body:%v", name, err)
		}
		return messageResponse.FirstID, nil
	} else if name == "Run" || name == "GetRunStatus" {
		// Unmarshal the response body into the RunResponse struct
		var runResponse RunResponse
		if err := json.NewDecoder(resp.Body).Decode(&runResponse); err != nil {
			return nil, fmt.Errorf("error at %s decoding response body:%v", name, err)
		}
		if name == "Run" {
			fmt.Println("Status:", runResponse.Status)
			fmt.Println("Status = in progress...")
			return runResponse.ID, nil
		} else if name == "GetRunStatus" {
			return runResponse.Status, nil
		}
	} else if name == "GetMessages" {
		// Decode the response body into the MessageResponse struct
		var messageResponse MessageResponse
		if err := json.NewDecoder(resp.Body).Decode(&messageResponse); err != nil {
			return "", fmt.Errorf("GetMessages Error decoding response body: %v", err)
		}
		// fmt.Println(messageResponse)
		return messageResponse.Data[0].Content[0].Text.Value, nil
	}
	return nil, nil
}
