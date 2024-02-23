package main

import (
	"ai_agents/agile_meth/model"
	"ai_agents/agile_meth/planning"
	"fmt"
)

func main(){
	// Create a new backlog
	backlog := planning.Backlog{}
	
	// Add some user stories to the backlog
	backlog.AddUserStory(model.UserStory{ID: 1, Description: "Implement login functionality", Priority: 1, EstimatedEffort: 8})
	backlog.AddUserStory(model.UserStory{ID: 2, Description: "Design homepage layout", Priority: 2, EstimatedEffort: 5})
	backlog.AddUserStory(model.UserStory{ID: 3, Description: "Fix bug in user profile page", Priority: 3, EstimatedEffort: 3})
	
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