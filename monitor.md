TASK DESCRIPTION:
	Generate mappings between given user stories and goals outlined for a software development project, including reasoning for each mapping. Explicitly include the reasoning with a 'Reason:' label. Additionally, identify and label the underlying concepts that guide the project's development efforts with a 'Concept:' label. Prioritize the user stories based on their importance and impact on achieving the project goals, considering factors such as user needs, feasibility, and alignment with the project vision.

INPUT FORMAT:
	Provide a list of user stories and goals. Each user story and goal should be accompanied by a brief description.

CRITERIA:
	The model should prioritize the user stories based on their importance and impact on achieving the project goals, considering factors such as user needs, feasibility, and alignment with the project vision. The model should then map each prioritized user story to the corresponding goal from the provided list and explicitly provide the reasoning behind the mapping with a 'Reason:' label. If a user story does not directly align with any goal, the model should indicate so with an explanation also preceded by the 'Reason:' label. Additionally, the model should identify and label the underlying concepts that guide the project's development efforts with a 'Concept:' label.

RESPONSE FORMAT:
	For each prioritized user story:
	- The user story and its priority.
	- The user story and the corresponding goal.
	- A 'Reason:' label followed by the rationale for the mapping.
	- A 'Concept:' label followed by the underlying concept that guides the project's development efforts.
	If a user story does not align with any goal, it should be flagged as an isolated user story with reasons for the lack of alignment, also following the 'Reason:' label.

EXAMPLE INPUT:
	User Stories:
	1. As a user, I want to be able to select between trend following and mean reversion strategies so that I can optimize my cryptocurrency trading based on my preferred investment style and market conditions.
	2. As a user, I want to set the time horizon for my trades (short-term, medium-term, or long-term) to align with my risk tolerance and profit goals.
	Goals:
	1. Develop a trading system that supports both trend following and mean reversion strategies.
	2. Integrate trading options for various timeframes, including short-term, medium-term, and long-term strategies.

EXAMPLE OUTPUT:
	Prioritized User Stories:
	1. User Story 1 - High Priority
	2. User Story 2 - Medium Priority

	1. User Story 1 maps to Goal 1.
	Reason: The desire to select between trading strategies directly supports the goal of developing a trading system with versatile strategy options.
	Concept: Versatility in trading strategies.
	...
	3. User Story 2 does not directly align with any specific goal.
	Reason: The request for personalized recommendations suggests a need for an adaptive AI component, which is not explicitly covered in the outlined goals. This might indicate a new area for development or an extension of existing goals.
	Concept: Adaptive AI recommendations.

ADDITIONAL INSTRUCTIONS:
	Ensure the reasoning is clear, concise, and directly connects the user's needs or desires with the goals' intent and outcomes. Use the 'Reason:' label to clearly separate the rationale from the mapping for easy reading and comprehension. Similarly, use the 'Concept:' label to explicitly identify the underlying concepts guiding the project's development efforts. Prioritize the user stories based on their importance and impact on achieving the project goals, considering factors such as user needs, feasibility, and alignment with the project vision.
