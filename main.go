package main

import (
	"ai_agents/agile_meth/config"
	"ai_agents/agile_meth/development"
	"ai_agents/agile_meth/model"
	"ai_agents/agile_meth/openaiapi"
	"ai_agents/agile_meth/planning"
	"log"
)

func main() {
	c := config.NewModelConfigs()["gpt4"]
	on := openaiapi.NewOpenAI(c.ApiKey, c.AssistUrl, c.ThreadUrl, c.RunUrl, c.Model)
	// Create instances of team members
    asst_id, err := on.CreateAssistant(development.DeveloperInst, "Developer", "code_interpreter")
    if err != nil{
        log.Fatalln(err)
    }
    thred_id, err := on.CreateThread()
    if err != nil{
        log.Fatalln(err)
    }
    Developer_John := development.NewDeveloper("John", asst_id, thred_id, c.ApiKey, c.ThreadUrl)
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
	Developer_John.WorkOn(user1)

	// tester.WorkOn(userStory)
	// designer.WorkOn(userStory)
}

func SimulateSprintPlanning(sprint *planning.Sprint, i int) {
	panic("unimplemented")
}
