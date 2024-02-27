package development

import (
	"ai_agents/agile_meth/model"
	"ai_agents/agile_meth/openai_api"
	"fmt"
	// "log"
)

// Define interface for team members
type TeamMember interface {
	WorkOn(userStory string)
}

// Developer struct representing a developer team member
type Developer struct {
	Name        string
    Engine openai_api.OpenAIServicer
}

func NewDeveloper(name string, engine openai_api.OpenAIServicer) *Developer {
    // _, err := engine.CreateAssistant(DeveloperInst, "Developer", "code_interpreter")
    // if err != nil{
    //     log.Fatalln(err)
    // }
    // _, err = engine.CreateThread()
    // if err != nil{
    //     log.Fatalln(err)
    // }
	return &Developer{
		Name:        name,
        Engine: engine,
	}
}

// CodeUserStory simulates the developer coding the given user story
func (d *Developer)CodeUserStory(us *model.UserStory) {
	// Example message content
	messageContent := "user story: " + us.Description + " immediate task: " + us.Tasks[0].Description
    resp, _ := d.Engine.ProcessAiMessage(messageContent)
	fmt.Println(resp["command"])
}
// DebugCode simulates the developer debugging the code for the given user story
func (d *Developer) DebugCode(userStory string) {
    fmt.Printf("Developer %s is debugging code for user story: %s\n", d.Name, userStory)
}
// IntegrateChanges simulates the developer integrating changes for the given user story
func (d *Developer) IntegrateChanges(userStory string) {
    fmt.Printf("Developer %s is integrating changes for user story: %s\n", d.Name, userStory)
}



