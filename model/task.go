package model

// Task represents a task or subtask within a user story
type Task struct {
    ID          int
    Description string
    AssignedTo  string // Name of the team member assigned to the task
    Completed   bool
}

// Function to assign a task to a team member
func (t *Task) AssignTo(memberName string) {
    t.AssignedTo = memberName
}

// Function to mark a task as completed
func (t *Task) MarkCompleted() {
    t.Completed = true
}

