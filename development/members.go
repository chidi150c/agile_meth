package development

import (
	"ai_agents/agile_meth/model"
	"ai_agents/agile_meth/openaiapi"
	"fmt"
	"log"
)

// Define interface for team members
type TeamMember interface {
	WorkOn(userStory string)
}

// Developer struct representing a developer team member
type Developer struct {
	Name string
	Aiapi   openaiapi.OpenAI                                     
}
func MewDeveloper (ai openaiapi.OpenAI)*Developer{
	return &Developer{
		Name: "",
		Aiapi: ai,
	}
}
// Implement the WorkOn method for Developer
func (d Developer) WorkOn(us *model.UserStory) {
	messages := []openaiapi.Message{
		{"system", fmt.Sprintf("Developer %s is working on the following user story:", d.Name)},
		{"user", fmt.Sprintf("User Story ID: %d", us.ID)},
		{"user", fmt.Sprintf("Description: %s", us.Description)},
		{"user", fmt.Sprintf("Priority: %d", us.Priority)},
		{"user", fmt.Sprintf("Estimated Effort: %d", us.EstimatedEffort)},
		{"user", fmt.Sprintf("Task: %s", us.Tasks[0].Description)},
		{"system", "Generate code for the task"},
	}
	output, err := d.Aiapi.ApiFetch(messages)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Printf("Developer %s is coding user story\n", d.Name)
	fmt.Printf("Now done with output as follows:\n %s \n", output)
}

// Tester struct representing a tester team member
type Tester struct {
	Name string
}

// Implement the WorkOn method for Tester
func (t Tester) WorkOn(userStory string) {
	fmt.Printf("Tester %s is testing user story: %s\n", t.Name, userStory)
}

// Designer struct representing a designer team member
type Designer struct {
	Name string
}

// Implement the WorkOn method for Designer
func (d Designer) WorkOn(userStory string) {
	fmt.Printf("Designer %s is designing user story: %s\n", d.Name, userStory)
}
