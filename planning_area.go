package main

import (
	"fmt"
)

// UserStory represents a single user story in the backlog
type UserStory struct {
	ID              int
	Description     string
	Priority        int
	EstimatedEffort int
	AssignedTo      string // New field to track the team member assigned to the user story
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
		fmt.Printf("ID: %d, Description: %s, Priority: %d, Estimated Effort: %d, Assigned To: %s\n", us.ID, us.Description, us.Priority, us.EstimatedEffort, us.AssignedTo)
	}
}

// Function to remove a user story from the backlog by ID
func (b *Backlog) RemoveUserStoryByID(id int) {
	for i, us := range b.UserStories {
		if us.ID == id {
			b.UserStories = append(b.UserStories[:i], b.UserStories[i+1:]...)
			fmt.Printf("User story with ID %d has been removed from the backlog.\n", id)
			return
		}
	}
	fmt.Printf("User story with ID %d not found in the backlog.\n", id)
}

// Function to assign a user story to a team member
func (b *Backlog) AssignUserStory(id int, assignee string) {
	for i, us := range b.UserStories {
		if us.ID == id {
			b.UserStories[i].AssignedTo = assignee
			fmt.Printf("User story with ID %d has been assigned to %s.\n", id, assignee)
			return
		}
	}
	fmt.Printf("User story with ID %d not found in the backlog.\n", id)
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

	// Remove a user story from the backlog
	backlog.RemoveUserStoryByID(2)

	// Assign a user story to a team member
	backlog.AssignUserStory(1, "John Doe")

	// Display the updated backlog
	backlog.DisplayBacklog()
}
