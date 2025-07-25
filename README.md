<h2>Go Project Template</h2> <br>
This repository serves as a starting point for a Go (Golang) project, demonstrating the standard setup using Go Modules.

<h2>About Go</h2> <br>
Go is a statically and strongly typed, compiled programming language known for its: <br>

Fast compilation times <br>
Built-in support for concurrency  <br>
Simplicity and efficiency <br>
Project Structure <br>
<h2>Go code is organized into:</h2>  <br>

Packages : A collection of Go source files (.go files) in the same directory. <br>
Modules : A collection of related Go packages. A module is defined by a go.mod file at its root. <br>
<h2>Getting Started</h2> <br>
Prerequisites <br>
Go installed (version 1.16 or later is recommended for module support). <br>
Initializing a Go Module <br>
To create a new Go project, you first initialize a Go module. This is done using the go mod init command. <br>
go mod init github.com/yourusername/your-repo-name <br>

<h2>Why this specific path?</h2> <br>

Using the path github.com/yourusername/your-repo-name is a widely adopted convention because: <br>

Consistency : It aligns with the URL of your remote Git repository (e.g., on GitHub).<br>
Discoverability : It makes the module easily identifiable and importable by others if you publish it.<br>
Best Practice : It follows standard practices within the Go community for module naming.<br>
This command creates a go.mod file, which tracks dependencies for your project.<br>

Creating Source Files
Create your Go source code files with the .go extension (e.g., main.go).

Running the Code
There are two primary ways to execute your Go program:

1. go run :
Compiles and runs the specified Go program in one step.
The compiled binary is temporary and discarded after execution.
Ideal for quick development and testing.
go run main.go

2. go build :
Compiles the Go program and generates a standalone executable binary file (e.g., main on Linux/macOS or main.exe on Windows).
The binary is saved to disk and can be executed independently.
Used when you want to distribute or run the application later.<br>
# Compile the program
go build main.go

# Run the generated executable (Linux/macOS)
./main
# Run the generated executable (Windows)
 .\main.exe
