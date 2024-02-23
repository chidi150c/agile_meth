package planning

import (
	"fmt"
	"time"
	"ai_agents/agile_meth/model"
)

// Sprint represents a sprint in the Agile project
type Sprint struct {
    ID           int
    StartDate    time.Time
    EndDate      time.Time
    UserStories  []model.UserStory // User stories included in the sprint
}

// Function to create a new sprint with the specified start and end dates
func NewSprint(id int, startDate, endDate time.Time) Sprint {
    return Sprint{
        ID:        id,
        StartDate: startDate,
        EndDate:   endDate,
    }
}

// Function to add a user story to the sprint
func (s *Sprint) AddUserStory(us model.UserStory) {
    s.UserStories = append(s.UserStories, us)
}
// Function to calculate the total estimated effort of all user stories in the sprint
func (s *Sprint) TotalEstimatedEffort() int {
    total := 0
    for _, story := range s.UserStories {
        total += story.EstimatedEffort
    }
    return total
}
// Function to calculate the total completed effort of all user stories in the sprint
func (s *Sprint) TotalCompletedEffort() int {
    total := 0
    for _, story := range s.UserStories {
        if story.Completed {
            total += story.EstimatedEffort
        }
    }
    return total
}
// Function to calculate the duration of the sprint in days
func (s *Sprint) Duration() int {
    duration := s.EndDate.Sub(s.StartDate)
    return int(duration.Hours() / 24) // Convert duration from hours to days
}
// Function to calculate the sprint velocity (completed effort per day) in the sprint
func (s *Sprint) Velocity() float64 {
    durationDays := float64(s.Duration())
    if durationDays == 0 {
        return 0 // Avoid division by zero
    }
    completedEffort := float64(s.TotalCompletedEffort())
    return completedEffort / durationDays
}
// Function to mark a user story as completed within the sprint
func (s *Sprint) MarkUserStoryCompleted(storyID int) {
    for i, story := range s.UserStories {
        if story.ID == storyID {
            s.UserStories[i].Completed = true
            fmt.Printf("User story with ID %d marked as completed in Sprint %d\n", storyID, s.ID)
            return
        }
    }
    fmt.Printf("User story with ID %d not found in Sprint %d\n", storyID, s.ID)
}
// Function to calculate the percentage of completed user stories in the sprint
func (s *Sprint) SprintProgress() float64 {
    completedCount := 0
    for _, story := range s.UserStories {
        if story.Completed {
            completedCount++
        }
    }
    return float64(completedCount) / float64(len(s.UserStories)) * 100
}
// Function to calculate the sprint velocity (total estimated effort of completed user stories)
func (s *Sprint) SprintVelocity() int {
    totalEffort := 0
    for _, story := range s.UserStories {
        if story.Completed {
            totalEffort += story.EstimatedEffort
        }
    }
    return totalEffort
}