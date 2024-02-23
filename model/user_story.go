package model

// UserStory represents a single user story in the backlog
type UserStory struct {
	ID              int
	Description     string
	Priority        int
	EstimatedEffort int
	AssignedTo      string
	Completed       bool // New field to track the completion status of the user story
}