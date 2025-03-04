# Frontend Building with Copilot

### Prompt

I need to build a UI/frontend for my language learning portal web-app. It should be saved under the directory @frontend_react. You will find all of the technical requirement specifications in the @Frontend_Technical_Specs.md file. First, I would like you to give me a summary of what you plan to do. Once I agree to the plan, we will walk through every point in your plan, step by step. You will explain each step first then when I agree you can go ahead and execute it.

### Plan Summary

#### 1. Project Setup

- Initialize a new React project using Vite.js.
- Configure TypeScript support.
- Install necessary dependencies: React Router, Tailwind CSS, ShadCN components.

#### 2. Project Structure

- Create a directory structure to organize components, pages, and styles.

#### 3. Routing

- Set up React Router with the specified routes.
- Implement route redirection from / to /dashboard.

#### 4. Global Components

- Create global components such as Navigation, Breadcrumbs, and Pagination.

#### 5. Pages Implementation

- Implement each page as specified in the document:
```text
- Dashboard Page
- Study Activities Page
- Study Activity Page
- Study Activity Launch Page
- Words Page
- Word Page
- Word Groups Page
- Group Page
- Study Sessions Page
- Study Session Page
- Settings Page
```
#### 6. API Integration

- Set up API endpoints for each page to fetch and display data.

#### 7. Styling

- Use Tailwind CSS for styling components and pages.
- Ensure consistent design and layout across the application.

#### 8. Testing

- Write unit tests for components and pages.
- Ensure all functionalities work as expected.

# Detailed Steps

## Step 1: Project Setup

### Explanation

#### 1. Initialize a new React project using Vite.js:

- Vite.js is a fast build tool that provides a great development experience for modern web projects.
- We will use Vite to create a new React project.

#### 2. Configure TypeScript support:

- TypeScript provides static typing, which helps catch errors early and improves code quality.
- We will configure the project to use TypeScript.

#### 3. Install necessary dependencies:

- React Router: For handling routing in the application.
- Tailwind CSS: For styling the components and pages.
- ShadCN components: For reusable UI components.

### Execution
```sh
# Navigate to the desired directory
cd /mnt/c/users/rosie/OneDrive/Desktop/My\ Workspace/MyRepo/free_genai_bootcamp_2025/lang-portal

# Initialize a new React project using Vite.js with TypeScript support
npm create vite@latest frontend_react --template react-ts

# Navigate to the project directory
cd frontend_react

# Install necessary dependencies
npm install react-router-dom tailwindcss postcss autoprefixer @shadcn/ui

# Initialize Tailwind CSS
npx tailwindcss init -p
```

## Step 2: Project Structure

### Explanation

We need to create a directory structure to organize our components, pages, and styles. This will help keep our project organized and maintainable. Here is the proposed directory structure:
```text
frontend_react/
├── public/
├── src/
│   ├── assets/
│   ├── components/
│   ├── pages/
│   ├── styles/
│   ├── App.tsx
│   ├── index.tsx
│   └── ...
├── tailwind.config.cjs
├── postcss.config.cjs
├── tsconfig.json
└── package.json
```

- `public`: Contains static assets like images, fonts, etc.
- `assets`: Contains project-specific assets.
- `src/components/`: Contains reusable React components.
- `src/pages/`: Contains page components for different routes.
- `src/styles/`: Contains global styles and Tailwind CSS configurations.
- `App.tsx`: The main application component.
- `src/index.tsx`: The entry point of the application.

#### Execution
```sh
# Navigate to the project directory
cd /mnt/c/users/rosie/OneDrive/Desktop/My\ Workspace/MyRepo/free_genai_bootcamp_2025/lang-portal/frontend_react

# Create the directory structure
mkdir -p src/assets src/components src/pages src/styles

# Create placeholder files
touch src/App.tsx src/index.tsx
```
## Step 3: Routing
### Explanation
We need to set up React Router to handle the routing in our application. This involves:

1. Installing React Router.
2. Creating the necessary routes as specified in the technical requirements.
3. Implementing route redirection from `/` to `/dashboard`.

### Execution
1. Install React Router
```sh
npm install react-router-dom
```

2. Set up routing in App.tsx:

- Create the necessary routes (check App.tsx for content).
- Implement route redirection from `/` to `/dashboard`.

## Step 4: Global Components
### Explanation
We need to create global components such as Navigation, Breadcrumbs, and Pagination. These components will be used across multiple pages in our application to ensure a consistent user experience.

### Execution
1. Create the Navigation component:
- This component will contain links to the main sections of the application.
2. Create the Breadcrumbs component:
- This component will display the current location within the application.
3. Create the Pagination component:
- This component will handle pagination for lists of items.

## Step 5: Pages Implementation
### Explanation
We need to implement each page as specified in the technical requirements. This involves creating the components for each page and setting up the necessary API integrations.

### Execution
Let's start by creating the components for each page. We have already created placeholder components for the pages in the previous steps. Now, we will implement the details for each page.

## Step 6: API Integration

### Explanation
We need to set up API endpoints for each page to fetch and display data. This involves making HTTP requests to the backend API and handling the responses to update the state of our components.

### Execution
1. Install Axios:
- Axios is a popular library for making HTTP requests. We will use it to interact with our backend API.
```sh
npm install axios
```
2. Create an API Service:
- Create a new file api.ts in the src directory to centralize our API calls.

3. Update Components to Use API Service:
- Update each component to use the API service for fetching data.

## Step 7: Styling

### Explanation
We need to style our components and pages to ensure a consistent design and layout across the application. We will use Tailwind CSS for styling, which provides utility-first CSS classes to build custom designs quickly.

### Execution
1. Configure Tailwind CSS:
- Ensure that Tailwind CSS is properly configured in your project.
```sh
npx tailwindcss init -p
```
- Update the tailwind.config.cjs file to include the paths to all of your template files.
- Update the src/styles/index.css file to include the Tailwind CSS directives.

2. Apply Tailwind CSS Classes:
- Apply Tailwind CSS classes to your components and pages to style them.

# Missed Steps

## Install ShadCN
1. Install
```sh
npm install @shadcn/ui @shadcn/tailwind
```
2. Configure Tailwind CSS
Next, you need to configure Tailwind CSS to work with shadCN. Update your tailwind.config.cjs file to include the shadCN plugin and theme.
```sh
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
    "./node_modules/@shadcn/ui/dist/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@shadcn/ui'),
  ],
}
```

3. Update CSS File
Ensure that your CSS file includes the Tailwind CSS directives and shadCN styles. Update your `src/styles/index.css` file:
```sh
@import "@shadcn/ui/dist/styles.css";
```
4. Use ShadCN Components
Now you can start using shadCN components in your project.

ex: 
```sh
import { Button } from '@shadcn/ui';
<Button></Button>
```