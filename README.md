# go-cli-todo-app

A simple command-line todo application built with Go. Manage your tasks directly from the terminal with support for adding, listing, completing, and editing tasks.

## Features

- Add new tasks
- List all tasks with their status and timestamps
- Mark tasks as complete
- Edit existing task titles
- Persistent storage via a local JSON file

## Requirements

- Go 1.25+

## Installation

```bash
git clone https://github.com/mohammed-bata/go-cli-todo-app.git
cd go-cli-todo-app
go build -o todo .
```

## Usage

### Add a task

```bash
./todo -add "Buy groceries"
```

### List all tasks

```bash
./todo -list
```

Output example:

```
ID   TITLE          STATUS   Created AT                  COMPLETED AT
--   -----          ------   ----------                  ------------
1    Buy groceries  ❌       Sat,28 Feb 2026 05:20:18
2    Read a book    ✅       Sat,28 Feb 2026 06:30:00    Sat,28 Feb 2026 06:34:58
```

### Complete a task

```bash
./todo -complete 1
```

### Edit a task

```bash
./todo -edit "1:New task title"
```

## Data Storage

Tasks are stored locally in a `tasks.json` file in the same directory as the binary.

## Project Structure

```
go-cli-todo-app/
├── main.go       # CLI flag parsing and command routing
├── todo.go       # Task struct and file operations (Read, Write, Complete, Edit)
├── tasks.json    # Local task storage
├── go.mod
└── README.md
```

## License

MIT
