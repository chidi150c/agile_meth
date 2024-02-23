package planning

import (
	"fmt"
	"time"
)
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


func main() {
    // Example usage
    startDate := time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC)
    endDate := time.Date(2024, time.March, 15, 0, 0, 0, 0, time.UTC)
    sprint := NewSprint(1, startDate, endDate)
    
    // Add some user stories to the sprint
    sprint.AddUserStory(UserStory{ID: 1, Description: "Implement login functionality", Priority: 1, EstimatedEffort: 8})
    sprint.AddUserStory(UserStory{ID: 2, Description: "Design homepage layout", Priority: 2, EstimatedEffort: 5})
    
    // Mark a user story as completed within the sprint
    sprint.MarkUserStoryCompleted(1)
    
    // Calculate and display sprint velocity
    velocity := sprint.Velocity()
    fmt.Printf("Sprint Velocity for Sprint %d: %.2f effort per day\n", sprint.ID, velocity)
}
