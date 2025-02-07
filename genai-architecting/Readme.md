## Architectural Design and Considerations

## Requirements

### 1. Business Requirements:

- Build a Gen AI app that will extend school language offering and augment the learning experience for students.

- Build a collection of learning apps using various different use-cases of AI.

- Maintain the learning experience from the existing learning portal using AI developer tools.

- This app should run in the cloud since the school does not own powerful hardware to run the model.

### 2. Functional Requirements:

- User Interface: Students homepage should be a personalized dashboard that includes their level/score and activities. They should be able to choose an activity from various options to practice (flashcards, game, reading...).

- Authentication: Students should be able to log in to their accounts.

- Core Features: Students should have personalized learning sessions depending on their level.

### 3. Non-functional Requirements:

- Performance: This app should be available 99.9% of the time to all 200 students.
- Scalability: This app will serve 200 students in the city of Vancouver learning french as a second language. This app must be able to accomodate the current number of active students with the ability to scale up with growing number of students.
- Security & Privacy: This app should put into consideration protecting the students information and the copyrighted material used.
- Usability: App should be user-friendly to serve both kids and adults. It should be easy to understand and navigate.

### 4. Tooling:

- GenAI (LLM)

## Risks:

- There are privacy and copyright risks that must be taken into account by applying responsible AI principles.

## Assumptions:

- Assuming that free-tier open-source LLMs can be powerful enough for the initial trial of the project until we set a budget and take it to the next level.
- Assuming that cloud resources will be able to serve 200 students.

## Constraints:

- Small budget for intial launch.
- Limited hardware compute capabilities.
