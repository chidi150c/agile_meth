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
// WriteTestCases simulates the tester writing test cases for the given user story
func (t *Tester) WriteTestCases(userStory string) {
    fmt.Printf("Tester %s is writing test cases for user story: %s\n", t.Name, userStory)
}

// ExecuteTests simulates the tester executing tests for the given user story
func (t *Tester) ExecuteTests(userStory string) {
    fmt.Printf("Tester %s is executing tests for user story: %s\n", t.Name, userStory)
}

// ReportBugs simulates the tester reporting bugs found during testing of the given user story
func (t *Tester) ReportBugs(userStory string, bugs []string) {
    fmt.Printf("Tester %s is reporting bugs for user story: %s\n", t.Name, userStory)
    for _, bug := range bugs {
        fmt.Printf(" - Bug: %s\n", bug)
    }
}
