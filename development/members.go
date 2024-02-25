package development

import (
	"ai_agents/agile_meth/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Define interface for team members
type TeamMember interface {
	WorkOn(userStory string)
}

// Developer struct representing a developer team member
type Developer struct {
	Name string
	Assistant_id string    
	Thread_id string  
	ApiKey string                             
}
func NewDeveloper (name string, assistID, threadID, apiKey string)*Developer{
	return &Developer{
		Name: name,
		Assistant_id: assistID,
		Thread_id: threadID,
		ApiKey: apiKey,
	}
}
// Implement the WorkOn method for Developer
func (d Developer) WorkOn(us *model.UserStory) {
	// Example thread ID
	threadID := d.Thread_id

	// Example message content
	messageContent := "user story: "+us.Description+" immediate task: "+us.Tasks[0].Description

	// Create the request body
	requestBody := MessageRequestBody{
		Role:    "user",
		Content: messageContent,
		// Add file IDs and metadata as needed
	}

	// Convert the request body to JSON
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling request body:", err)
		return
	}

	// Send the POST request to create the message
	url := fmt.Sprintf("https://api.openai.com/v1/threads/%s/messages", threadID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+d.ApiKey)
	req.Header.Set("OpenAI-Beta", "assistants=v1")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Message created successfully!")

}

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



