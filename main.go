package main

import (
	"ai_agents/agile_meth/model"
	"ai_agents/agile_meth/planning"
	"fmt"
	"time"
)

func main(){
	// Create a new backlog
	backlog := planning.NewBacklog()

	userStory1 := model.UserStory{ID: 1, Description: "Implement login functionality", Priority: 1, EstimatedEffort: 8}
	userStory2 := model.UserStory{ID: 2, Description: "Design homepage layout", Priority: 2, EstimatedEffort: 5}
	userStory3 := model.UserStory{ID: 3, Description: "Fix bug in user profile page", Priority: 3, EstimatedEffort: 3}

	// Add some user stories to the backlog
	backlog.AddUserStory(userStory1)
	backlog.AddUserStory(userStory2)
	backlog.AddUserStory(userStory3)

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


	// Example usage
	startDate := time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, time.March, 15, 0, 0, 0, 0, time.UTC)
	sprint := planning.NewSprint(1, startDate, endDate)

	// Add some user stories to the sprint
	sprint.AddUserStory(userStory1)
	sprint.AddUserStory(userStory2)

	// Mark a user story as completed within the sprint
	sprint.MarkUserStoryCompleted(1)

	// Calculate and display sprint velocity, total estimated effort and completed effort within the sprint
	velocity := sprint.Velocity()
	totalEstimatedEffort := sprint.TotalEstimatedEffort()
	totalCompletedEffort := sprint.TotalCompletedEffort()
	fmt.Printf("Total estimated effort within Sprint %d: %d\n", sprint.ID, totalEstimatedEffort)
	fmt.Printf("Total completed effort within Sprint %d: %d\n", sprint.ID, totalCompletedEffort)
	fmt.Printf("Sprint Velocity for Sprint %d: %.2f effort per day\n", sprint.ID, velocity)

	// Add tasks to the user story
	userStory1.AddTask(model.Task{ID: 1, Description: "Design login form", AssignedTo: "John Doe"})
	userStory1.AddTask(model.Task{ID: 2, Description: "Backend implementation", AssignedTo: "Jane Smith"})
	
	// Function to display all tasks within the user story
    fmt.Printf("Tasks for User Story %d: \n", userStory1.ID)
    for _, task := range userStory1.Tasks {
        status := "Not Completed"
        if task.Completed {
            status = "Completed"
        }
        fmt.Printf("ID: %d, Description: %s, Assigned To: %s, Status: %s\n", task.ID, task.Description, task.AssignedTo, status)
    }

}