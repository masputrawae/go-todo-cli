# Todo CLI

A simple and efficient command-line interface for managing your todo tasks, built with Go.

## Features

- ✅ Create, read, update, and delete todo tasks
- 📋 List all tasks with their status
- 🔍 Filter tasks by status (completed/pending)
- 💾 Persistent storage of tasks
- ⚡ Fast and lightweight CLI tool

## Installation

### From Source

Make sure you have [Go](https://golang.org/doc/install) installed (version 1.16 or higher).

```bash
git clone https://github.com/masputrawae/todo-cli.git
cd todo-cli
go build -o todo-cli
```

Then move the binary to your PATH:

```bash
sudo mv todo-cli /usr/local/bin/
```

## Usage

### Add a new task

```bash
todo-cli add "Your task description"
```

### List all tasks

```bash
todo-cli list
```

### List completed tasks

```bash
todo-cli list --status completed
```

### List pending tasks

```bash
todo-cli list --status pending
```

### Mark a task as completed

```bash
todo-cli complete <task-id>
```

### Delete a task

```bash
todo-cli delete <task-id>
```

### Update a task

```bash
todo-cli update <task-id> "Updated description"
```

## Examples

```bash
# Add a new task
$ todo-cli add "Buy groceries"
Task created with ID: 1

# List all tasks
$ todo-cli list
ID | Status    | Description
1  | pending   | Buy groceries

# Mark task as completed
$ todo-cli complete 1

# List again
$ todo-cli list
ID | Status    | Description
1  | completed | Buy groceries
```

## Project Structure

```
todo-cli/
├── main.go
├── cmd/
│   ├── add.go
│   ├── list.go
│   ├── delete.go
│   ├── complete.go
│   └── update.go
├── models/
│   └── task.go
├── storage/
│   └── storage.go
├── go.mod
└── README.md
```

## Configuration

Tasks are stored in a local JSON file. By default, the storage location is `~/.todo-cli/tasks.json`.

You can customize the storage location by setting the `TODO_CLI_HOME` environment variable:

```bash
eport TODO_CLI_HOME=/custom/path
```

## Development

### Prerequisites

- Go 1.16 or higher
- Git

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o todo-cli
```

### Running

```bash
go run main.go list
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Masputra Wae**

## Changelog

### v1.0.0
- Initial release
- Basic CRUD operations for tasks
- Task filtering by status

## Support

If you encounter any issues or have suggestions, please open an [issue](https://github.com/masputrawae/todo-cli/issues) on GitHub.