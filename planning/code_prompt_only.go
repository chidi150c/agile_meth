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
	Updated Prompt for Generating Go Code Based on Application Architecture Requirements:

Given:

A vision statement that outlines the purpose and key features of the application.
A comprehensive list of goals that the application aims to achieve.
Detailed tasks and sub-tasks that break down the steps to accomplish these goals.
An existing Go code snippet that includes "Models" and "Workers" representing the current state of the application architecture.
Objective:

Generate Go code that evolves the application architecture, focusing on:

Maintaining and Utilizing Existing Models and Workers: Enhance or modify existing "Models" (Go structs representing users or objects) and "Workers" (Go structs tasked with operations) based on the new requirements outlined in the vision statement, goals, tasks, and sub-tasks.

Creating New Models and Workers: Where existing structures do not suffice, introduce new "Models" and "Workers". These should seamlessly integrate with the existing codebase, maintaining the coherence and scalability of the application.

Output Requirements:

Go Structs for Models: For each identified model in the vision statement, define or update Go structs. Include relevant fields that accurately represent the model's role within the application.

Go Structs for Workers: For each worker, define or refine Go structs with methods that directly contribute to achieving the application's goals. These methods should interact with the models to enable the application's desired functionalities.

Code Best Practices:

Ensure the generated Go code:

Adheres to Go programming conventions and best practices.
Includes error handling where appropriate.
Contains documentation comments to improve code readability and maintenance.
This structured approach ensures the generated Go code is both functional and aligned with the application's strategic vision, facilitating the realization of its goals through well-defined tasks and sub-tasks.`


	CodePrompt3 = `
	Given a vision statement, list of goals, and detailed tasks and sub-tasks, generate Go code that defines the application architecture. Focus on creating Go structs for "Models" based on the users or objects highlighted in the vision statement. Then, define "Worker" structs that will serve these models, with their responsibilities encapsulated as methods corresponding to the goals, tasks, and sub-tasks provided.

Output Requirements:

Go Structs for Models: Define Go structs for each model identified in the vision statement. Include fields that are relevant to the model's representation in the application.

Go Structs for Workers: For each worker, create a Go struct with methods that reflect the application's goals, tasks, and sub-tasks. These methods should logically operate on or with the models, facilitating the desired functionality of the application.

Note: Ensure the Go code adheres to best practices, including appropriate naming conventions, error handling, and documentation comments for ease of understanding and maintenance.
	`

)