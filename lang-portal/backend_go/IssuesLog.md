# Issues Log

## 1. 'gpls' Error

### Explanation

Needed to upgrade Go.

### Solution
	1. Remove old version
		- sudo rm -rf /usr/local/go
	2. Download latest version of go
		- go to website to get the link to the latest version of go
		- cd ~/
		- wget wget https://golang.org/dl/go1.24.0.linux-amd64.tar.gz
	3. Extract the downloaded archive
		- sudo tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz
	4. Update the PATH environment variable:
		- nano ~/.profile or bashrc (open file in nano)
		- export PATH=$PATH:/usr/local/go/bin (add this line)
		- save and exit nano
	5. Reload the file
		- source ~/.profile or bashrc
	6. Check go's version
		- go version
	
## 2. compiler (BrokenImport) Error

### Explanation

Could not import ./database (no required module provides package "./database") 

### Solution
	Change module's name to backend_go
		- Open go.mod
		- Change the name from lang-portal to backend_go
		- save and close
		
## 3. Command 'mage' not found Error

### Explanation
While trying to run mage init.
The path was not setup correctly.

### Solution
	Add the correct directory to your PATH variable:
		1. Verify the installation directory:
			- ls $HOME/go/bin/mage or $GOPATH/bin/mage (defaults)
		2. nano ~/.bashrc
			- export PATH=$PATH:$HOME/go/bin/mage
			- save and exit
		3. reload bashrc
			- source ~/.bashrc
		4. Verify Mage Installation
			- mage -version
			
## 4. No .go files marked with the mage build tag in this directory error

### Explanation
After fixed the correct directoy in PATH and trying to rerun mage init.
Copilot had created a tasks directory and stored the mageFile.go in it while it should be saved in the root directory of the project.

### Solution
	Move MageFile.go to the root of the project (backend_go)

## 5. Error: exec: "sqlite3": executable file not found in $PATH

### Solution
	Install Sqlite3
		- sudo apt-get update
		- sudo apt-get install sqlite3
		- sqlite3 --version

## 6. Error: in prepare, near "<": syntax error (1) Error: exit status 1

### Explanation

The error you're encountering is due to the incorrect usage of the exec.Command function with the redirection operator (<). The redirection should be handled by the shell, not by the exec.Command function directly.

### Solution
    Change the function to execute from the shell:
        - cmd := exec.Command("sh", "-c", "sqlite3 words.db < migrations/0001_init.sql")

## 7. Problems> github.com/magefile/mage is not used in this module.

### Explanation
The error "github.com/magefile/mage is not used in this module" typically occurs when a dependency is listed in the go.mod file but is not actually imported in any of the Go source files in your project. This can happen if the import statement for Mage is missing or if the file containing the import statement is not being recognized correctly.

To resolve this issue, ensure that the Magefile.go file is correctly importing the Mage package and that the file is in the root directory of your Go module.

### Solution

    Ensure the Magefile.go contains the necessary import statements for Mage:
        - "github.com/magefile/mage/mg"
        - "github.com/magefile/mage/sh"

## 7. When saving file the import statements disappear.

### Explanation
The import statements for Mage are disappearing because they are not being used directly in the code. Go's goimports tool, which is often run automatically by editors like VS Code, removes unused imports when you save the file.

### Solution
    To prevent goimports from removing the Mage import statements, you need to use them in your code.
        - In Seed() add:
            - sh.Run("echo", "Seeding database...");
        - In Migrate() add:
            - mg.Deps(Init)

## 8. Running main.go: cannot open db Error

### Explanation
The name and directory of the database were not correct.

### Solution
	- Moved the database file into the database folder.
	- Revised the MageFile.go and main.go files with correct database name and path.