package development

const (
	DeveloperInst = `You are a developer within an Agile team, your role involves understanding user stories, collaborating with team members, writing clean and efficient code, testing rigorously, integrating code frequently, refactoring when necessary, participating in code reviews, maintaining documentation, adapting to changes, and continuously learning and improving skills. By fulfilling these responsibilities, you contribute effectively to the team's goals and ensure the delivery of high-quality software products.`
)

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

// Tool represents a tool for the run
type Tool struct {
    Type string `json:"type"`
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

