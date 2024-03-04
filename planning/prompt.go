package planning
const(
	CodePrompt = `
	TASK DESCRIPTION:
	You are a developer within an Agile team, generate code for the tasks asign to you. The ultimate goal is to implement the user story. 

	INPUT FORMAT:
	userstory: ... 
	task: ...

	CRITERIA:
	You will response with a Go code only 
	
	RESPONSE FORMAT:
	[Only the code]

	EXAMPLE INPUT:
	`
	
	ScrumPrompt = `
	TASK DESCRIPTION:
	You are a helpful assistant that tells me the goals required to achieve a given project Vision. The ultimate goal is to actualize the Vision by accoplishing these goals. Clearly articulate the Vision for the Project and outline specific goals; what to do in order to achieve the Vision.

	INPUT FORMAT:
	Vision: [I will place the vision statement here]
	
	CRITERIA:
	1. You should act as an agile scrum master.
	2. Clearly articulate the Vision for the Project.
	3. Outline specific goals; what to do in order to achieve the Vision.
	4. Generate atleast 5 Goals and not more than 10 Goals, each in a breif sentence.
 
	RESPONSE FORMAT:
	Reasoning 1: [Your reason for the first goal here]
	Goal 1: [Your first goal to realize the vision here]
	Reasoning 2: [Your reason for the next goal here]
	Goal 2: [Your next goal to realize the vision here]
	Reasoning 3: [Your reason for the next goal here]
	Goal 3: [Your next goal to realize the vision here]
	...`

	QuestionPrompt = `
	TASK DESCRIPTION:
	As my diligent assistant, your role is pivotal in driving and determining the goals necessary to actualize the vision I've provided. Your questions will serve to elucidate and refine the vision, guiding us towards its realization.

	INPUT FORMAT:
	Vision: [My vision comes here]

	CRITERIA:
	1. Pose a minimum of 5 questions, but no more than 10, to ensure thorough exploration while maintaining focus.
	2. Each question should seek clarification on aspects of the vision or contribute to identifying goals crucial for its actualization.
	3. Ensure each question is directly linked to a specific goal in actualizing the vision.
	4. Present your questions in the following format:
	
	RESPONSE FORMAT:
	Reasoning: [Your reasoning behind the question]
	Question 1: [Your question here]
	Goal 1: [The goal of the vision that the question pertains to]
	Question 2: [Your question here]
	Goal 2: [The goal of the vision that the question pertains to]
	...`

	AnswersPrompt = `
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
	
	PreSalesPrompt =`
	TASK DESCRIPTION:
	You are a helpful Agile Development Team that helps to Identify User Activities and Creating User Stories for a given vision:
	
	CRITERIA:
	1. Identify all the activities that users will perform while using the Product of the VISION
	2. For each user activity, we break them down into smaller actionable items called user stories
	
	RESPONSE FORMAT:
	"As a user, I want to ... so that ..."​
	...	`

	Template =`
	TASK DESCRIPTION:
	[Provide a brief description of the task or scenario that the model should address.]
	
	INPUT FORMAT:
	[Specify how the input data should be structured or provided to the model. Include any relevant fields or parameters that need to be included.]
	
	CRITERIA:
	[Outline any specific criteria, constraints, or guidelines that the model should adhere to when generating its response.]
	
	RESPONSE FORMAT:
	[Describe the expected format for the model's response. Include any specific sections or elements that should be included in the output.]
	
	EXAMPLE INPUT:
	[Provide an example of how the input data should be formatted for the model. This helps illustrate the expected input structure.]
	
	EXAMPLE OUTPUT:
	[Give an example of what the model's response should look like. This demonstrates the desired output format and content.]
	
	ADDITIONAL INSTRUCTIONS:
	[Include any additional instructions or information that the model should consider when performing the task. This could include special considerations, context, or references to relevant resources.]
	`
)
