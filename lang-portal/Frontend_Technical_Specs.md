# Frontend Technical Specs

## Dashboard Page `/dashboard`

### Purpose

The purpose of this page is to provide a summary of the user's learning experience
and act as the default page when a user visits the app.

### Components

- Last Study Session
  - shows last activity practiced
  - shows date of the last activitiy
  - shows correct count vs wrong count from last activity
  - shows a link to the last practiced group of words
- Study Progress
  - shows total number of words studied from all groups out of all available words in our database (ex. 3/124)
  - shows a mastery progress bar showing the above number's percentage (ex. 10%)
- Quick Stats
  - shows the success rate by calculating the percentage of the average rate of correct words across all study sessions (ex. 80%)
  - shows the total number of study sessions (ex. 4)
  - shows the total number of active groups (ex. 3)
  - shows the study streak, which is calculated by the number of consecutive days the user has practiced counting retrospective from today (ex. 4 days)
- `Start Studying` Button
  - shows a button linking to the study activities page

### API Endpoints

- GET /api/dashboard/last_study_session
- GET /api/dashboard/study_progress
- GET /api/dashboard/quick_stats

## Study Activities Page `/activities`

### Purpose

The purpose of this page is to display a collection of study activities. Each activity shown in a card.

### Components

- Study Activity Card
  - shows a thumbnail of the study activity
  - shows the name of the study activity
  - shows a `Launch` button to take us to the launch page
  - shows a `View` button to take us to the study activity page

### API Endpoints

- GET /api/activities

## Study Activity Page `/activities/:id`

### Purpose

The purpose of this page is to show the details of a study activity with all of its past study sessions.

### Components

- shows the name of the study activity
- shows a thumbnail of the study activity
- shows a description of the study activity
- shows a `Launch` button
- shows a study activities paginated list with the following columns:
  - id
  - activity name
  - group name
  - start time
  - end time (inferred by the last review_word submitted)
  - number of review words

### API Endpoints

- GET /api/activities/:id
- GET /api/activities/:id/sessions

## Study Activity Launch Page `/activities/:id/launch`

### Purpose

The purpose of this page is to launch a study activity.

### Components

- Shows the name of the study activity
- shows a launch form having:
  - a combo box showing all groups allowing user to select a group of words to study
  - a `Launch Now` button to allow user to launch the activity

### Behaviour

- After the form is submitted:
  - a new tab opens with the study activity based on its URL provided in the database.
  - the page will redirect to the study session page

### API Endpoints

- POST /api/activities

## Words Page `/words`

### Purpose

The purpose of this page is to display all words available in our database.

### Components

- shows a paginated Word List with the columns:
  - Word in French
  - Word in English
  - Correct Count
  - Wrong Count
- Pagination will display 100 items per page
- Clicking the French word will take us to the its corresponding word page

### API Endpoints

- GET /api/words

## Word Page `/words/:id`

### Purpose

The purpose of this page is to show information about a specific word.

### Components

- Word in French
- Word in English
- Study Statistics
  - Correct Count
  - Wrong Count
- Word Groups this word belongs to
  - display as a series of pills like tags
  - when group name is clicked it will take us to this group's page

### API Endpoints

- GET /api/words/:id

## Word Groups Page `/groups`

### Purpose

The purpose of this page is to show a list of groups in our database.

### Components

- shows a paginated group list with columns:
  - Group Name
  - Word Count
- clicking the group name will take us to this group's page

### API Endpoints

- GET /api/groups

## Group `/groups/:id`

### Purpose

The purpose of this page is to show information about a specific group.

### Components

- shows the group name
- shows the total words count in this group
- shows the list of words in this group (Paginated List of Words)
  - Should use the same component as the words page
- shows all study sessions of this group(Paginated List of Study Sessions)
  - Should use the same component as the study sessions page

### API Endpoints

- GET /api/groups/:id
- GET /api/groups/:id/words
- GET /api/groups/:id/sessions

## Study Sessions Page `/sessions`

### Purpose

The purpose of this page is to show a list of study sessions in our database.

### Components

- show a paginated study session list with columns:
  - Id
  - Activity Name
  - Group Name
  - Start Time
  - End Time
  - Number of Review Words
- clicking the study session id will take us to the study session page

### API Endpoints

- GET /api/sessions

## Study Session Page `/sessions/:id`

### Purpose

The purpose of this page is to show information about a specific study session.

### Components

- shows the details of this study session:
  - Activity Name
  - Group Name
  - Start Time
  - End Time
  - Number of Review Words
- shows a list of all review words (Paginated List of Words)
  - Should use the same component as the words page

### API Endpoints

- GET /api/sessions/:id
- GET /api/sessions/:id/words

## Settings Page `/settings`

### Purpose

The purpose of this page is to modify the configurations of the study portal.

### Components

- shows a theme selection option (ex. Light, Dark, System Default)
- shows a `Reset History` Button
  - this button will delete all study sessions and review words
- shows a `Full Reset` Button
  - this button will drop all tables and re-create with seed data

### API Endpoints

- POST /api/reset_history
- POST /api/full_reset
