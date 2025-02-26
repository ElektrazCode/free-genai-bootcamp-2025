# Backend Building Steps with Copilot

## Prompt

I need to build an api backend for my lang portal app. It should be saved under the directory @backend_go. You will find all of the technical requirement specifications in the @Backend_Technical_Specs.md file. First, I would like you to give me a summary of what you plan to do. Once I agree to the plan, we will walk through every point in your plan, step by step. You will explain each step first then when I agree you can go ahead and execute it.

## Plan Summary

1. Setup Project Structure
	- Ensure the project structure is organized as specified.
	- Create necessary directories and files.
2. Initialize Go Modules
	- Initialize the Go module for dependency management.
	- Add required dependencies (Gin, SQLite3).
3. Database Setup
	- Implement database initialization in database/database.go.
	- Create migration scripts to set up the database schema in migrations/0001_init.sql.
4. API Endpoints Implementation
	- Implement the following API endpoints in main.go and corresponding handler files:
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
5. Task Runner Setup
	- Implement tasks for initializing, migrating, and seeding the database.
6. Testing
	- Write unit tests for the API endpoints.
	- Ensure the API meets the specifications and works as expected.

## Detailed Steps

### Step 1: Setup Project Structure

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

#### Explanation

- main.go: The entry point of the application.
- database/database.go: Contains the database initialization and connection logic.
- handlers/: Contains the handler functions for the API endpoints.
- migrations/: Contains the SQL migration files to set up the database schema.
- seeds/: Contains the seed data files.
- tasks/tasks.go: Contains the task runner logic for initializing, migrating, and seeding the database.

#### Execution

mkdir -p backend_go/database
mkdir -p backend_go/handlers
mkdir -p backend_go/migrations
mkdir -p backend_go/seeds
mkdir -p backend_go/tasks

touch backend_go/main.go
touch backend_go/database/database.go
touch backend_go/handlers/dashboard.go
touch backend_go/handlers/activities.go
touch backend_go/handlers/words.go
touch backend_go/handlers/groups.go
touch backend_go/handlers/sessions.go
touch backend_go/migrations/0001_init.sql
touch backend_go/seeds/words.json
touch backend_go/tasks/tasks.go

### Step 2: Initialize Go Modules

#### Explanation

1. Initialize the Go module for dependency management: This will create a go.mod file to manage dependencies.
2. Add required dependencies: We will add the necessary dependencies for the project, including Gin and SQLite3.

#### Execution

cd backend_go
go mod init backend_go
go get -u github.com/gin-gonic/gin
go get -u github.com/mattn/go-sqlite3

### Step 3: Database Setup

#### Explanation

1. Implement database initialization: We will create a database.go file to handle the database connection and initialization.
2. Create migration scripts: We will create a migration script 0001_init.sql to set up the database schema.

#### Execution

1. Create database.go (check file for content)
2. Create 0001_init.sql (check file for content)

#### Step 4: API Endpoints Implementation
We will implement the following API endpoints in main.go and corresponding handler files:

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

We will start by setting up the basic structure in main.go and then implement each handler function in the corresponding files.

#### Execution

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

### Step 4: Task Runner Setup

#### Explanation

We will implement tasks for initializing, migrating, and seeding the database using Mage, a task runner for Go.

#### Execution

1. Setup Mage:
	- Install Mage if it is not already installed:
		- if says "no required module provide package"
		- go get github.com/magefile/mage

2. Create Magefile:
	- Create a Magefile.go in the tasks directory (Check file for content).

3. Run Mage tasks:
	- Run the Mage tasks to initialize and seed the database:
		- cd backend_go
		- mage init
		- mage seed
		- mage migrate