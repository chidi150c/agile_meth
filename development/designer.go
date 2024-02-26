package development

import "fmt"

// Designer struct representing a designer team member
type Designer struct {
	Name string
}

// Implement the WorkOn method for Designer
func (d Designer) WorkOn(userStory string) {
	fmt.Printf("Designer %s is designing user story: %s\n", d.Name, userStory)
}