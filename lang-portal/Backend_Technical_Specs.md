# Backend Server Technical Specs

## Business Goal

A language learning school wants to build a prototype of a learning portal which will act as three things:

- An inventory of possible vocabulary that can be learned.
- A learning record store (LRS), providing correct and wrong scores on practice vocabulary
- A unified launchpad to launch different learning activities

## Technical Requirements

- The backend will be built using Go
- The database will be SQLite3
- The API will be built using Gin
  -Mage is a task runner for Go.
- The API will always return JSON
- There will be no authentication/authorization
- The app will serve a single user

## Database Schema

Our database will be a single sqlite database called `words.db` that will be in the root of the project folder of `backend_go`

### Tables

- words - stored vocabulary words
  - id - integer
  - french - string
  - english - string
  - parts - json
- words_groups - many-to-many join table for words and groups
  - id - integer
  - word_id - integer
  - group_id - integer
- groups - thematic groups of words
  - id - integer
  - name - string
- sessions - records of study sessions grouping review_words
  - id - integer
  - group_id - integer
  - created_at - datetime
  - activity_id - integer
- activities - study activities, linking a study session to a group of words
  - id - integer
  - session_id - integer
  - group_id - integer
  - created_at - datetime
- review_words - a record of word practice, determining if the word was correct or not
  - word_id - integer
  - session_id - integer
  - correct - boolean
  - created_at - datetime

## API Endpoints

### GET /api/dashboard/last_study_session

Returns information about the most recent study session.

#### JSON Response

```json
{
  "id": 123,
  "group_id": 456,
  "created_at": "2025-02-08T17:20:23-05:00",
  "activity_id": 789,
  "group_id": 456,
  "group_name": "Basic Greetings"
}
```

### GET /api/dashboard/study_progress

Returns study progress statistics.
Please note that the frontend will determine progress bar based on total words studied and total available words.

#### JSON Response

```json
{
  "total_words_studied": 3,
  "total_available_words": 124
}
```

### GET /api/dashboard/quick-stats

Returns quick overview statistics.

#### JSON Response

```json
{
  "success_rate": 80.0,
  "total_sessions": 4,
  "total_active_groups": 3,
  "streak_days": 4
}
```

### GET /api/activities

Returns a list of study activities.

#### JSON Response

```json
{
  "items": [
    {
      "id": 1,
      "name": "Vocabulary Quiz",
      "thumbnail_url": "https://example.com/thumbnail.jpg",
      "description": "Practice your vocabulary with flashcards"
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "total_items": 1,
    "items_per_page": 100
  }
}
```

### GET /api/activities/:id

Returns details of a specific study activity.

#### JSON Response

```json
{
  "id": 1,
  "name": "Vocabulary Quiz",
  "thumbnail_url": "https://example.com/thumbnail.jpg",
  "description": "Practice your vocabulary with flashcards"
}
```

### GET /api/activities/:id/sessions

Returns a list of sessions for a specific activity.

```json
{
  "items": [
    {
      "id": 123,
      "activity_name": "Vocabulary Quiz",
      "group_name": "Basic Greetings",
      "start_time": "2025-02-08T17:20:23-05:00",
      "end_time": "2025-02-08T17:30:23-05:00",
      "review_words_count": 20
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_items": 100,
    "items_per_page": 20
  }
}
```

### POST /api/activities

Creates a new study activity.

#### Request Params

- group_id integer
- activity_id integer

#### JSON Response

```json
{
  "id": 124,
  "group_id": 123
}
```

### GET /api/words

Returns a list of words.

#### JSON Response

```json
{
  "items": [
    {
      "french": "salut",
      "english": "hello",
      "correct_count": 5,
      "wrong_count": 2
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_items": 500,
    "items_per_page": 100
  }
}
```

### GET /api/words/:id

Returns details of a specific word.

#### JSON Response

```json
{
  "french": "salut",
  "english": "hello",
  "stats": {
    "correct_count": 5,
    "wrong_count": 2
  },
  "groups": [
    {
      "id": 1,
      "name": "Basic Greetings"
    }
  ]
}
```

### GET /api/groups

Returns a list of groups.

#### JSON Response

```json
{
  "items": [
    {
      "id": 1,
      "name": "Basic Greetings",
      "word_count": 20
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "total_items": 10,
    "items_per_page": 100
  }
}
```

### GET /api/groups/:id

Returns details of a specific group.

#### JSON Response

```json
{
  "id": 1,
  "name": "Basic Greetings",
  "stats": {
    "total_word_count": 20
  }
}
```

### GET /api/groups/:id/words

Returns a list of words in a specific group.

#### JSON Response

```json
{
  "items": [
    {
      "french": "salut",
      "english": "hello",
      "correct_count": 5,
      "wrong_count": 2
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "total_items": 20,
    "items_per_page": 100
  }
}
```

### GET /api/groups/:id/sessions

Returns a list of sessions for a specific group.

#### JSON Response

```json
{
  "items": [
    {
      "id": 123,
      "activity_name": "Vocabulary Quiz",
      "group_name": "Basic Greetings",
      "start_time": "2025-02-08T17:20:23-05:00",
      "end_time": "2025-02-08T17:30:23-05:00",
      "review_words_count": 20
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "total_items": 5,
    "items_per_page": 100
  }
}
```

### GET /api/sessions

Returns a list of study sessions.

#### JSON Response

```json
{
  "items": [
    {
      "id": 123,
      "activity_name": "Vocabulary Quiz",
      "group_name": "Basic Greetings",
      "start_time": "2025-02-08T17:20:23-05:00",
      "end_time": "2025-02-08T17:30:23-05:00",
      "review_words_count": 20
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_items": 100,
    "items_per_page": 100
  }
}
```

### GET /api/sessions/:id

Returns details of a specific study session.

#### JSON Response

```json
{
  "id": 123,
  "activity_name": "Vocabulary Quiz",
  "group_name": "Basic Greetings",
  "start_time": "2025-02-08T17:20:23-05:00",
  "end_time": "2025-02-08T17:30:23-05:00",
  "review_words_count": 20
}
```

### GET /api/sessions/:id/words

Returns a list of words in a specific study session.

#### JSON Response

```json
{
  "items": [
    {
      "french": "salut",
      "english": "hello",
      "correct_count": 5,
      "wrong_count": 2
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "total_items": 20,
    "items_per_page": 100
  }
}
```

### POST /api/reset_history

Resets the study history.

#### JSON Response

```json
{
  "success": true,
  "message": "Study history has been reset"
}
```

### POST /api/full_reset

Performs a full system reset.

#### JSON Response

```json
{
  "success": true,
  "message": "System has been fully reset"
}
```

### POST /api/sessions/:id/words/:word_id/review

Submits a review for a word in a study session.

#### Request Params

- session_id integer
- word_id integer
- correct boolean

#### Request Payload

```json
{
  "correct": true
}
```

#### JSON Response

```json
{
  "success": true,
  "word_id": 1,
  "session_id": 123,
  "correct": true,
  "created_at": "2025-02-08T17:33:07-05:00"
}
```

## Task Runner Tasks

Lets list out possible tasks we need for our lang portal.

### Initialize Database

This task will initialize the sqlite database called `words.db`

### Migrate Database

This task will run a series of migrations sql files on the database.

Migrations live in the `migrations` folder.
The migration files will be run in order of their file name.
The file names should looks like this:

```sql
0001_init.sql
0002_create_words_table.sql
```

### Seed Data

This task will import json files and transform them into target data for our database.

All seed files live in the `seeds` folder.

In our task we should have DSL to specify each seed file and its expected group word name.

```json
[
  {
    "french": "payer",
    "english": "to pay"
  }
  ...
]
```
