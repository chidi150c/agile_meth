package model

type VisionGoal struct{
	Goal string
	Reasoning string
}
type ConceptGoal struct{
	Question string
	Answer string
	Goal string
	Intervention bool
}
type ConceptGoalsReasoning struct{
	Reasoning string
	BaseVision string
	Goals []VisionGoal
}
type Mapping struct {
    UserStory string
    Goal      string
    Reason    string
    Concept   string
}
