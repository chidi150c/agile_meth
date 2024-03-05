package main

import (
	"ai_agents/agile_meth/config"
	"ai_agents/agile_meth/planning"
	// "ai_agents/agile_meth/ai_model/ollama"
	"ai_agents/agile_meth/ai_model/chat"
	"fmt"
)

func main() {
	cchat := config.NewModelConfigs()["gpt4"]
	devonc := chat.NewOpenAI("ChatDeveloper", planning.CodePrompt, cchat.ApiKey, cchat.Url, cchat.Model)
	scronc :=  chat.NewOpenAI("ChatScrum", planning.ScrumPrompt, cchat.ApiKey, cchat.Url, cchat.Model)
	queonc := chat.NewOpenAI("ChatQuestion", planning.QuestionPrompt, cchat.ApiKey, cchat.Url, cchat.Model)
	ansonc :=chat.NewOpenAI("ChatAnswer", planning.AnswersPrompt, cchat.ApiKey, cchat.Url, cchat.Model)
	preonc :=chat.NewOpenAI("ChatPresale", planning.PreSalesPrompt, cchat.ApiKey, cchat.Url, cchat.Model)
	maponc :=chat.NewOpenAI("ChatMapsale", planning.MappingPrompt, cchat.ApiKey, cchat.Url, cchat.Model)
	sumonc :=chat.NewOpenAI("ChatSumary", planning.Summarizer, cchat.ApiKey, cchat.Url, cchat.Model)

	// collama := config.NewModelConfigs()["ollama"]
    // devono := ollama.NewOllama("OllamaDeveloper", planning.CodePrompt, collama.Url, collama.Model)
    // scrono := ollama.NewOllama("OllamaScrum", planning.ScrumPrompt, collama.Url, collama.Model)
    // queono := ollama.NewOllama("OllamaQuestion", planning.QuestionPrompt, collama.Url, collama.Model)
    // ansono := ollama.NewOllama("OllamaAnswer", planning.AnswersPrompt, collama.Url, collama.Model)
    // preono := ollama.NewOllama("OllamaPresale", planning.PreSalesPrompt, collama.Url, collama.Model)
    // mapono := ollama.NewOllama("OllamaMap", planning.MappingPrompt, collama.Url, collama.Model)

	sm := planning.NewAgileProject("ScrumMaster", devonc, scronc, queonc, ansonc, preonc, maponc, sumonc)
	// vision := `designing and developing a worldwide game with immersive and engaging gameplay. The game is aimed at attracting millions of players globally through a mix of competitive and collaborative gameplay mechanics, strategy elements, customization options, and regular content updates. The vision includes creating a virtual environment using AI technology and fostering interaction, competition, and social interactions among players. The goals include implementing gameplay mechanics, strategy elements, content updates, social features, and feedback mechanisms. Additionally, user stories outline specific player desires such as competition, collaboration, customization, and content updates that align with the overall vision and goals for the game.`
	// vision := `To Build an App that maximizes cryptocurrency trading profit with automatic trading​`
	// vision := `Enhance your understanding and practical application of Go's concurrency features by developing a simple but efficient concurrent web crawler. This project will require you to utilize goroutines for parallel processing and channels for communication and synchronization among goroutines.`
    // vision := `Build a world wide game that millions around the world could play in the virtual environment generated by AI.`
	vision := `The vision is to create a robust and efficient VectorDB tailored specifically for the generative AI application, facilitating seamless data management and manipulation to support the AI's creative processes and enhance its overall performance.`
   	// vision := `Vision: Build a world wide game that millions around the world could play in the virtual environment generated by AI. Adopt Agile methodology and step by step code by code session by session implement the game using agile methodology.`
	vision =  sm.VisionArgumentation(vision) 
	goals := sm.BreakDownVisionIntoGoals(vision)
	
	sm.AllBuilder.WriteString(fmt.Sprintf("THE FOLLOWING ARE %d GOALS TO ACHIEVE THE VISION: %s\n\n", len(goals), sm.UpdatedVision))

	sm.MapBuilder.WriteString("Goals:\n")
    for k, v := range goals{
        sm.MapBuilder.WriteString(fmt.Sprintf("%d. %s\n", k+1, v.Goal))
    }
	sm.MapBuilder.WriteString("\nUser Stories:\n")
	userStories := sm.CreateUserStory(vision, 0)
    for k, v := range userStories{
        sm.MapBuilder.WriteString(fmt.Sprintf("%d. As a user, %s\n", k+1, v.Description))
    }	
	sm.MapBuilder.WriteString("\n")
	input := sm.MapBuilder.String()
	sm.AllBuilder.WriteString(fmt.Sprintln(input))
	sm.MapGoalAndUserStory(input)
	input = sm.AllBuilder.String()
	fmt.Println(input)
	input, _ = sm.SummaryEngine.ProcessAiMessage(input)
	
	fmt.Println(input)
	for k, v := range sm.ResearchBackLog{
		fmt.Printf("\nResearch %d: %s\n",k+1, v)
	}
}

