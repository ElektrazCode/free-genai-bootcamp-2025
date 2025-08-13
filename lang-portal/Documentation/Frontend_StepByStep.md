# Frontend Step By Step

## Step 1: Project Setup

### Explanation
1. Initialize React project with Vite and TypeScript
2. Install and configure TailwindCSS and ShadCN
3. Set up project structure
4. Configure ESLint and Prettier

### Execution
#### 1. Initialize React project with Vite and TypeScript
```sh
# Create new Vite project
npm create vite@latest frontend_react -- --template react-ts

# Navigate to project directory
cd frontend_react

# Install dependencies
npm install

# Install additional core dependencies
npm install @tanstack/react-query
npm install react-router-dom
npm install class-variance-authority
npm install clsx
npm install tailwind-merge
```
#### 2. Install and configure TailwindCSS and ShadCN
1. TailwindCSS
```sh
npm install -D tailwindcss postcss autoprefixer
```
- Create the `tailwind.config.js` file.
- Create the `index.css` file.
- Udpdate the main `App.tsx` file.

2. ShadCN
```sh
cd frontend_react
npm install shadcn-ui
npx shadcn@latest init
```
When prompted during initialization, choose:

Would you like to use TypeScript (recommended)? › Yes
Which style would you like to use? › Default
Which color would you like to use as base color? › Slate
Where is your global CSS file? › src/index.css
Would you like to use CSS variables for colors? › Yes
Where is your tailwind.config.js located? › tailwind.config.js
Configure the import alias for components? › @/components
Configure the import alias for utils? › @/lib/utils
Are you using React Server Components? › No

Update Configuration Files:
- `tailwind.config.js`
- `index.css.js`

#### 3. Project Structure
```sh
mkdir -p src/{components,pages,services,hooks,types,utils}/{common,dashboard,activities,words,sessions}
```

