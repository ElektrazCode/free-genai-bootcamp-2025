# Frontend Technical Specs

## Business Goal

We are building a langauge learning web-app which serves the following purposes:

- A unified launchpad to launch different learning activities
- An inventory of possible vocabulary words grouped by themes that can be learned.
- A learning record store (LRS), providing correct and wrong scores on practice vocabulary to review study progress.
- The web-app is intended for desktop only, so we don't have to be concerned with mobile layouts.

## Technical Requirements

- React.js as the frontend library
- React router as the router library
- Tailwind CSS as the css framework
- Vite.js as the local development server
- Typescript for the programming language
- ShadCN for components
- There will be no authentication/authorization
- The app will serve a single user

## Frontend Routes
- This is a list of routes for our web-app we are building.
- Each of these routes will serve a seperate page.
- We'll describe each page in more details under the pages heading.
```text
  /dashboard
  /activties
  /activties/:id
  /words
  /words/:id
  /groups
  /groups/:id
  /sessions
  /settings
```

The default route `/` should forward to `/dashboard`

## Global Components
### Navigation
There will be a horizontal navigation bar with the following links:

Dashboard
Activities
Words
Word Groups
Sessions
Settings

### Breadcrumbs
Beneath the navigation there will be breadcrumbs so users can easily see where they are.
Examples of breadcrumbs:

Dashboard > Activities > Flashcards > Animals

### Pagination

- Some pages will contain long lists that will need pagination.
- Each pagination page will be limited to a maximum of 50 rows/page.
- The pagination will be shown as a bar at the bottom of the page/table.
- Previous button: grey out if you cannot go further back.
- Page 1 of 3: With the current page bolded.
- Next button: greyed out if you cannot go any further forward.
- All table headings should be sortable, clicking it will toggle between ascending and descending order. An ascii arrow should indicate direction: ascending order arrow will be pointing down and descending order arrow will be pointing up.

# Pages

## Dashboard Page

### Route
`/dashboard`

### Purpose

The purpose of this page is to provide a summary of the user's learning experience and progress.
and act as the default page when a user visits the app.

### Components
Each of the following components is displayed in a card:
  - Last Study Session
    - shows the name of the last activity practiced in the last session.
    - shows the date of the last session.
    - shows the total correct counts vs total wrong counts from the words reviewed in the last session.
    - shows a link to the group of words practiced in the last session. 
  - Study Progress
    - shows the total number of words studied from all groups out of all available words in our database (ex. 3/124).
    - shows a mastery progress bar showing the above number's percentage (ex. 10%).
  - Quick Stats
    - shows the success rate by calculating the percentage of the average rate of correct words across all study sessions (ex. 80%).
    - shows the total number of study sessions (ex. 4).
    - shows the total number of active groups (groups that have been studied in study sessions) (ex. 3).
    - shows the study streak, which is calculated by the number of consecutive days the user has practiced counting retrospective from today (ex. 4 days).

- `Start Studying` Button
  - shows a button linking to the study activities page.

### API Endpoints

- GET /api/dashboard/last_study_session
- GET /api/dashboard/study_progress
- GET /api/dashboard/quick_stats

## Study Activities Page

### Route
 `/activities`

### Purpose

The purpose of this page is to display a collection of study activities. Each activity shown in a card.

### Components

- Study Activity Card:
  - shows a thumbnail of the study activity.
  - shows the name of the study activity.
  - shows a `Launch` button to take us to the `Study Activity Launch Page`.
  - shows a `View` button to take us to the study activity page.

### API Endpoints

- GET /api/activities

## Study Activity Page

### Route
`/activities/:id`

### Purpose

The purpose of this page is to show the details of a study activity with all of its past study sessions.

### Components

- shows the name of the study activity
- shows a thumbnail of the study activity
- shows a description of the study activity
- shows a `Launch` button
- shows all study sessions realted to this activity in a paginated list with the following columns:
  - Group Name: This will be a link to the Group Page
  - Start Time: When the session was created in YYY-MM-DD HH:MM format (12 hours)
  - End Time: When the session ended in YYY-MM-DD HH:MM format (12 hours)
  - Review Words: Total count of review words in this session.

### API Endpoints

- GET /api/activities/:id
- GET /api/activities/:id/sessions

## Study Activity Launch Page

### Route
`/activities/:id/launch`

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

## Words Page

### Route
`/words`

### Purpose

The purpose of this page is to display all words available in our database.

### Components

- shows a paginated Word List with the columns:
  - French (This is a link that will take us to its corresponding `Word Page`)
  - English
  - Correct Count
  - Wrong Count
- Pagination will display 50 rows (words) per page.


### API Endpoints

- GET /api/words

## Word Page

### Route
`/words/:id`

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

## Word Groups Page

### Route
`/groups`

### Purpose

The purpose of this page is to show a list of groups in our database.

### Components

- shows a paginated group list with columns:
  - Group Name
  - Word Count (Total count of words associated with this group)
- clicking the group name will take us to this group's page

### API Endpoints

- GET /api/groups

## Group

### Route

`/groups/:id`

### Purpose

The purpose of this page is to show information about a specific group.

### Components

- shows the group name
- shows the total word count associated with this group
- shows the list of words in this group (Paginated List of Words)
  - Should use the same component as the words page
- shows all study sessions of this group(Paginated List of Study Sessions)
  - Should use the same component as the study sessions page

### API Endpoints

- GET /api/groups/:id
- GET /api/groups/:id/words
- GET /api/groups/:id/sessions

## Study Sessions Page

### Route
`/sessions`

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

## Study Session Page

### Route
`/sessions/:id`

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

## Settings Page

### Route
`/settings`

### Purpose

The purpose of this page is to modify the configurations of the study portal.

### Components

- Dark Mode Toggle: This is a toggle that changes from light to dark theme.
- shows a `Reset History` Button
  - this button will delete all study sessions and review words
- shows a `Full Reset` Button
  - this button will drop all tables and re-create with seed data
  - We need to confirm this action in a dialog and type the word `reset me` to confirm.

### API Endpoints

- POST /api/reset_history
- POST /api/full_reset
