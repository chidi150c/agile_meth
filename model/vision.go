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