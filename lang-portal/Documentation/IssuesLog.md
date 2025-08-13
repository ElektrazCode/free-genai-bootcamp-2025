# Issues Log

## 1. ShadCN-UI deprecated
The 'shadcn-ui' package is deprecated. Please use the 'shadcn' package instead:

### Solution
```sh 
npx shadcn@latest init
```

## 2. Running `npx shadcn@latest init`
No Tailwind CSS configuration found.

### Solution
1. Initialize Tailwind CSS properly
```sh
npx tailwindcss init -p
```
2. Create/Update TypeScript Configuration
`tsconfig.json`

3. Update Vite Configuration
`vite.config.ts`

4. Initialize TailwindCSS Configuration
`tailwind.config.js`

5. Create Utils Directory and File
`utils.ts`
6. Run:
```sh
npx shadcn-ui@latest init
```
The initialization should work now with these settings:

TypeScript: Yes
Style: Default
Base color: Slate
Global CSS: src/index.css
CSS variables: Yes
tailwind.config.js location: tailwind.config.js
Components path: @/components
Utils path: @/lib/utils
React Server Components: No

## 3. Installing Node

Checking availability of unzip... Missing!
Not installing fnm due to missing dependencies.

#### Solution

```sh
# For Ubuntu/Debian based systems
sudo apt-get update
sudo apt-get install unzip
# Install fnm
curl -fsSL https://fnm.vercel.app/install | bash

# Reload shell configuration
source ~/.bashrc

# Install and use Node.js LTS version
fnm install --lts
fnm use lts-latest
# Check Node.js installation
node --version
npm --version
```

