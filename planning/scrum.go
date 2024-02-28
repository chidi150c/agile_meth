package planning

import (
	"ai_agents/agile_meth/ai_model"
	"ai_agents/agile_meth/model"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

// AgileProject represents an Agile project with its backlog of user stories.
type AgileProject struct {
	Vision string
	VisionClarityQuestions []model.VisionClarityQuestion
	Backlog *Backlog
	Engine  ai_model.AIServicer
}

func NewAgileProject(name string, engine ai_model.AIServicer) *AgileProject {
	// _, err := engine.CreateAssistant(Prompt2, name, "")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// _, err = engine.CreateThread()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	
	return &AgileProject{
		Engine: engine,
		Backlog: NewBacklog(),
	}
}

func fillClarityQuestions(input string) ([]model.VisionClarityQuestion, error) {

	fmt.Printf("input: %s", input)
	// Regular expressions to extract questions and concepts.
	questionRegex := regexp.MustCompile(`Question \d+: (.+)`)
	conceptRegex := regexp.MustCompile(`Concept \d+: (.+)`)

	// Split the input string by newlines to process each line separately.
	lines := strings.Split(input, "\n")

	// Slice to store the extracted questions and concepts.
	var visionQuestions []model.VisionClarityQuestion

	// Iterate over each line in the input string.
	for _, line := range lines {
		// Try to find a question and concept in the line.
		questionMatch := questionRegex.FindStringSubmatch(line)
		conceptMatch := conceptRegex.FindStringSubmatch(line)

		// If both a question and a concept are found in the line, add them to the visionQuestions slice.
		if len(questionMatch) > 1 && len(conceptMatch) > 1 {
			visionQuestions = append(visionQuestions, model.VisionClarityQuestion{
				Question: questionMatch[1],
				Concept:  conceptMatch[1],
			})
		}
	}

	// Check if any questions and concepts were extracted.
	if len(visionQuestions) == 0 {
		return nil, errors.New("no vision clarity questions and concepts found")
	}

	// Print the extracted questions and concepts.
	for _, q := range visionQuestions {
		fmt.Printf("Question: %s\n", q.Question)
		fmt.Printf("Concept: %s\n\n", q.Concept)
	}

	return visionQuestions, nil
}


func fillGoalandReason(v string)*model.VisionGoal{
	s := strings.Split(v, "Reasoning:")[1]
	p := strings.Split(s, "Goal:")
	return &model.VisionGoal{Reason: p[0], Goal: p[1]}
}
// Function to clarify project vision 
func (project *AgileProject) VisionClarification(vision string) []model.VisionClarityQuestion {
	project.Vision = vision
	goals, err := project.Engine.ProcessAiMessage(vision)
	if err != nil{
		log.Fatalln(err)
	}
	project.VisionClarityQuestions, err = fillClarityQuestions(goals)
	if err != nil{
		log.Fatalln(err)
	}
	return project.VisionClarityQuestions
}

// Function to break down project vision into 5 goals
func (project *AgileProject) BreakDownVision(vision string) []*model.VisionGoal {
	project.Vision = vision
	goals, err := project.Engine.ProcessAiMessage(vision)
	if err != nil{
		log.Fatalln(err)
	}
	GoalsAndReasonings := strings.Split(goals, "\n\n")
	for _, v :=  range GoalsAndReasonings{
		project.Backlog.VisionGoals = append (project.Backlog.VisionGoals, fillGoalandReason(v))
	}
	return project.Backlog.VisionGoals
}

// CreateUserStory creates a new user story with the given description and priority.
func (project *AgileProject) CreateUserStory(description string, priority int) *model.UserStory {
	return &model.UserStory{Description: description, Priority: priority}
}

// PrioritizeBacklog prioritizes the backlog of user stories based on their priority.
func (project *AgileProject) PrioritizeBacklog() {
	// Sorting logic here...
	fmt.Println("Prioritizing backlog...")
	time.Sleep(1 * time.Second) // Simulating sorting process
	fmt.Println("Backlog prioritized successfully!")
}

// RefineGoalsAndBacklog refines the project goals and backlog based on new insights.
func (project *AgileProject) RefineGoalsAndBacklog(newGoals []string) {
	fmt.Println("Refining project goals and backlog based on new insights...")
	// Logic to incorporate new goals into project goals and backlog
	for _, goal := range newGoals {
		fmt.Println("Goal = ", goal)
		// Create user story for each new goal and add it to the backlog
		// project.Backlog.UserStories = append(project.Backlog.UserStories, project.CreateUserStory(goal, len(project.Backlog)+1))
	}
	fmt.Println("Project goals and backlog refined successfully!")
}

// // PrintBacklog prints the backlog of user stories.
// func (project AgileProject) PrintBacklog() {
// 	fmt.Println("Product Backlog:")
// 	for _, story := range project.Backlog {
// 		fmt.Printf("- %s (Priority: %d)\n", story.Description, story.Priority)
// 	}
// }





// func main() {
// 	// Create Agile project
// 	project := AgileProject{}

// 	// Initial project goals
// 	initialGoals := []string{
// 		"Develop a Scalable Game Infrastructure",
// 		"Implement AI-driven Virtual Environment",
// 		"Ensure Global Accessibility",
// 		"Promote User Engagement and Retention",
// 		"Iterate Based on User Feedback",
// 	}

	// Populate project backlog with initial goals
// 	project.Backlog = append(project.Backlog, createUserStories(initialGoals)...)

// 	// Prioritize backlog
// 	project.PrioritizeBacklog()

// 	// Print initial backlog
// 	project.PrintBacklog()

// 	// Simulate gathering new insights and refining goals
// 	newInsights := []string{
// 		"Incorporate Virtual Reality (VR) technology",
// 		"Enhance community-building features",
// 	}

// 	// Refine project goals and backlog based on new insights
// 	project.RefineGoalsAndBacklog(newInsights)

// 	// Prioritize backlog after refinement
// 	project.PrioritizeBacklog()

// 	// Print refined backlog
// 	project.PrintBacklog()
// }

// // createUserStories creates user stories from a slice of goal descriptions.
// func createUserStories(goals []string) []UserStory {
// 	var userStories []UserStory
// 	for _, goal := range goals {
// 		userStories = append(userStories, CreateUserStory(goal, len(userStories)+1))
// 	}
// 	return userStories
// }
