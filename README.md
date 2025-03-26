# GO_IOT_Notes
# Go Program Setup and Execution in VS Code

## Step 1: Create a New Go Project or File

1. Open **VS Code** and navigate to the folder where you want to create your Go program.
2. Create a new file with a `.go` extension (e.g., `main.go`).
3. Write your Go program in this file. For example:

   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, World!")
   }
Step 2: Run Your Go Program
There are a couple of ways to run your Go program in VS Code.

Method 1: Using the Terminal
Open the integrated terminal in VS Code (`Ctrl + `` or via the menu Terminal > New Terminal).

Navigate to the directory where your Go file is located (if not already there).

Run the Go program using the command:

go run main.go

This will compile and run your Go program. If everything is set up correctly, you should see the output of your program in the terminal, for example:
output : 

Hello, World!


