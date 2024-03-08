package planning

import (
	"ai_agents/agile_meth/ai_model"
	"ai_agents/agile_meth/model"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Project represents an Agile project with its backlog of user stories.
type Project struct {
	ID              int
	 Vision          *model.Vision
	// Backlog         *Backlog
	AI ai_model.AIModelServicer
	AllBuilder strings.Builder
	ResearchBackLog []string
}

func NewProject(ai ai_model.AIModelServicer) *Project {
	return &Project{
		// Backlog:         NewBacklog(),
		AI: ai,
		Vision: &model.Vision{
			Goals: make(map[string]*model.Goal),
			UserStories: make(map[string]*model.UserStory),
		},
	}
}

// DeriveGoalsFromVision simulates the process of deriving goals from the given vision using the specified AI model.
func (pj *Project) DeriveGoalsFromVision(vision *model.Vision) (*model.Vision, []model.Goal, error) {
	pj.AllBuilder.Reset()
	goslsForQuestion := pj.QuestionsForVisionClarification(vision)
	vi := ""
	i := 1
	pj.AllBuilder.WriteString("\n\nBELOW ARE QUESTIONS AND ANSWERS TO CLARIFY THE FEATURES AND CHALLENGES EXPECTED FOR THE VISION")
	for k, v := range goslsForQuestion {
		pj.AllBuilder.WriteString(fmt.Sprintf("\n\nQuestion %d: %s\nGoal %d: %s \n", k+1, v.Question, k+1, v.Description))
		v.Answer = pj.AnswerToQuestionForVisionClarity(v.Question)
		pj.AllBuilder.WriteString(fmt.Sprintf("Answer %d: %s", k+1, v.Answer))
		vi = fmt.Sprintf("Feature %d: %s", i, v.Answer)
		if strings.Contains(vi, "Unknown:") {
			//send vi to R&D
			pj.ResearchBackLog = append(pj.ResearchBackLog, fmt.Sprintf("%s\n%s", v.Question, vi))
		} else {
			vision.Description = fmt.Sprintf("%s\n%s", vision.Description, vi)
			i++
		}
	}
	
	pj.AllBuilder.WriteString("\n\nTHE FOLLOWING IS THE UPDATED VISION WITH SOME FEATURES:\n")
	pj.AllBuilder.WriteString(fmt.Sprintf("\n%s\n\n", vision.Description))
	var goals []model.Goal
	goals, vision = pj.BreakDownVisionIntoGoals(vision)
	pj.AllBuilder.WriteString(fmt.Sprintf("THE FOLLOWING ARE %d GOALS TO ACHIEVE THE VISION: %s\n\n", len(goals), vision.UpdatedVision))
	pj.AllBuilder.WriteString("Goals:\n")
	for k, v := range vision.DarftedGoals {
		pj.AllBuilder.WriteString(fmt.Sprintf("%d. %s\n", k+1, v.Description))
	}
	fmt.Println((pj.AllBuilder.String()))
	pj.Vision = vision
	return vision, goals, nil
}
// CreateUserStory creates a new user story with the given description and priority.
func (pj *Project) CreateUserStoriesForTheVision(description string) ([]model.UserStory, error){
	PreSalesPrompt :=`
	TASK DESCRIPTION:
	You are a helpful Agile Development Team that helps to Identify User Activities and Creating User Stories for a given vision and Prioritizing the user stories based on their importance and impact on achieving the project goals, considering factors such as user needs, feasibility, and alignment with the project vision.
	
	CRITERIA:
	1. Identify all the activities that users will perform while using the Product of the VISION and for each user activity, break them down into smaller actionable items called user stories each preceded with "As a user," as label.
	2. We then prioritize these user stories based on their importance and impact on achieving the project goals, considering factors such as user needs, feasibility, and alignment with the project vision: with "Priority:" as label.
	
	RESPONSE FORMAT:
	For each user story:
	- The user story
	- Its Priority

	EXAMPLE OUTPUT:]
	1. As a user, I want to experience immersive graphics in the game so that I can be fully engaged in the virtual environment."
	Priority: High

	2. As a user, I want the game to have accessibility features for players with disabilities so that everyone can enjoy the game.
	Priority: Medium

	...

	8. As a user, I want diverse cultural elements and representation in the game world, to feel connected and represented.    
	Priority: Low

	...`
	var err error
	description, err = pj.AI.PromptAI(PreSalesPrompt, description)
	if err != nil{
		return nil, err
	}
	userStories := pj.fillUserStory(description)
	pj.AllBuilder.WriteString(fmt.Sprintf("THE FOLLOWING ARE %d USER STORIES OF THE VISION: %s\n\n", len(userStories), pj.Vision.UpdatedVision))
	pj.AllBuilder.WriteString("\nUser Stories:\n")
	for k, v := range userStories {
		pj.AllBuilder.WriteString(fmt.Sprintf("%d. As a user, %s\n", k+1, v.Description))
	}
	pj.AllBuilder.WriteString("\n")
	pj.Vision.DraftedUserStories = userStories
	fmt.Println((pj.AllBuilder.String()))
	return pj.Vision.DraftedUserStories, nil
}

func (pj *Project) MapUserStoriesToGoals() {
	pj.AllBuilder.Reset()
	pj.AllBuilder.WriteString("Goals:\n")
	for k, v := range pj.Vision.DarftedGoals {
		pj.AllBuilder.WriteString(fmt.Sprintf("%d. %s\n", k+1, v.Description))
	}
	pj.AllBuilder.WriteString("\n")
	pj.AllBuilder.WriteString("\nUser Stories:\n")
	for k, v := range pj.Vision.DraftedUserStories {
		pj.AllBuilder.WriteString(fmt.Sprintf("%d. As a user, %s\n", k+1, v.Description))
	}
	pj.AllBuilder.WriteString("\n")

	input := pj.AllBuilder.String()

	MappingPrompt := `
	TASK DESCRIPTION:
	Generate mappings between given user stories and goals outlined for a software development project, including reasoning for each mapping. Explicitly include the reasoning with a 'Reasoning:' label. Additionally, identify and label the underlying concepts that guide the project's development efforts with a 'Concept:' label.

	INPUT FORMAT:
	Provide a list of user stories and goals. Each user story and goal should be accompanied by a brief description.

	CRITERIA:
	The model should map each user story to the corresponding goal from the provided list and explicitly provide the reasoning behind the mapping with a 'Reasoning:' label. If a user story does not directly align with any goal, the model should indicate so with an explanation also preceded by the 'Reasoning:' label. Additionally, the model should identify and label the underlying concepts that guide the project's development efforts with a 'Concept:' label.

	RESPONSE FORMAT:
	For each mapping, the model's response should include:
	- The user story and the corresponding goal.
	- A 'Reasoning:' label followed by the rationale for the mapping.
	- A 'Concept:' label followed by the underlying concept that guides the project's development efforts.
	If a user story does not align with any goal, it should be flagged as an isolated user story with reasonings for the lack of alignment, also following the 'Reasoning:' label.

	EXAMPLE INPUT:
	User Stories:
	1. As a user, I want to be able to select between trend following and mean reversion strategies so that I can optimize my cryptocurrency trading based on my preferred investment style and market conditions.
	2. As a user, I want to set the time horizon for my trades (short-term, medium-term, or long-term) to align with my risk tolerance and profit goals.
	Goals:
	1. Develop a trading system that supports both trend following and mean reversion strategies.
	2. Integrate trading options for various timeframes, including short-term, medium-term, and long-term strategies.

	EXAMPLE OUTPUT:
	1. User Story 1 maps to Goal 1.
	Reasoning: The desire to select between trading strategies directly supports the goal of developing a trading system with versatile strategy options.
	Concept: Versatility in trading strategies.
	...
	3. User Story 2 does not directly align with any specific goal.
	Reasoning: The request for personalized recommendations suggests a need for an adaptive AI component, which is not explicitly covered in the outlined goals. This might indicate a new area for development or an extension of existing goals.
	Concept: Adaptive AI recommendations.

	ADDITIONAL INSTRUCTIONS:
	Ensure the reasoning is clear, concise, and directly connects the user's needs or desires with the goals' intent and outcomes. Use the 'Reasoning:' label to clearly separate the rationale from the mapping for easy reading and comprehension. Similarly, use the 'Concept:' label to explicitly identify the underlying concepts guiding the project's development efforts.
`
	output, _ := pj.AI.PromptAI(MappingPrompt, input)
	// Split the output into individual mappings
	fmt.Println(output)
	mappings := strings.Split(output, "\n\n")
	// Initialize a slice to store parsed mappings
	var storyMappings []model.UserStory
	var goalMappings []model.Goal
	// Parse each mapping
	for _, mapping := range mappings {
		lines := strings.Split(mapping, "\n")
		// Extract user story, goal, reasoning, and concept from each mapping
		reasoning := strings.TrimSpace(strings.Split(lines[1], "Reasoning:")[1])
		concept := strings.TrimSpace(strings.Split(lines[2], "Concept:")[1])
		ug := strings.Split(lines[0], ".")[1]
		if strings.Contains(ug, "does not directly align") {
			pj.ResearchBackLog = append(pj.ResearchBackLog, fmt.Sprintf("%s: %s\n", concept, reasoning))
		} else {
			reGoal, _ := regexp.Compile(`Goal (\d+)`)
			reUStory, _ := regexp.Compile(`User Story (\d+)`)

			gid := reGoal.FindStringSubmatch(ug)[1]
			uid := reUStory.FindStringSubmatch(ug)[1]

			nuid, err := strconv.Atoi(uid)
			if err != nil{
				log.Fatalln(err)
			}
			ngid, err := strconv.Atoi(gid)
			if err != nil{
				log.Fatalln(err)
			}
			
			// Populate the Mapping struct
			storyMappings = append(storyMappings, model.UserStory{
				ID: nuid,
				MappedGoals: ngid,
				Reasoning:          reasoning,
				Concept:         concept,
			})
			// Populate the Mapping struct
			goalMappings = append(goalMappings, model.Goal{
				ID: ngid,
				MappedUserStories: nuid,
				Reasoning:          reasoning,
				Concept:         concept,
			})
		}
		
	}
	
	// Print parsed mappings
	for _, m := range storyMappings {
		pj.Vision.DraftedUserStories[m.ID-1].Reasoning = m.Reasoning
		pj.Vision.DraftedUserStories[m.ID-1].Concept = m.Concept
		pj.Vision.UserStories[m.Concept] = &pj.Vision.DraftedUserStories[m.ID-1]
		fmt.Printf("\nUser Story: %d\nGoal: %d\nReasoning: %s\nConcept: %s\n\n", m.ID, m.MappedGoals, m.Reasoning, m.Concept)
		// pj.AllBuilder.WriteString(fmt.Sprintf("User Story: %d\nGoal: %s\nReasoning: %s\nConcept: %s\n\n", m.ID, m.MappedGoals, m.Reasoning, m.Concept))
	}
	for _, m := range goalMappings {
		pj.Vision.DarftedGoals[m.ID-1].Reasoning = m.Reasoning
		pj.Vision.DarftedGoals[m.ID-1].Concept = m.Concept
		pj.Vision.Goals[m.Concept]	= &pj.Vision.DarftedGoals[m.ID-1]
		fmt.Printf("\nGoal: %d\nUser Story: %d\nReasoning: %s\nConcept: %s\n\n", m.ID, m.MappedUserStories, m.Reasoning, m.Concept)
		// pj.AllBuilder.WriteString(fmt.Sprintf("User Story: %d\nGoal: %s\nReasoning: %s\nConcept: %s\n\n", m.ID, m.MappedGoals, m.Reasoning, m.Concept))
	}
}

// DeriveTasksFromGoal simulates the process of deriving tasks from the given goal.
func (pj *Project) DeriveTasksFromGoal(goal model.Goal) []model.Task {
	// Prompt AI model to derive tasks based on the goal
	
	TasksPrompt := `TASK DESCRIPTION:
	Break down the provided Goal into smaller, actionable tasks or sub-goals, that can be easily managed and implemented by an Agile Development Team. This will help in planning, tracking progress, and ensuring that each component of the goal is addressed effectively.

	INPUT FORMAT:
	[I, not you, will provide the Goal to be broken down into tasks]

	CRITERIA:
	- You will break down the goal into smaller, actionable tasks or sub-goals.
	- Tasks or sub-goals must be clear, concise, and directly related to the original goal.
	- Include any relevant details or considerations needed to understand and implement the task.
	- Prioritize tasks or sub-goals that are critical for the goal's completion.
	- Identify dependencies between tasks or sub-goals where applicable.

	RESPONSE FORMAT:
	- List the smaller, actionable tasks or sub-goals.
	- Provide a brief description of each task or sub-story.
	- If applicable, note any dependencies between tasks.
	- Highlight any tasks that are particularly critical for achieving the goal's objective.

	EXAMPLE OUTPUT:
	1. Design a login interface.
	2. Implement authentication mechanism (e.g., OAuth, JWT).
	3. Develop multi-factor authentication feature.
	4. Create user session management.
	5. Implement encryption for user credentials.
	6. Test the login process for security vulnerabilities.
	...

	ADDITIONAL INSTRUCTIONS:
	- Ensure that the breakdown of the goal is comprehensive and covers all necessary aspects to achieve the goal's objective.
	- Consider the technical and resource constraints that may influence the implementation of tasks.
	- Use clear and understandable language to ensure that all team members can easily grasp the tasks and their purposes.
	`
	response, err := pj.AI.PromptAI(TasksPrompt, goal.Description)
	if err != nil{
		log.Fatalln(err)
	}
	// Split the response into individual tasks
	taskDescriptions := strings.Split(response, "\n")
	var tasks []model.Task
	for _, desc := range taskDescriptions {
		tasks = append(tasks, model.Task{Description: strings.TrimSpace(desc)})
	}
	return tasks
}

func (pj *Project) BreakDownUserStoryIntoTasks() {
	var TaskBuilder strings.Builder
	TaskBuilder.WriteString(fmt.Sprintf("Vision: %s\n\n", pj.Vision.UpdatedVision))
	k :=1
	for _, v := range pj.Vision.UserStories {
		TaskBuilder.WriteString(fmt.Sprintf("User Story %s: %s\n\n", k, v.Description))
	}

	TasksPrompt := `TASK DESCRIPTION:
	Break down provided user stories into smaller, actionable tasks or sub-stories that can be easily managed and implemented by an Agile Development Team. This will help in planning, tracking progress, and ensuring that each component of the user story is addressed effectively.

	INPUT FORMAT:
	Provide a list of user stories. Each user story should be clearly stated, focusing on what the end-user wants and why. User stories should be formatted and numbered for easy reference.

	CRITERIA:
	- Each user story should be broken down into smaller, actionable tasks or sub-stories.
	- Tasks or sub-stories must be clear, concise, and directly related to the original user story.
	- Include any relevant details or considerations needed to understand and implement the task.
	- Prioritize tasks or sub-stories that are critical for the user story's completion.
	- Identify dependencies between tasks or sub-stories where applicable.

	RESPONSE FORMAT:
	For each user story:
	- List the smaller, actionable tasks or sub-stories.
	- Provide a brief description of each task or sub-story.
	- If applicable, note any dependencies between tasks.
	- Highlight any tasks that are particularly critical for achieving the goal of the user story.

	EXAMPLE INPUT:
	User Story 1: As a user, I want to securely log in to the application so that I can access my personal dashboard.

	User Story 2: As an admin, I want to be able to review user activity logs to ensure system security and compliance.

	EXAMPLE OUTPUT:
	Breakdown for User Story 1:
		1. Design a login interface.
		2. Implement authentication mechanism (e.g., OAuth, JWT).
		3. Develop multi-factor authentication feature.
		4. Create user session management.
		5. Implement encryption for user credentials.
		6. Test the login process for security vulnerabilities.

	Breakdown for User Story 2:
		1. Develop functionality to capture user activity logs.
		2. Design and implement a log review interface for admins.
		3. Integrate filtering and search capabilities for logs.
		4. Implement access controls to ensure only admins can review logs.
		5. Test log review features for usability and security.

	ADDITIONAL INSTRUCTIONS:
	- Ensure that the breakdown of each user story is comprehensive and covers all necessary aspects to achieve the user story's objective.
	- Consider the technical and resource constraints that may influence the implementation of tasks.
	- Use clear and understandable language to ensure that all team members can easily grasp the tasks and their purposes.
	`
	tasks, err := pj.AI.PromptAI(TasksPrompt,"")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("TTTTTaslks=: ", tasks)
	tasksOfEachUserStories := strings.Split(tasks, "\n\n")
	for k, v := range tasksOfEachUserStories {
		_,_ =k,v
		// fmt.Println(v)
		// tasksOfAUserStory := strings.Split(v, "\n")
		// for y, v := range tasksOfAUserStory {
		// 	fmt.Println(v)
		// 	pj.Vision.UserStories[k].Tasks[]&model.Task{
		// 		ID:          y,
		// 		Description: v,
		// 	}
		// }
	}
}


func (pj *Project) QuestionsForVisionClarification(vision *model.Vision) []*model.Goal{
	QuestionPrompt := `
    TASK DESCRIPTION:
    As my diligent assistant, your role is pivotal in driving and determining the goals necessary to actualize the vision I've provided. Your questions will serve to guide us towards the realization of the vision.

    INPUT FORMAT:
    Vision: [My vision comes here]

    CRITERIA:
	1. To clearly understand the vision or identify features and/or options specific to actualizing the vision, ask relevant questions that leads to realization of the vision.
    2. Each question should be aimed at clarifying specific aspects of the vision or propose a potential objectives that aligns with the goals of the vision..
    3. Pose a minimum of 5 and not more than 10 questions that directly contribute to identifying goals crucial for the actualization of the provided vision.
    4. Focus solely on generating insightful inquiries to guide the goal-setting process.

    RESPONSE FORMAT:
    Reasoning: [Your reasoning behind the question that contributes to identifying crucial goals for actualizing the vision]
    Question 1: [Your question here]
    Question 2: [Your question here]
    ...
    
    ADDITIONAL INSTRUCTIONS:
    Ensure your reasoning is directly related to the vision statement and highlights the importance of the questions in guiding the goal-setting process.

	`
	goals, err := pj.AI.PromptAI(QuestionPrompt, vision.Description)
	if err != nil {
		log.Fatalln(err)
	}
	gl := fillClarityQuestions(goals)
	return gl
}
// AnswerToQuestionForVisionClarity function generates an answer for the question using LLM.
func (pj *Project) AnswerToQuestionForVisionClarity(question string) string {
	input := fmt.Sprintf("Vision : %s\nQuestion : %s", pj.Vision.Description, question)
	AnswersPrompt := `
	TASK DESCRIPTION:
	As my dedicated assistant, I'm seeking your expertise to clarify the vision for our project. Your insightful answers to questions related to the project's vision will guide us through this process.

	CRITERIA:
	1. I'll share the project's vision statement with you and pose a question based on it.
	2. I trust your judgment to select the most efficient approach from multiple approaches/options to achieve the vision.
	3. Ensure your responses are very brief and begin with "Answer:" for clarity.
	4. Your answers should reflect a proactive approach towards achieving the vision.
	5. In situations where there are several approaches/choices to decision/choose from, decide on or choose a specific approach/choice that would lead us towards achieving the vision fastest with less resources.
	6. In situations where specific information is lacking, simply respond with "Answer: Unknown:".
	7. In situations where no specific decision/choice can be made, simply respond with "Answer: Unknown:" and number out the choices to select from.
	8. Ensure that every response contributes directly to realizing the project's vision.
	
	RESPONSE FORMAT:
	Answer: [Unknown:][Your answer here][Numbered choices]
	`
	qaAnswer, err := pj.AI.PromptAI(AnswersPrompt, input)
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Split(qaAnswer, "Answer: ")[1]
}

// Function to break down pj vision into 5 goals
func (pj *Project) BreakDownVisionIntoGoals(vision *model.Vision)([]model.Goal, *model.Vision){
	ScrumPrompt := `
	TASK DESCRIPTION:
	You are a helpful assistant that tells me the goals required to achieve a given project Vision. The ultimate goal is to actualize the Vision by accoplishing these goals. Clearly articulate the Vision for the Project and outline specific goals; what to do in order to achieve the Vision.

	INPUT FORMAT:
	Vision: [I will place the vision statement here]
	
	CRITERIA:
	1. You should act as an agile scrum master.
	2. Clearly articulate the Vision for the Project.
	3. Outline specific goals; what to do in order to achieve the Vision.
	4. Generate atleast 5 Unique Goals and not more than 10 Goals, each in a breif sentence.
 
	RESPONSE FORMAT:
	Vision: [Place the clearly articulated vision statement here]
	Reasoning 1: [Your reasoning for the first goal here]
	Goal 1: [Your first goal to realize the vision here]
	Reasoning 2: [Your reasoning for the next goal here]
	Goal 2: [Your next goal to realize the vision here]
	Reasoning 3: [Your reasoning for the next goal here]
	Goal 3: [Your next goal to realize the vision here]
	...`

	goals, err := pj.AI.PromptAI(ScrumPrompt, vision.Description)
	if err != nil {
		log.Fatalln(err)
	}
	vgoals, updateVision := fillGoalandReasoning(goals)
	vision.UpdatedVision = updateVision //This is a clearly acticulated vision
	vision.DarftedGoals = vgoals
	return vgoals, vision
}

func fillClarityQuestions(input string) []*model.Goal {
    // Define regular expressions to match question and reasoning patterns.
    questionRegex := regexp.MustCompile(`Question \d+: (.+)`)
    reasoningRegex := regexp.MustCompile(`Reasoning : (.+)`)

    // Slice to store the extracted questions and reasoning.
    var visionQuestions []*model.Goal

    // Split the input string by newlines to process each line separately.
    lines := strings.Split(input, "\n\n")

    // Iterate over each line in the input string.
    for _, line := range lines {
        // Try to find a question and reasoning in the line.
        questionMatch := questionRegex.FindStringSubmatch(line)
        reasoningMatch := reasoningRegex.FindStringSubmatch(line)

        // If both a question and reasoning are found in the line, add them to the visionQuestions slice.
        if len(questionMatch) > 1 && len(reasoningMatch) > 1 {
            visionQuestions = append(visionQuestions, &model.Goal{
                Question:  questionMatch[1],
                Reasoning: reasoningMatch[1],
            })
        } else if len(questionMatch) > 1 {
            // If only a question is found, add it to the visionQuestions slice without reasoning.
            visionQuestions = append(visionQuestions, &model.Goal{
                Question: questionMatch[1],
            })
        }
    }
    return visionQuestions
}


// Function to parse the text and fill the VisionGoal struct
func fillGoalandReasoning(input string) ([]model.Goal, string) {
	var (
		visionGoals    []model.Goal
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
			visionGoals = append(visionGoals, model.Goal{
				Reasoning: strings.TrimSpace(reasoningMatch[1]),
				Description:      strings.TrimSpace(goalMatch[1]),
			})
		}
	}
	return visionGoals, strings.TrimSpace(visionMatch[1])
}

func (pj *Project) fillUserStory(description string) []model.UserStory {
	var (
		userStories    []model.UserStory
		userStoryMatch []string
		priorityMatch  []string
	)
	userStoryRegex := regexp.MustCompile(`As a user, (.+)`)
	priorityRegex := regexp.MustCompile(`Priority: (.+)`)
	lines := strings.Split(description, "\n\n")
	for _, line := range lines {
		userStoryMatch = userStoryRegex.FindStringSubmatch(line)
		priorityMatch = priorityRegex.FindStringSubmatch(line)
		if userStoryMatch != nil && priorityMatch != nil {
			userStories = append(userStories, model.UserStory{
				Description: strings.TrimSpace(userStoryMatch[1]),
				Priority:    strings.TrimSpace(priorityMatch[1]),
			})
		}
	}
	return userStories
}


















