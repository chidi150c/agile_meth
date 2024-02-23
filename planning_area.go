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
	AssignedTo      string
	Completed       bool // New field to track the completion status of the user story
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
		status := "Not Completed"
		if us.Completed {
			status = "Completed"
		}
		fmt.Printf("ID: %d, Description: %s, Priority: %d, Estimated Effort: %d, Assigned To: %s, Status: %s\n", us.ID, us.Description, us.Priority, us.EstimatedEffort, us.AssignedTo, status)
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

// Function to mark a user story as completed
func (b *Backlog) MarkUserStoryCompleted(id int) {
	for i, us := range b.UserStories {
		if us.ID == id {
			b.UserStories[i].Completed = true
			fmt.Printf("User story with ID %d has been marked as completed.\n", id)
			return
		}
	}
	fmt.Printf("User story with ID %d not found in the backlog.\n", id)
}

// Function to calculate the total estimated effort of all user stories in the backlog
func (b *Backlog) TotalEstimatedEffort() int {
	total := 0
	for _, us := range b.UserStories {
		total += us.EstimatedEffort
	}
	return total
}

// Function to calculate the total number of completed user stories in the backlog
func (b *Backlog) TotalCompletedUserStories() int {
	count := 0
	for _, us := range b.UserStories {
		if us.Completed {
			count++
		}
	}
	return count
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

	// Mark a user story as completed
	backlog.MarkUserStoryCompleted(1)

	// Display the updated backlog
	backlog.DisplayBacklog()

	// Calculate and display the total estimated effort of all user stories in the backlog
	totalEffort := backlog.TotalEstimatedEffort()
	fmt.Printf("Total estimated effort of all user stories in the backlog: %d\n", totalEffort)

	// Calculate and display the total number of completed user stories in the backlog
	completedUserStories := backlog.TotalCompletedUserStories()
	fmt.Printf("Total completed user stories in the backlog: %d\n", completedUserStories)
}
