package main

import (
	"ai_agents/agile_meth/config"
	"ai_agents/agile_meth/development"
	"ai_agents/agile_meth/model"
	"ai_agents/agile_meth/openai_api"
	"ai_agents/agile_meth/planning"
)

func main() {
	c := config.NewModelConfigs()["gpt4"]
	on := openai_api.NewOpenAI(c.ApiKey, c.AssistUrl, c.ThreadUrl, c.Model)
	// Create instances of team members
    Developer_John := development.NewDeveloper("John", on)
    user1 := &model.UserStory{
    	ID:            -  1,
    	Description:     "As a User, I want a sign-in mechanism, so that I can securely access my personal data and features exclusive to members.",
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
	Developer_John.CodeUserStory(user1)

	// tester.WorkOn(userStory)
	// designer.WorkOn(userStory)
}

func SimulateSprintPlanning(sprint *planning.Sprint, i int) {
	panic("unimplemented")
}
