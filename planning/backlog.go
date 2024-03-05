package planning

import (
	"ai_agents/agile_meth/model"
	"fmt"
)

// Backlog represents the collection of user stories for the project
type Backlog struct {
	UserStories []*model.UserStory
	VisionGoals []*model.VisionGoal
	Mapping []model.Mapping
}


func NewBacklog ()*Backlog{
	return &Backlog{
		UserStories: make([]*model.UserStory,0),
    	VisionGoals: make([]*model.VisionGoal,0),
	}
}
// Function to add a new user story to the backlog
func (b *Backlog) AddUserStory(us *model.UserStory) {
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


