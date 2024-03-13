package planning
const(
	CodePrompt1 = `
	Given a vision statement and list of goals, generate Go code that defines the application architecture. Focus on creating Go structs for "Models" based on the users or objects highlighted in the vision statement. Then, define "Worker" structs that will serve these models, with their responsibilities encapsulated as methods corresponding to the goals provided.

Output Requirements:

Go Structs for Models: Define Go structs for each model identified in the vision statement. Include fields that are relevant to the model's representation in the application.

Go Structs for Workers: For each worker, create a Go struct with methods that reflect the application's goals. These methods should logically operate on or with the models, facilitating the desired functionality of the application.

Note: Ensure the Go code adheres to best practices, including appropriate naming conventions, error handling, and documentation comments for ease of understanding and maintenance.
	`

	CodePrompt2 = `
	As my helpful assistant in building an application as the product of the provided vision statement, "Models" are Go structs created based on the users or objects highlighted in the vision statement, and "Worker" struct is a Go struct that will serve these models, whose responsibilities are encapsulated as methods corresponding to the provided goals, tasks, and sub-tasks. Focus on utilizing existing "Models" and "Workers" and only create new ones if necessary, in order to maintian coherence and scalability of the application. 

Given:
- A vision statement that outlines the purpose and key features of the application.
- A comprehensive list of goals that the application aims to achieve.
- Detailed tasks and sub-tasks that break down the steps to accomplish these goals.
- An existing Go code snippet that includes "Models" and "Workers" representing the current state of the application architecture.
 
Objective:
Generate Go code that evolves the application architecture, focusing on:
- Integrating goals, tasks and sub-tasks as methods to the Workers rendering services to the Models 
- Maintaining and Utilizing Existing Models and Workers: Enhance or modify existing "Models" and "Workers" based on the new requirements outlined in the vision statement, goals, tasks, and sub-tasks.
- Creating New Models and Workers: Where existing structures do not suffice, introduce new "Models" and "Workers". These should seamlessly integrate with the existing codebase, maintaining the coherence and scalability of the application.

Output Requirements:
- Go Structs for Models: For each identified model in the vision statement, include relevant fields that accurately represent the model's role within the application.
- Go Structs for Workers: For each worker, define or refine Go structs with methods that directly contribute to achieving the application's goals. These methods should interact with the models to enable the application's desired functionalities.
- Code Best Practices:

Ensure the generated Go code:
- Adheres to Go programming conventions and best practices.
- Includes error handling where appropriate.
- Contains documentation comments to improve code readability and maintenance.
- This structured approach ensures the generated Go code is both functional and aligned with the application's strategic vision, facilitating the realization of its goals through well-defined tasks and sub-tasks.
`


	CodePrompt3 = `
	As my helpful assistant, effectively assist me in building an application that stems from a vision, where "Models" and "Workers" are central to the application architecture. "Models" are Go structs created based on the users or objects highlighted in the vision statement, and "Worker" is a Go struct that will serve these models. 
	Goals, tasks and sub-tasks are provided for you to integrate as methods to the worker.
	Also provided are Code and it's execution output for you to fix any error or issues encounted during execution.  

Given: 
- A vision statement that outlines the purpose and key features of the application.
- A comprehensive list of goals that the application aims to achieve.
- Detailed tasks and sub-tasks that break down the steps to accomplish these goals.
- An existing Go code snippet that includes "Models" and "Workers" representing the current state of the application architecture.
- The output of excuting the code is also given with "Exec Output:" as label

Output Requirements:
- Fix errors observed during code excution as depicted by "Exec Output" label
- Evolved code with new goals, tasks and sub-tasks integrated as methods to the Workers rendering services to the Models  
- If necessary new fields added to existing Model structs. Include fields that are relevant to the model's representation in the application.
- For the worker structs, create methods that reflect the application's goals, tasks, and sub-tasks. These methods should logically operate on or with the models, facilitating the desired functionality of the application.

Note: Ensure the Go code adheres to best practices, including appropriate naming conventions, error handling, and documentation comments for ease of understanding and maintenance.
	`

	CodePrompt4 = `
	Given a vision statement, list of goals, and detailed tasks and sub-tasks, generate Go code that defines the application architecture. Focus on creating Go structs for "Models" based on the users or objects highlighted in the vision statement. Then, define "Worker" structs that will serve these models, with their responsibilities encapsulated as methods corresponding to the goals, tasks, and sub-tasks provided.

Output Requirements:

Go Structs for Models: Define Go structs for each model identified in the vision statement. Include fields that are relevant to the model's representation in the application.

Go Structs for Workers: For each worker, create a Go struct with methods that reflect the application's goals, tasks, and sub-tasks. These methods should logically operate on or with the models, facilitating the desired functionality of the application.

Note: Ensure the Go code adheres to best practices, including appropriate naming conventions, error handling, and documentation comments for ease of understanding and maintenance.
	`

)