package main

import (
	"ai_agents/agile_meth/model"
	"ai_agents/agile_meth/planning"
	"fmt"
	"time"
)

func main() {
    // Create a new sprint
    startDate := time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC)
    endDate := time.Date(2024, time.March, 15, 0, 0, 0, 0, time.UTC)
    sprint := planning.NewSprint(1, startDate, endDate)
    
    // Add user stories to the sprint
    sprint.AddUserStory(model.UserStory{ID: 1, Description: "Implement login functionality", Priority: 1, EstimatedEffort: 8})
    sprint.AddUserStory(model.UserStory{ID: 2, Description: "Design homepage layout", Priority: 2, EstimatedEffort: 5})
    
    // Simulate sprint planning meeting
    planning.SimulateSprintPlanning(&sprint, 10) // Assuming team capacity is 10
    
    // Generate burndown chart data
    days, remainingEffort := planning.GenerateBurndownData(&sprint)
    
    // Plot burndown chart
    if err := planning.PlotBurndownChart(days, remainingEffort); err != nil {
        fmt.Println("Error plotting burndown chart:", err)
    } else {
        fmt.Println("Burndown chart saved as burndown_chart.png")
    }
}

func SimulateSprintPlanning(sprint *planning.Sprint, i int) {
	panic("unimplemented")
}
