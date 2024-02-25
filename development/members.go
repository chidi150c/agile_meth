package development

import (
	"ai_agents/agile_meth/model"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define interface for team members
type TeamMember interface {
	WorkOn(userStory string)
}

// Developer struct representing a developer team member
type Developer struct {
	Name        string
	AssistantID string
	ThreadID    string
	RunID       string
	ApiKey      string
	BaseUrl     string
}

func NewDeveloper(name, assistID, threadID, apiKey, threadUrl string) *Developer {
	return &Developer{
		Name:        name,
		AssistantID: assistID,
		ThreadID:    threadID,
		ApiKey:      apiKey,
		BaseUrl:     threadUrl + "/" + threadID,
	}
}

// Implement the WorkOn method for Developer
func (d *Developer) WorkOn(us *model.UserStory) {
	// Example message content
	messageContent := "user story: " + us.Description + " immediate task: " + us.Tasks[0].Description

	// Create the request body
	requestBody := MessageRequestBody{
		Role:    "user",
		Content: messageContent,
		// Add file IDs and metadata as needed
	}
	err := d.APICall("workOn", "POST", "/messages", requestBody)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Message created successfully!")
	d.Run()
	d.GetRunStatus()
}
func (d *Developer) APICall(name, method, url string, requestBody interface{}) error {
	// Convert the request body to JSON
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("%s Error marshaling request body: %v", name, err)
	}
	// Send the POST request to create the message
	Url := fmt.Sprintf("%s%s", d.BaseUrl, url)
	req, err := http.NewRequest(method, Url, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		return fmt.Errorf("%s Error creating request: %v", name, err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+d.ApiKey)
	req.Header.Set("OpenAI-Beta", "assistants=v1")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("%s Error making request: %v", name, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s %s Message Error: Unexpected status code: %v", name, Url, resp.StatusCode)
	}

	// Unmarshal the response body into the RunResponse struct
	var runResponse RunResponse

	if err := json.NewDecoder(resp.Body).Decode(&runResponse); err != nil {
		return fmt.Errorf("%s Error decoding response body: %v", name, err)
	}
	if name == "run" {
		fmt.Println("Status:", runResponse.Status)
		d.RunID = runResponse.ID
	} else if name == "runStatus" {
		fmt.Println("Status:", runResponse.Status)
	}
	return nil
}

// RunRequest represents the structure of the request body for creating a run
type RunRequest struct {
	AssistantID            string            `json:"assistant_id"`
	ThreadID               string            `json:"thread_id"`
	Model                  string            `json:"model,omitempty"`
	Instructions           string            `json:"instructions,omitempty"`
	AdditionalInstructions string            `json:"additional_instructions,omitempty"`
	Tools                  []Tool            `json:"tools,omitempty"`
	Metadata               map[string]string `json:"metadata,omitempty"`
}

// Tool represents a tool for the run
type Tool struct {
	Type string `json:"type"`
}

// RunResponse represents the structure of the response body for creating a run
type RunResponse struct {
	ID           string            `json:"id"`
	Object       string            `json:"object"`
	CreatedAt    int64             `json:"created_at"`
	AssistantID  string            `json:"assistant_id"`
	ThreadID     string            `json:"thread_id"`
	Status       string            `json:"status"`
	StartedAt    int64             `json:"started_at"`
	ExpiresAt    int64             `json:"expires_at,omitempty"`
	CancelledAt  int64             `json:"cancelled_at,omitempty"`
	FailedAt     int64             `json:"failed_at,omitempty"`
	CompletedAt  int64             `json:"completed_at,omitempty"`
	LastError    string            `json:"last_error,omitempty"`
	Model        string            `json:"model"`
	Instructions string            `json:"instructions,omitempty"`
	Tools        []Tool            `json:"tools,omitempty"`
	FileIDs      []string          `json:"file_ids,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
	Usage        interface{}       `json:"usage,omitempty"`
}

func (d *Developer) Run() {
	// Define the request body
	requestBody := RunRequest{
		AssistantID: d.AssistantID,
		// You can add other fields here as needed
	}
	err := d.APICall("run", "POST", "/runs", requestBody)
	if err != nil {
		log.Fatalln(err)
	}
	// Print other relevant fields as needed
}

// Example: Get run's status
func (d *Developer) GetRunStatus() {
	err := d.APICall("runStatus", "GET", "/runs/"+d.RunID, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// Implement similar functions for other API requests

// Tester struct representing a tester team member
type Tester struct {
	Name string
}

// Implement the WorkOn method for Tester
func (t Tester) WorkOn(userStory string) {
	fmt.Printf("Tester %s is testing user story: %s\n", t.Name, userStory)
}

// Designer struct representing a designer team member
type Designer struct {
	Name string
}

// Implement the WorkOn method for Designer
func (d Designer) WorkOn(userStory string) {
	fmt.Printf("Designer %s is designing user story: %s\n", d.Name, userStory)
}

type MessageRequestBody struct {
	Role     string            `json:"role"`
	Content  string            `json:"content"`
	FileIDs  []string          `json:"file_ids,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}
