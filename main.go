package main

import (
	"ai_agents/agile_meth/ai_model/chat"
	"ai_agents/agile_meth/config"
	"ai_agents/agile_meth/model"
	"ai_agents/agile_meth/planning"
	"fmt"
	"log"
)





func main() {
	// Define AI configuration
	config := config.NewModelConfigs()["gpt4"]

	// Create an instance of the real AI model
	openAI := chat.NewOpenAI(config)

	vision := model.NewVision()
	// Define the initial vision
	vision.Description = "To revolutionize the fitness industry by empowering users to achieve their fitness goals through personalized workout plans on a mobile application."
	
	pj := planning.NewProject(openAI, vision)

	// Iteratively refine goals and tasks based on new insights or changes
	var (
		goals []model.Goal
		// err error
	)
	
	fmt.Printf("\nBase Vision: %s\n", vision.Description)
	fmt.Printf("\nActiculated Vision: %s\n\n", pj.ActiculatedVision(vision))
	
	for i := 1; i <= 5; i++ {
		fmt.Printf("\nIteration %d:\n\n", i)
		if i == 1{
			goals, vision = pj.BreakDownVisionIntoGoals(vision)
			fmt.Printf("The Following are Derived %d Goals for the Vision\n\n", len(goals))
			for _, goal := range goals {
				fmt.Printf("Goal %d: %s\n", goal.ID, goal.Description)
			}
		}else{
			goals, vision = pj.BreakDownVisionIntoNextGoals(vision)
			fmt.Printf("\n\nThe Following are Derived %d Goals for the Vision\n\n", len(goals))
			for _, goal := range goals {
				fmt.Printf("Goal %d: %s\n", goal.ID, goal.Description)
			}			 
		}
		// for j, goal := range goals {
		// 	fmt.Printf("\nDerived Tasks for Goal %d: %s\n", j+1, goal.Description)
		// 	// Derive tasks from each goal
		// 	tasks := pj.DeriveTasksFromGoal(goal)

		// 	// Display derived tasks
		// 	for _, task := range tasks {
		// 		fmt.Printf("%s\n", task.Description)
		// 	}
		// 	fmt.Println()
		// }

		
		//Create User Story from updated Vision
		// if i == 1{
			_, err := pj.CreateUserStoriesForTheVision(vision)
			if err != nil {
				log.Fatalln(err)
			}
		// }
		fmt.Printf("\nThe Following are Derived %d User Sttories for the Vision\n\n", len(vision.UserStories))
		for _, userStory := range vision.DraftedUserStories{
			if userStory.MappedGoals == 0 {
				fmt.Printf("%d: As a user, %s\n", userStory.ID, userStory.Description)
			}
		}

		pj.MapUserStoriesToGoals(vision)
		// input := pj.AllBuilder.String()
		
		// fmt.Printf("Description: %v\n\n", vision.Description)
		// fmt.Printf("UpdatedVision: %v\n\n", vision.UpdatedVision)
		// fmt.Printf("DraftedGoals: %v\n\n", vision.DraftedGoals)
		// fmt.Printf("DraftedUserStories: %v\n\n", vision.DraftedUserStories)
		// fmt.Printf("Goals: %v\n\n", vision.Goals)
		// fmt.Printf("UserStories: %v\n\n", vision.UserStories)

		fmt.Printf("\nLenth of Backlog %d and lenght of goals %d", len(pj.UmappedBackLog), len(vision.Goals))
		features := ""
		if len(pj.UmappedBackLog) != 0{
			for _,v := range pj.UmappedBackLog{
				features = features + "; "+v
			}
			vision.Description, _ = pj.AI.PromptAI("Summarize in a brief paragraph of an acticulated vision satement", fmt.Sprintf("%s\nFeature: %s",vision.UpdatedVision, features))
			pj.UmappedBackLog = []string{}
		}
		fmt.Printf("\nUpdated Vision:\n%s\n\n", vision.Description)
	}
}

