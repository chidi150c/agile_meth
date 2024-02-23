package model

// UserStory represents a single user story in the backlog
type UserStory struct {
	ID              int
	Description     string
	Priority        int
	EstimatedEffort int
	AssignedTo      string
	Completed       bool // New field to track the completion status of the user story
    Tasks           []Task // List of tasks/subtasks for the user story
}

// Task represents a task or subtask within a user story
type Task struct {
    ID          int
    Description string
    AssignedTo  string
    Completed   bool
}

// Function to add a new task to the user story
func (us *UserStory) AddTask(task Task) {
    us.Tasks = append(us.Tasks, task)
}


