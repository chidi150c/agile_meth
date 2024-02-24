package main

import (
	"ai_agents/agile_meth/config"
	"ai_agents/agile_meth/development"
	"ai_agents/agile_meth/model"
	"ai_agents/agile_meth/openaiapi"
	"ai_agents/agile_meth/planning"
)

func main() {
	c := config.NewModelConfigs()["gpt3"]
	on := openaiapi.NewOpenAI(c.ApiKey, c.Url, c.Model)
	// Create instances of team members
	dev := development.Developer{
		Name:  "John",
		Aiapi: on,
	}
	// tester := Tester{Name: "Alice"}
	// designer := Designer{Name: "Bob"}
    user1 := &model.UserStory{
    	ID:              1,
    	Description:     "Implement login functionality",
    	Priority:        1,
    	EstimatedEffort: 8,
    	Assigned:        false,
    	AssignedTo:      "",
    	Completed:       false,
    	Tasks:           []model.Task{
                            {Description: "create a signin form",},
                        },
    }
	// Simulate work on user stories
	dev.WorkOn(user1)

	// tester.WorkOn(userStory)
	// designer.WorkOn(userStory)
}

func SimulateSprintPlanning(sprint *planning.Sprint, i int) {
	panic("unimplemented")
}
