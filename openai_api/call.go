package openai_api

type MessageRequestBody struct {
	Role     string            `json:"role"`
	Content  string            `json:"content"`
	FileIDs  []string          `json:"file_ids,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

type MessageText struct {
	Value       string        `json:"value"`
	Annotations []interface{} `json:"annotations"`
}

type MessageContent struct {
	Type string      `json:"type"`
	Text MessageText `json:"text"`
}

type Message struct {
	ID          string           `json:"id"`
	Object      string           `json:"object"`
	CreatedAt   int64            `json:"created_at"`
	AssistantID string           `json:"assistant_id,omitempty"`
	ThreadID    string           `json:"thread_id"`
	RunID       string           `json:"run_id,omitempty"`
	Role        string           `json:"role"`
	Content     []MessageContent `json:"content"`
	FileIDs     []string         `json:"file_ids"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type MessageResponse struct {
	Object  string     `json:"object"`
	Data    []Message  `json:"data"`
	FirstID string     `json:"first_id"`
	LastID  string     `json:"last_id"`
	HasMore bool       `json:"has_more"`
}

// RunRequest represents the structure of the request body for creating a run
type RunRequest struct {
    AssistantID         string            `json:"assistant_id"`
    Model               string            `json:"model,omitempty"`
    Instructions        string            `json:"instructions,omitempty"`
    AdditionalInstructions string        `json:"additional_instructions,omitempty"`
    Tools               []Tool            `json:"tools,omitempty"`
    Metadata            map[string]string `json:"metadata,omitempty"`
}

// RunResponse represents the structure of the response body for creating a run
type RunResponse struct {
    ID           string `json:"id"`
    Object       string `json:"object"`
    CreatedAt    int64  `json:"created_at"`
    AssistantID  string `json:"assistant_id"`
    ThreadID     string `json:"thread_id"`
    Status       string `json:"status"`
    StartedAt    int64  `json:"started_at"`
    ExpiresAt    int64  `json:"expires_at,omitempty"`
    CancelledAt  int64  `json:"cancelled_at,omitempty"`
    FailedAt     int64  `json:"failed_at,omitempty"`
    CompletedAt  int64  `json:"completed_at,omitempty"`
    LastError    string `json:"last_error,omitempty"`
    Model        string `json:"model"`
    Instructions string `json:"instructions,omitempty"`
    Tools        []Tool `json:"tools,omitempty"`
    FileIDs      []string `json:"file_ids,omitempty"`
    Metadata     map[string]string `json:"metadata,omitempty"`
    Usage        interface{} `json:"usage,omitempty"`
}

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

