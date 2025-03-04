# Frontend Technical Specifications

## Technology Stack
- React 18+ with TypeScript
- Vite for build tooling
- TailwindCSS for styling
- ShadcnUI for component library
- React Router for navigation
- React Query for API data fetching
- Vitest for testing

## Application Structure

### Core Features
1. Dashboard
   - Last study session summary
   - Study progress overview
   - Quick stats display
   
2. Activities
   - List available learning activities
   - Activity details and launch interface
   - Session history per activity
   
3. Words Management
   - Browse vocabulary list
   - Word details with stats
   - Group-based word organization
   
4. Study Sessions
   - Active session interface
   - Session history
   - Performance review

### Directory Structure
```
frontend_react/
├── src/
│   ├── components/
│   │   ├── common/
│   │   ├── dashboard/
│   │   ├── activities/
│   │   ├── words/
│   │   └── sessions/
│   ├── pages/
│   ├── services/
│   ├── hooks/
│   ├── types/
│   └── utils/
├── public/
└── tests/
```

## Component Pages

### Dashboard Page
- Display last study session
- Show progress statistics
- Present quick stats
- Navigation to main features

### Activities Page
- List all available activities
- Activity details view
- Launch activity interface
- Session history per activity

### Words Page
- Vocabulary list with search and filters
- Word details view
- Group management interface
- Study progress per word

### Sessions Page
- Active session interface
- Historical sessions list
- Session performance review
- Export functionality

## Data Management
- React Query for API integration
- Local storage for session persistence
- TypeScript interfaces matching API responses
- Error boundary implementation

## User Interface
- Responsive design
- Dark/Light mode support
- Accessibility compliance
- Loading states and error handling

## Development Guidelines
- Component-based architecture
- TypeScript for type safety
- Unit testing with Vitest
- ESLint and Prettier configuration
