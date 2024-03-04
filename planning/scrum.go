package planning

import (
	"ai_agents/agile_meth/ai_model"
	"ai_agents/agile_meth/model"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

// AgileProject represents an Agile project with its backlog of user stories.
type AgileProject struct {
	Vision                 string
	ConceptGoals []model.ConceptGoal
	Backlog                *Backlog
	DeveloperEngine                 ai_model.AIServicer
	QuestionsEngine                 ai_model.AIServicer
	AnswersEngine                 ai_model.AIServicer
	ScrumEngine                 ai_model.AIServicer
	PreSalesEngine   ai_model.AIServicer
	ResearchBackLog []string
}

func NewAgileProject(name string, deveng, srceng, queeng, anseng, preeng ai_model.AIServicer) *AgileProject {

	return &AgileProject{
		Backlog: NewBacklog(),
		DeveloperEngine: deveng,
		QuestionsEngine: queeng,
		AnswersEngine: anseng,
		ScrumEngine: srceng ,
		PreSalesEngine: preeng,
	}
}
func (project *AgileProject)VisionArgumentation(vision string)string{
	questions := project.QuestionsForVisionClarification(vision)
	vi := ""
	i := 1
	fmt.Printf("\n\nBELOW ARE QUESTIONS AND ANSWERS TO CLARIFY THE FEATURES AND CHALLENGES EXPECTED FOR THE VISION")
    for k, v := range questions{
        fmt.Printf("\n\nQuestion %d: %s\nGoal %d: %s \n", k+1, v.Question, k+1, v.Goal)
		v.Answer = project.AnswerQuestionForVisionClarity(v.Question)
		fmt.Printf("Answer %d: %s", k+1, v.Answer)
		vi = fmt.Sprintf("Answer: %s", v.Answer)
		if strings.Contains(vi, "Unknown:"){
			//send vi to R&D
			project.ResearchBackLog = append(project.ResearchBackLog, fmt.Sprintf("%s\n%s", v.Question, vi))
		}else{
			vi = strings.Replace(vi, "Answer:", "Feature "+fmt.Sprint(i)+":", -1)
			if i == 1{
				vision = fmt.Sprintf("%s\n\n%s", vision, vi)
			}else{
				vision = fmt.Sprintf("%s\n%s", vision, vi)
			}
			i++
		}
    }

	fmt.Printf("\n\nTHE FOLLOWING IS THE UPDATED VISION WITH SOME FEATURES:\n")
    fmt.Printf("\n%s\n\n", vision)
	return vision
}
// Function to clarify project vision
func (project *AgileProject) QuestionsForVisionClarification(vision string) []model.ConceptGoal {
	project.Vision = vision
	goals, err := project.QuestionsEngine.ProcessAiMessage(vision)
	if err != nil {
		log.Fatalln(err)
	}
	project.ConceptGoals = fillClarityQuestions(goals)
	return project.ConceptGoals
}
// AnswerQuestionForVisionClarity function generates an answer for the question using LLM.
func (project *AgileProject)AnswerQuestionForVisionClarity(question string) string {
	s := fmt.Sprintf("Vision : %s\nQuestion : %s", project.Vision, question)
	qaAnswer, err := project.AnswersEngine.ProcessAiMessage(s)
	if err != nil{
		log.Fatalln(err)
	}
	return strings.Split(qaAnswer,"Answer: ")[1]
}
// Function to break down project vision into 5 goals
func (project *AgileProject) BreakDownVisionIntoGoals(vision string) []*model.VisionGoal {
	project.Vision = vision
	goals, err := project.ScrumEngine.ProcessAiMessage(vision)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Printf("virtuallizing Output before split : %v\n\n", goals)
	project.Backlog.VisionGoals = fillGoalandReason(goals)
	return project.Backlog.VisionGoals
}
// CreateUserStory creates a new user story with the given description and priority.
func (project *AgileProject) CreateUserStory(description string, priority int) *model.UserStory {
	description, _ = project.PreSalesEngine.ProcessAiMessage(description)
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

func fillClarityQuestions(input string) []model.ConceptGoal {
	// fmt.Printf("\n\nIn fillClarityQuestions start%send\n\n", input)
	// Regular expressions to extract questions and goals.
	questionRegex := regexp.MustCompile(`Question \d+: (.+)`)
	goalRegex := regexp.MustCompile(`Goal \d+: (.+)`)

	// Slice to store the extracted questions and goals.
	var visionQuestions []model.ConceptGoal

	// Split the input string by newlines to process each line separately.
	lines := strings.Split(input, "\n\n")
	// fmt.Println(lines)
	// Iterate over each line in the input string.
	for _, line := range lines {
		// Try to find a question and goal in the line.
		questionMatch := questionRegex.FindStringSubmatch(line)
		goalMatch := goalRegex.FindStringSubmatch(line)

		// If both a question and a goal are found in the line, add them to the visionQuestions slice.
		if len(questionMatch) > 1 && len(goalMatch) > 1 {
			visionQuestions = append(visionQuestions, model.ConceptGoal{
				Question: questionMatch[1],
				Goal:  goalMatch[1],
			})
		}
	}
	return visionQuestions
}

// Function to parse the text and fill the VisionGoal struct
func fillGoalandReason(input string) []*model.VisionGoal {
	var (
		visionGoals []*model.VisionGoal
		reasoningMatch []string
		goalMatch []string
	)
	reasoningRegex := regexp.MustCompile(`Reasoning \d+: (.+)`)
	goalRegex := regexp.MustCompile(`Goal \d+: (.+)`)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if strings.Contains(line, "Reasoning"){
			reasoningMatch = reasoningRegex.FindStringSubmatch(line)
		}else if strings.Contains(line, "Goal"){
			goalMatch = goalRegex.FindStringSubmatch(line)
		}
		if reasoningMatch != nil && goalMatch != nil {
			visionGoals = append(visionGoals, &model.VisionGoal{
				Reasoning: strings.TrimSpace(reasoningMatch[1]),
				Goal:      strings.TrimSpace(goalMatch[1]),
			})
		}
	}
	return visionGoals
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
