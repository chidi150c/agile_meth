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

	pj := planning.NewProject(openAI)

	// Define the initial vision
	vision := &model.Vision{Description: "To create a seamless online shopping experience for customers by integrating AI technology for personalized recommendations and efficient customer service."}

	// Iteratively refine goals and tasks based on new insights or changes
	var (
		goals []model.Goal
		err error
	)
	for i := 1; i <= 2; i++ {
		fmt.Printf("\nIteration %d:\n", i)
		fmt.Printf("\nVision: %s\n\n", vision.Description)

		// Derive goals from the current vision
		vision, goals, err = pj.DeriveGoalsFromVision(vision)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("The Following are Derived %d Goals for the Vision\n\n", len(goals))
		for j, goal := range goals {
			fmt.Printf("Goal %d: %s\n", j+1, goal.Description)
		}
		for j, goal := range goals {
			fmt.Printf("\nDerived Tasks for Goal %d: %s\n", j+1, goal.Description)
			// Derive tasks from each goal
			tasks := pj.DeriveTasksFromGoal(goal)

			// Display derived tasks
			for _, task := range tasks {
				fmt.Printf("%s\n", task.Description)
			}
			fmt.Println()
		}

		// Simulate updating the vision based on new insights or changes
		vision.Description = vision.UpdatedVision
		fmt.Printf("Update Vision:\n%s\n\n", vision.Description)
		pj.Vision = vision
	}
}


// 	vision := ""

// 	// Derive goals from the vision using the real AI model
// 	_, err := pj.DeriveGoalsFromVision(vision)
// 	if err != nil {
// 		log.Fatalln("DeriveGoalsFromVision Error:", err)
// 		return
// 	}
// 	_, err = pj.CreateUserStoriesForTheVision(pj.Vision.UpdatedVision)
// 	if err != nil {
// 		log.Fatalln("CreateUserStoriesForTheVision Error:", err)
// 		return
// 	}
// 	pj.MapUserStoriesToGoals()
// }
