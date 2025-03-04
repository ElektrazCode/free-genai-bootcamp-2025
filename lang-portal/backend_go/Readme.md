# Backend Building with Copilot

### Prompt

I need to build an api backend for my lang portal app. It should be saved under the directory @backend_go. You will find all of the technical requirement specifications in the @Backend_Technical_Specs.md file. First, I would like you to give me a summary of what you plan to do. Once I agree to the plan, we will walk through every point in your plan, step by step. You will explain each step first then when I agree you can go ahead and execute it.

### Plan Summary

#### 1. Setup Project Structure
- Ensure the project structure is organized as specified.
- Create necessary directories and files.

#### 2. Initialize Go Modules
- Initialize the Go module for dependency management.
- Add required dependencies (Gin, SQLite3).

#### 3. Database Setup
- Implement database initialization in database/database.go.
- Create migration scripts to set up the database schema in migrations/0001_init.sql.
#### 4. API Endpoints Implementation
- Implement the following API endpoints in main.go and corresponding handler files:
```text
	- GET /api/dashboard/last_study_session
	- GET /api/dashboard/study_progress
	- GET /api/dashboard/quick-stats
	- GET /api/activities
	- GET /api/activities/:id
	- GET /api/activities/:id/sessions
	- POST /api/activities
	- GET /api/words
	- GET /api/words/:id
	- GET /api/groups
	- GET /api/groups/:id
	- GET /api/groups/:id/words
	- GET /api/groups/:id/sessions
	- GET /api/sessions
	- GET /api/sessions/:id
	- GET /api/sessions/:id/words
	- POST /api/reset_history
	- POST /api/full_reset
	- POST /api/sessions/:id/words/:word_id/review
```
#### 5. Task Runner Setup
- Implement tasks for initializing, migrating, and seeding the database.
#### 6. Testing
- Write unit tests for the API endpoints.
- Ensure the API meets the specifications and works as expected.

# Detailed Steps

## Step 1: Setup Project Structure

We will create the following directories and files:

- backend_go
    - main.go
    - database/
        - database.go
    - handlers/
        - dashboard.go
        - activities.go
        - words.go
        - groups.go
        - sessions.go
    - migrations/
        - 0001_init.sql
    - seeds/
        - words.json
    - tasks/
        - tasks.go

### Explanation

- `main.go`: The entry point of the application.
- `database/database.go`: Contains the database initialization and connection logic.
- `handlers/`: Contains the handler functions for the API endpoints.
- `migrations/`: Contains the SQL migration files to set up the database schema.
- `seeds/`: Contains the seed data files.
- `tasks/tasks.go`: Contains the task runner logic for initializing, migrating, and seeding the database.

### Execution
```sh
- mkdir -p backend_go/database
- mkdir -p backend_go/handlers
- mkdir -p backend_go/migrations
- mkdir -p backend_go/seeds
- mkdir -p backend_go/tasks

- touch backend_go/main.go
- touch backend_go/database/database.go
- touch backend_go/handlers/dashboard.go
- touch backend_go/handlers/activities.go
- touch backend_go/handlers/words.go
- touch backend_go/handlers/groups.go
- touch backend_go/handlers/sessions.go
- touch backend_go/migrations/0001_init.sql
- touch backend_go/seeds/words.json
- touch backend_go/tasks/tasks.go
```

## Step 2: Initialize Go Modules

### Explanation

1. Initialize the Go module for dependency management: This will create a go.mod file to manage dependencies.
2. Add required dependencies: We will add the necessary dependencies for the project, including Gin and SQLite3.

### Execution
```sh
- cd backend_go
- go mod init backend_go
- go get -u github.com/gin-gonic/gin
- go get -u github.com/mattn/- go-sqlite3
```

## Step 3: Database Setup

### Explanation

1. Implement database initialization: We will create a database.go file to handle the database connection and initialization.
2. Create migration scripts: We will create a migration script 0001_init.sql to set up the database schema.

### Execution

1. Create database.go (check file for content)
2. Create 0001_init.sql (check file for content)

### Step 4: API Endpoints Implementation
We will implement the following API endpoints in main.go and corresponding handler files:
```sh
	- GET /api/dashboard/last_study_session
	- GET /api/dashboard/study_progress
	- GET /api/dashboard/quick-stats
	- GET /api/activities
	- GET /api/activities/:id
	- GET /api/activities/:id/sessions
	- POST /api/activities
	- GET /api/words
	- GET /api/words/:id
	- GET /api/groups
	- GET /api/groups/:id
	- GET /api/groups/:id/words
	- GET /api/groups/:id/sessions
	- GET /api/sessions
	- GET /api/sessions/:id
	- GET /api/sessions/:id/words
	- POST /api/reset_history
	- POST /api/full_reset
	- POST /api/sessions/:id/words/:word_id/review
```
We will start by setting up the basic structure in main.go and then implement each handler function in the corresponding files.

### Execution

1. Setup basic structure in main.go:
	- Create the main.go file (Check file for content)
		- Initilize the database
		- Create a new Gin router
		- Define all routes
		- Start the server
2. Implement handler functions
	- Create the handler files. (Check files for content)
		1. `/dashboard.go` handlers:
			- GetLastStudySession 
			- GetStudyProgress
			- GetQuickStats
		2. `/activities.go` handlers:
			- GetActivities
			- GetActivityByID
			- GetActivitySessions
			- CreateActivity
		3. `/words.go` handlers:
			- GetWords
			- GetWordByID
		4. `/groups.go`
			- GetGroups
			- GetGroupByID
			- GetGroupWords
			- GetGroupSessions
		5. `/sessions.go` handlers:
			- GetSessions
			- GetSessionByID
			- GetSessionWords
			- ResetHistory
			- FullReset
			- ReviewWord

## Step 4: Task Runner Setup

### Explanation

We will implement tasks for initializing, migrating, and seeding the database using Mage, a task runner for Go.

### Execution

1. Setup Mage:
	- Install Mage if it is not already installed:
		- if says "no required module provide package"
		```sh
			go get github.com/magefile/mage
		```

2. Create Magefile:
	- Create a Magefile.go in the tasks directory (Check file for content).

3. Run Mage tasks:
	- Run the Mage tasks to initialize and seed the database:
	```sh
		- cd backend_go
		- mage init
		- mage seed
		- mage migrate
	```

# Missed Steps

## 1. Generate data for the Seed File:

### Prompt

Could you please generate the top 100 beginner level french vocabulary words in json format like the following:
```sh
{
	french: french_word,
	english: english_translation,
	parts: part_of_speech,
} 
```
### Execution

Check Words.json content for result.

## 2. Complete the Implementation of the Seed Function

### Explanation

To complete the Seed function, you need to:

1. Read the seed data from the JSON files in the seeds folder.
2. Parse the JSON data.
3. Insert the parsed data into the appropriate tables in the SQLite database.


### Execution 

Check the `MageFile.go` file to see the content of the `Seed()` function.

## 3. Complete the Implementation of the Migrate() Function

### Explanation

To complete the implementation of the Migrate function in the `MageFile.go` file, you need to add logic to run the migration scripts located in the migrations folder. These scripts will set up the database schema as specified in the `Backend_Technical_Specs.md` file.

This updated Migrate function does the following:

1. Ensures the database is initialized by calling the Init function.
2. Retrieves the list of migration files from the migrations folder.
3. Opens the SQLite database.
4. Reads and executes each migration file in order.

### Execution 

Check the `MageFile.go` file to see the content of the `Migrate()` function.

# Suggested Optimizations

The project appears to be well-structured and includes the necessary files and directories as outlined in the `Readme.md` and `Backend_Technical_Specs.md`. However, there are a few potential areas that might need attention:

#### Unit Tests:
There are no unit tests mentioned or included in the provided files. Adding unit tests for the API endpoints would be beneficial to ensure the API meets the specifications and works as expected.

#### Error Handling:
While the handlers include basic error handling, you might want to ensure that all possible errors are handled gracefully and consistently across the application.

#### Logging:
Consider adding more detailed logging throughout the application to help with debugging and monitoring.

#### Configuration Management:
If there are any configuration settings (e.g., database connection strings, server ports), it might be useful to externalize these into a configuration file or environment variables.

#### Documentation:
Ensure that the documentation is complete and up-to-date, including any additional setup steps, usage instructions, and API documentation.

#### Security:
Although the specs mention no authentication/authorization, consider if there are any other security measures needed, such as input validation to prevent SQL injection.

#### Pagination and Filtering:
The API endpoints for listing resources (e.g., /api/words, /api/groups) should support pagination and filtering to handle large datasets efficiently.

#### Performance Optimization:
Review the database queries and application logic for any potential performance bottlenecks.

# Testing

## Postman Extension in VSCode
I used the Postman extension to manually test all of the api endpoints.

For the time being, I still need to work out the logic for the dashboard and the word review. Otherwise, all of the apis were tested successfully.

## Using Copilot to generate Unit Tests

### Prompt
I would like to add unit tests in the directory tests/apis to test all of the api endpoints. You will find all api handlers in the backend_go/handlers directory. Could you walk me through each step? Explain to me what needs to be done, then after I approve you can go ahead and implement it.

### Response
1. Create a new directory for tests: We'll create a new directory named `tests/apis` inside the `backend_go` directory to store our test files.

2. Set up the test environment: We'll need to set up a test environment using the testing package in Go. This includes initializing a test database and creating a test router.
- Created the File `api_test.go`
	- This file sets up the test environment by initializing a test database and creating a test router with all the routes defined.
- Install testify
```sh
	go get github.com/stretchr/testify/assert
```

3. Write test cases for each endpoint: For each API endpoint, we'll write test cases to verify the expected behavior. This includes sending HTTP requests to the endpoints and checking the responses.

4. Run the tests: Finally, we'll run the tests to ensure everything is working as expected.

# Missed Steps

- Create the test_database, sample data, function to insert the sample data into the test_database.

- The tables were not matching the schema. I had to manually revise the insert queries to use the correct columns.

- Only 2 test cases were created for 2 endpoints. I had to ask for the remaining test cases to be done.

- The output was not clear I had to add some print lines to get more information about the reponse of each request.