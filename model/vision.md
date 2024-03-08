package model


// Task represents a task or subtask within a user story
type Task struct {
    ID          int
	UserStoryID int
    Description string
    AssignedTo  string // Name of the team member assigned to the task
    Completed   bool
    Comments    []string // Comments added by team members
}


type Vision struct {
	ID int
	UpdatedVision string
	Reasoning  string
	BaseVision string
	Goals      map[string]VisionGoal
	UserStories map[string]UserStory
	ConceptGoals    []VisionGoal
}