package model


// Vision represents the overarching objective of the project.
type Vision struct {
	Description string
	UpdatedVision string
	DarftedGoals []Goal
	DraftedUserStories []UserStory
	Goals map[string]*Goal
	UserStories map[string]*UserStory
}

// Goal represents a specific objective derived from the vision.
type Goal struct {
	ID        int
	Reasoning string
	Concept   string
	Description string
	MappedUserStories int
	Question     string
	Answer       string
	Completed       bool   // New field to track the completion status of the user story
}
// UserStory represents a single user story in the backlog
type UserStory struct {
	ID              int
	Concept   string
	Reasoning string
	Description     string
	Priority        string
	EstimatedEffort int
	Assigned        bool // Indicates whether the user story is assigned to the sprint
	AssignedTo      string
	Completed       bool   // New field to track the completion status of the user story
	Tasks           []*Task // List of tasks/subtasks for the user story
	MappedGoals 	int
}
// Task represents an individual action or activity required to achieve a goal.
type Task struct {
	ID int
	Description string
}

