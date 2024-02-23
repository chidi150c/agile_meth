package model

// Task represents a task or subtask within a user story
type Task struct {
    ID          int
    Description string
    AssignedTo  string // Name of the team member assigned to the task
    Completed   bool
    Comments    []string // Comments added by team members
}

// Function to assign a task to a team member
func (t *Task) AssignTo(memberName string) {
    t.AssignedTo = memberName
}

// Function to mark a task as completed
func (t *Task) MarkCompleted() {
    t.Completed = true
}

// Function to add a comment to the task
func (t *Task) AddComment(comment string) {
    t.Comments = append(t.Comments, comment)
}

