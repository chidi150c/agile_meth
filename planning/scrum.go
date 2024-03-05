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
	Vision          string
	UpdatedVision   string
	ConceptGoals    []model.ConceptGoal
	Backlog         *Backlog
	DeveloperEngine ai_model.AIServicer
	QuestionsEngine ai_model.AIServicer
	AnswersEngine   ai_model.AIServicer
	ScrumEngine     ai_model.AIServicer
	MappingEngine   ai_model.AIServicer
	PreSalesEngine  ai_model.AIServicer
	SummaryEngine  ai_model.AIServicer
	ResearchBackLog []string
	MapBuilder strings.Builder
	AllBuilder strings.Builder
}

func NewAgileProject(name string, deveng, srceng, queeng, anseng, preeng, mapeng, sumeng ai_model.AIServicer) *AgileProject {

	return &AgileProject{
		Backlog:         NewBacklog(),
		DeveloperEngine: deveng,
		QuestionsEngine: queeng,
		AnswersEngine:   anseng,
		ScrumEngine:     srceng,
		PreSalesEngine:  preeng,
		MappingEngine:   mapeng,
		SummaryEngine: sumeng,
	}
}
func (project *AgileProject) VisionArgumentation(vision string) string {
	questions := project.QuestionsForVisionClarification(vision)
	vi := ""
	i := 1
	project.AllBuilder.WriteString(fmt.Sprintf("\n\nBELOW ARE QUESTIONS AND ANSWERS TO CLARIFY THE FEATURES AND CHALLENGES EXPECTED FOR THE VISION"))
	for k, v := range questions {
		project.AllBuilder.WriteString(fmt.Sprintf("\n\nQuestion %d: %s\nGoal %d: %s \n", k+1, v.Question, k+1, v.Goal))
		v.Answer = project.AnswerQuestionForVisionClarity(v.Question)
		project.AllBuilder.WriteString(fmt.Sprintf("Answer %d: %s", k+1, v.Answer))
		vi = fmt.Sprintf("Answer: %s", v.Answer)
		if strings.Contains(vi, "Unknown:") {
			//send vi to R&D
			project.ResearchBackLog = append(project.ResearchBackLog, fmt.Sprintf("%s\n%s", v.Question, vi))
		} else {
			vi = strings.Replace(vi, "Answer:", "Feature "+fmt.Sprint(i)+":", -1)
			vision = fmt.Sprintf("%s\n%s", vision, vi)
			i++
		}
	}

	project.AllBuilder.WriteString(fmt.Sprintf("\n\nTHE FOLLOWING IS THE UPDATED VISION WITH SOME FEATURES:\n"))
	project.AllBuilder.WriteString(fmt.Sprintf("\n%s\n\n", vision))
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
func (project *AgileProject) AnswerQuestionForVisionClarity(question string) string {
	s := fmt.Sprintf("Vision : %s\nQuestion : %s", project.Vision, question)
	qaAnswer, err := project.AnswersEngine.ProcessAiMessage(s)
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Split(qaAnswer, "Answer: ")[1]
}

// Function to break down project vision into 5 goals
func (project *AgileProject) BreakDownVisionIntoGoals(vision string) []*model.VisionGoal {
	project.Vision = vision
	goals, err := project.ScrumEngine.ProcessAiMessage(vision)
	if err != nil {
		log.Fatalln(err)
	}
	project.Backlog.VisionGoals, project.UpdatedVision = fillGoalandReason(project.ScrumEngine.GetEngineName(), goals)
	
	return project.Backlog.VisionGoals
}
func (project *AgileProject) fillUserStory(description string)[]*model.UserStory {
	var (
		userStories    []*model.UserStory
		userStoryMatch []string
	)
	userStoryRegex := regexp.MustCompile(`As a user, (.+)`)
	lines := strings.Split(description, "\n\n")
	for _, line := range lines {
		userStoryMatch = userStoryRegex.FindStringSubmatch(line)
		if userStoryMatch != nil{
			userStories = append(userStories, &model.UserStory{
				Description: strings.TrimSpace(userStoryMatch[1]),
			})
		}
	}
	return userStories
}
// CreateUserStory creates a new user story with the given description and priority.
func (project *AgileProject) CreateUserStory(description string, priority int) []*model.UserStory {
	description, _ = project.PreSalesEngine.ProcessAiMessage(description)
	userStories := project.fillUserStory(description)
	return userStories
}
func (project *AgileProject) MapGoalAndUserStory(input string)[]model.Mapping{
	output, _ := project.MappingEngine.ProcessAiMessage(input)
	// Split the output into individual mappings
    mappings := strings.Split(output, "\n\n")

    // Initialize a slice to store parsed mappings
    parsedMappings := make([]model.Mapping, len(mappings))

    // Parse each mapping
    for i, mapping := range mappings {
        lines := strings.Split(mapping, "\n")

        // Extract user story, goal, reason, and concept from each mapping
        reason := strings.TrimSpace(strings.Split(lines[1], "Reason:")[1])
        concept := strings.TrimSpace(strings.Split(lines[2], "Concept:")[1])		
		ug := strings.Split(lines[0], ".")[1]     
		var userStory, goal string    
		if strings.Contains(ug, "does not directly align"){
			project.ResearchBackLog = append(project.ResearchBackLog, fmt.Sprintf("%s: %s\n", concept, reason))
		}else{
			userStory = strings.TrimSpace(strings.Split(ug, "maps to")[0])
			goal = strings.TrimSpace(strings.Split(ug, "maps to")[1])
		}

        // Populate the Mapping struct
        parsedMappings[i] = model.Mapping{
            UserStory: userStory,
            Goal:      goal,
            Reason:    reason,
            Concept:   concept,
        }
    }

    // Print parsed mappings
    for _, m := range parsedMappings {
        project.AllBuilder.WriteString(fmt.Sprintf("User Story: %s\nGoal: %s\nReason: %s\nConcept: %s\n\n", m.UserStory, m.Goal, m.Reason, m.Concept))
    }
	return parsedMappings 
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
	questionRegex := regexp.MustCompile(`Question \d+: (.+)`)
	goalRegex := regexp.MustCompile(`Goal \d+: (.+)`)

	// Slice to store the extracted questions and goals.
	var visionQuestions []model.ConceptGoal

	// Split the input string by newlines to process each line separately.
	lines := strings.Split(input, "\n\n")
	// Iterate over each line in the input string.
	for _, line := range lines {
		// Try to find a question and goal in the line.
		questionMatch := questionRegex.FindStringSubmatch(line)
		goalMatch := goalRegex.FindStringSubmatch(line)

		// If both a question and a goal are found in the line, add them to the visionQuestions slice.
		if len(questionMatch) > 1 && len(goalMatch) > 1 {
			visionQuestions = append(visionQuestions, model.ConceptGoal{
				Question: questionMatch[1],
				Goal:     goalMatch[1],
			})
		}
	}
	return visionQuestions
}

// Function to parse the text and fill the VisionGoal struct
func fillGoalandReason(name, input string) ([]*model.VisionGoal, string) {
	var (
		visionGoals    []*model.VisionGoal
		reasoningMatch []string
		goalMatch      []string
		visionMatch    []string
	)
	visionRegex := regexp.MustCompile(`Vision: (.+)`)
	reasoningRegex := regexp.MustCompile(`Reasoning \d+: (.+)`)
	goalRegex := regexp.MustCompile(`Goal \d+: (.+)`)
	visionMatch = visionRegex.FindStringSubmatch(input)
	lines := strings.Split(input, "\n\n")
	for _, line := range lines {
		reasoningMatch = reasoningRegex.FindStringSubmatch(line)
		goalMatch = goalRegex.FindStringSubmatch(line)
		if reasoningMatch != nil && goalMatch != nil {
			visionGoals = append(visionGoals, &model.VisionGoal{
				Reasoning: strings.TrimSpace(reasoningMatch[1]),
				Goal:      strings.TrimSpace(goalMatch[1]),
			})
		}
	}
	return visionGoals, strings.TrimSpace(visionMatch[1])
}
