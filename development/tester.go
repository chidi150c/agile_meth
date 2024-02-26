package development

import "fmt"

// Tester struct representing a tester team member
type Tester struct {
	Name string
}


// Implement the WorkOn method for Tester
func (t Tester) WorkOn(userStory string) {
	fmt.Printf("Tester %s is testing user story: %s\n", t.Name, userStory)
}