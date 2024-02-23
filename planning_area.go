package main

import (
	"fmt"
)

// UserStory represents a single user story in the backlog
type UserStory struct {
	ID          int
	Description string
	Priority    int
	EstimatedEffort int
}

// Backlog represents the collection of user stories for the project
type Backlog struct {
	UserStories []UserStory
}

// Function to add a new user story to the backlog
func (b *Backlog) AddUserStory(us UserStory) {
	b.UserStories = append(b.UserStories, us)
}

// Function to display all user stories in the backlog
func (b *Backlog) DisplayBacklog() {
	fmt.Println("Project Backlog:")
	for _, us := range b.UserStories {
		fmt.Printf("ID: %d, Description: %s, Priority: %d, Estimated Effort: %d\n", us.ID, us.Description, us.Priority, us.EstimatedEffort)
	}
}

func main() {
	// Create a new backlog
	backlog := Backlog{}

	// Add some user stories to the backlog
	backlog.AddUserStory(UserStory{ID: 1, Description: "Implement login functionality", Priority: 1, EstimatedEffort: 8})
	backlog.AddUserStory(UserStory{ID: 2, Description: "Design homepage layout", Priority: 2, EstimatedEffort: 5})
	backlog.AddUserStory(UserStory{ID: 3, Description: "Fix bug in user profile page", Priority: 3, EstimatedEffort: 3})

	// Display the backlog
	backlog.DisplayBacklog()
}
