# Users-Posts Backend
This is the backend server for the Web Developer Assignment, built with Go, Gin and SQLite3.

## Prerequisites
- Go
- Gin
- SQLite3

## Installation
1. Install dependencies:

```bash
go mod tidy
```

2. Build the project:

```bash
go build ./src/main.go
```
This will compile the go project into an executable named main.

3. Running the Server

Start the server in production mode:

```bash
./main
```

## Project Structure
```angular2html
backend/
├── src/           # Go source files
├── config/        # Configuration files
└── ...
```

# Important Notes
Make sure the Go files are properly compiled into an executable before running the server
The production server runs from the compiled files in the project directory