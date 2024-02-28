package planning

import (
	"fmt"
	"time"
	"ai_agents/agile_meth/model"
    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"
    
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
// Function to calculate the remaining effort in the sprint
func (s *Sprint) RemainingEffort() int {
    return s.TotalEstimatedEffort() - s.TotalCompletedEffort()
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

// Function to simulate sprint planning meeting
func SimulateSprintPlanning(sprint *Sprint, teamCapacity int) {
    // Calculate total estimated effort of all user stories
    
    // Assign user stories to the sprint based on team capacity
    assignedEffort := 0
    for i, story := range sprint.UserStories {
        if assignedEffort+story.EstimatedEffort <= teamCapacity {
            sprint.UserStories[i].Assigned = true
            assignedEffort += story.EstimatedEffort
        } else {
            break
        }
    }
    fmt.Printf("Sprint Planning: Assigned %d user stories to the sprint.\n", len(sprint.UserStories))
}

// Function to generate data for burndown chart
func GenerateBurndownData(sprint *Sprint) ([]float64, []float64) {
    var days []float64
    var remainingEffort []float64
    
    // Calculate total estimated effort of all user stories
    totalEstimatedEffort := float64(sprint.TotalEstimatedEffort())
    remaining := totalEstimatedEffort
    
    // Iterate over each day of the sprint
    for day := 0; day <= int(sprint.EndDate.Sub(sprint.StartDate).Hours()/24); day++ {
        days = append(days, float64(day))
        remaining -= float64(sprint.Velocity()) // Assuming constant velocity for simplicity
        remainingEffort = append(remainingEffort, remaining)
    }
    
    return days, remainingEffort
}


// Function to plot the burndown chart
func PlotBurndownChart(days, remainingEffort []float64) error {
    // Create a new plot
    p := plot.New()

    // Create a new scatter plotter with markers for data points
    scatter, err := plotter.NewScatter(makeXYs(days, remainingEffort))
    if err != nil {
        return err
    }

    // Add the scatter plotter to the plot
    p.Add(scatter)

    // Set plot title and labels
    p.Title.Text = "Sprint Burndown Chart"
    p.X.Label.Text = "Day"
    p.Y.Label.Text = "Remaining Effort"

    // Save the plot to a PNG file
    if err := p.Save(10*vg.Inch, 6*vg.Inch, "burndown_chart.png"); err != nil {
        return err
    }

    return nil
}

// Helper function to convert days and remainingEffort slices to XYs slice
func makeXYs(days, remainingEffort []float64) plotter.XYs {
    xy := make(plotter.XYs, len(days))
    for i := range days {
        xy[i].X = days[i]
        xy[i].Y = remainingEffort[i]
    }
    return xy
}


