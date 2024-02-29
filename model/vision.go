package model


type VisionGoal struct{
	Goal string
	Reason string
}
type VisionClarityQuestion struct{
	Question string
	Answer string
	Goal string
	Intervention bool
}