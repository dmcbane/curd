# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

CURD (Change to one of a User's Recurrent Directories) is a Go-based command-line tool for quick directory navigation. It allows users to save frequently visited directories with keywords and quickly jump to them. The tool works across multiple platforms (Windows, macOS, Linux) and shell environments.

## Build and Development Commands

**Note**: If you encounter Go version conflicts, use the provided shell scripts or prefix commands with `unset GOROOT &&`.

```bash
# Build the application
./build.sh
# Or: unset GOROOT && go build

# Run all tests
./test.sh
# Or: unset GOROOT && go test ./...

# Install the application to GOPATH/bin
unset GOROOT && go install github.com/dmcbane/curd/v2@v2.0.1

# Run tests for a specific package
unset GOROOT && go test ./config
unset GOROOT && go test ./execute

# Run tests with verbose output
unset GOROOT && go test -v ./...

# Run a specific test
unset GOROOT && go test -run TestExecuteCommand_Save ./execute

# Clean up dependencies
unset GOROOT && go mod tidy
```

## Architecture

The codebase is organized into three main packages:

1. **args** (`args/args.go`): Command-line argument parsing using docopt-go library. Handles all CLI argument parsing and validation.

2. **config** (`config/config.go`): Configuration file management using YAML. Manages reading/writing the `.curdrc` configuration file that stores saved directory paths.

3. **execute** (`execute/execute.go`): Core command execution logic. Implements all CURD commands (clean, list, save, remove, read) and bash completion.

The main entry point (`main.go`) orchestrates these packages:
1. Parses arguments via `args.NewArgs()`
2. Loads configuration via `config.NewConfig()`
3. Executes commands via `execute.ExecuteCommand()`

## Key Implementation Details

- **Configuration Storage**: Paths are stored in a YAML file (default: `~/.curdrc`) as a map of keyword->path pairs
- **Default Keyword**: The "default" keyword is special and used when no keyword is specified
- **Shell Integration**: Since Go programs can't change the parent shell's directory, CURD outputs the path to STDOUT, and shell wrapper functions (`curr.*` files) handle the actual directory change
- **Bash Completion**: Implemented via `BashCompletionHelper` in execute package and `curd_completion.bash` script

## Testing

The project has comprehensive test coverage in:
- `config/config_test.go`: Tests configuration file operations
- `execute/execute_test.go`: Tests all command execution paths

Tests use temporary files and directories to avoid affecting the real filesystem.