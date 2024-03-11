package main

import (
	"ai_agents/agile_meth/ai_model/chat"
	"ai_agents/agile_meth/config"
	"ai_agents/agile_meth/model"
	"ai_agents/agile_meth/planning"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)





func main() {
	
	reader := bufio.NewReader(os.Stdin)
	
	// Define AI configuration
	config := config.NewModelConfigs()["gpt4"]

	// Create an instance of the real AI model
	openAI := chat.NewOpenAI(config)

	vision := model.NewVision()

	// vision.Description = "To revolutionize the fitness industry by empowering users to achieve their fitness goals through personalized workout plans on a mobile application."
	vision.Description = "To build a user-friendly platform that seamlessly connects pet owners with trustworthy pet sitters, enhancing the overall pet care experience."
	
	pj := planning.NewProject(openAI, vision, reader)

	// Define the initial vision
	fmt.Print("\nEnter your vision statement and press ENTER: \n")
	inputCode, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputCode = strings.TrimSpace(inputCode)
	if inputCode != ""{
		vision.Description = inputCode
	}

	// Iteratively refine goals and tasks based on new insights or changes
	var (
		goals []model.Goal
		// err error
	)
	//uncheck
	_ = goals
	// fmt.Printf("\nBase Vision: %s\n", vision.Description)
	// fmt.Printf("\nActiculated Vision: %s\n\n", vacticulated)
	pj.Fnum = 1
	pj.FilePath = ""
	for{	
		pj.VisionEnhancement(vision)
		for{
			fmt.Print("\nEnter manual or auto: ")
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(choice)
			if choice != "manual" && choice != "auto"{
				fmt.Println("\n Invalid Input!!!!")
				continue
			}else if choice == "manual"{	
				pj.ManualRun(vision, reader)	
				break	
			}else if choice == "auto"{	
				pj.AutomaticRun(vision)		
			}
		}
	}
}

