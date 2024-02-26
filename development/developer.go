package development

import (
	"ai_agents/agile_meth/model"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
	MessageID string
	BaseUrl string
	RunID string
	ApiKey      string
}

func NewDeveloper(name, assistID, threadID, apiKey, threadUrl string) *Developer {
	return &Developer{
		Name:        name,
		AssistantID: assistID,
		ThreadID:    threadID,
		ApiKey:      apiKey,
		BaseUrl: threadUrl+"/"+threadID,
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

	// Convert the request body to JSON
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling request body:", err)
		return
	}

	// Send the POST request to create the message
	url := fmt.Sprintf("%s/messages", d.BaseUrl)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+d.ApiKey)
	req.Header.Set("OpenAI-Beta", "assistants=v1")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Message Error: Unexpected status code:", resp.StatusCode)
		return
	}
	 // Decode the response body into the MessageResponse struct
	 var messageResponse MessageResponse
	 if err := json.NewDecoder(resp.Body).Decode(&messageResponse); err != nil {
		 fmt.Println("Error decoding response body:", err)
		 return
	 }
 
	d.MessageID = messageResponse.FirstID

	fmt.Println("Message created successfully!")
	d.Run()

	for{
		if status := d.GetRunStatus(); status == "completed"{
			break
		}
	}
	message, err := d.GetMessages()
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(message)
	return
}

func (d *Developer)Run(){
    // Define the request body
    requestBody := RunRequest{
        AssistantID: d.AssistantID,
        // You can add other fields here as needed
    }
    // Marshal the request body to JSON
    jsonBody, err := json.Marshal(requestBody)
    if err != nil {
        fmt.Println("Error marshalling the request body:", err)
        return
    }

    // Create a new HTTP request
    req, err := http.NewRequest("POST", d.BaseUrl+"/runs", bytes.NewBuffer(jsonBody))
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    // Set request headers
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
    req.Header.Set("OpenAI-Beta", "assistants=v1")

    // Send the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }
    defer resp.Body.Close()

    // Check response status code
    if resp.StatusCode != http.StatusOK {
        fmt.Println("Unexpected status code:", resp.StatusCode)
        return
    }

    // Unmarshal the response body into the RunResponse struct
    var runResponse RunResponse
    if err := json.NewDecoder(resp.Body).Decode(&runResponse); err != nil {
        fmt.Println("Error decoding response body:", err)
        return
    }

	// d.RunURL = fmt.Sprintf("%s/runs/%s",d.BaseUrl, runResponse.ID)
    // Print the response data
    fmt.Println("Status:", runResponse.Status)
	d.RunID = runResponse.ID
    // Print other relevant fields as needed
}

func (d *Developer)GetMessages() (string, error){
    // Create a new HTTP request
    req, err := http.NewRequest("GET", fmt.Sprintf("%s/messages/%s", d.BaseUrl, d.MessageID), nil)
    if err != nil {
        return "", fmt.Errorf("Error creating request: %v", err)
    }

    // Set request headers
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
    req.Header.Set("OpenAI-Beta", "assistants=v1")

    // Send the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("GetMessages Error sending request: %v", err)
    }
    defer resp.Body.Close()

    // Check response status code
    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("GetMessages Unexpected status code: %v", resp.StatusCode)
    }

    // Decode the response body into the MessageResponse struct
    var messageResponse MessageResponse
    if err := json.NewDecoder(resp.Body).Decode(&messageResponse); err != nil {
        return "", fmt.Errorf("GetMessages Error decoding response body: %v", err)
    }
	fmt.Println(messageResponse)
    return messageResponse.Data[0].Content[0].Text.Value, nil
}

// Example: Get run's status
func (d *Developer)GetRunStatus() string{
    // Create a new HTTP GET request
    req, err := http.NewRequest("GET", d.BaseUrl+"/runs/"+d.RunID, nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return ""
    }

    // Set request headers
    req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
    req.Header.Set("OpenAI-Beta", "assistants=v1")

    // Send the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return ""
    }
    defer resp.Body.Close()

     // Check response status code
	 if resp.StatusCode != http.StatusOK {
        fmt.Println("Unexpected status code:", resp.StatusCode)
        return ""
    }

    // Unmarshal the response body into the RunResponse struct
    var runResponse RunResponse
    if err := json.NewDecoder(resp.Body).Decode(&runResponse); err != nil {
        fmt.Println("Error decoding response body:", err)
        return ""
    }
    fmt.Println("Status:", runResponse.Status)
	return runResponse.Status
}

