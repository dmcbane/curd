---
layout: default
title: Command Reference - CURD
---

# Command Reference

[← Back to Home](index.html)

## Overview

CURD provides commands for saving, retrieving, and managing directory bookmarks. All commands are run using the `curd` binary, while navigation uses the `curr` shell function.

## Basic Syntax

```bash
curd [COMMAND] [KEYWORD] [OPTIONS]
curr [KEYWORD]  # For navigation (requires shell integration)
```

## Commands

### save

Save a directory with a keyword for quick access.

**Syntax:**
```bash
curd save [KEYWORD] [--dir <directory>] [--config <file>] [--verbose]
```

**Options:**
- `KEYWORD` - Name to associate with the directory (default: "default")
- `--dir <directory>` - Directory to save (default: current directory)
- `--config <file>` - Use alternate config file
- `--verbose` - Show detailed output

**Examples:**
```bash
# Save current directory as default
curd save

# Save current directory with keyword
cd ~/projects/website
curd save web

# Save specific directory
curd save docs --dir=~/Documents

# Use verbose output
curd save myproject --verbose
```

### list (ls)

List all saved directories and their keywords.

**Syntax:**
```bash
curd list [-k | --keywords-only] [--config <file>] [--verbose]
curd ls   # Alias for list
```

**Options:**
- `-k, --keywords-only` - Show only keywords, not paths
- `--config <file>` - Use alternate config file
- `--verbose` - Show detailed output

**Examples:**
```bash
# List all saved paths
curd list

# Output:
# default - /home/user
# docs - /home/user/Documents
# project - /home/user/projects/current

# List keywords only
curd ls -k

# Output:
# docs  project
```

### remove (rm)

Remove a saved directory from the configuration.

**Syntax:**
```bash
curd remove [KEYWORD] [--config <file>] [--verbose]
curd rm [KEYWORD]  # Alias for remove
```

**Options:**
- `KEYWORD` - Keyword of the path to remove (default: "default")
- `--config <file>` - Use alternate config file
- `--verbose` - Show detailed output

**Examples:**
```bash
# Remove a specific keyword
curd remove oldproject

# Remove default keyword
curd rm default

# Remove with verbose output
curd remove test --verbose
```

### clean

Remove all saved paths that no longer exist in the filesystem.

**Syntax:**
```bash
curd clean [--config <file>] [--verbose]
```

**Options:**
- `--config <file>` - Use alternate config file
- `--verbose` - Show detailed output

**Examples:**
```bash
# Clean up non-existent paths
curd clean

# Clean with verbose output to see what was removed
curd clean --verbose
```

### Read (Default Command)

Retrieve the path associated with a keyword. This is the default command when no command is specified.

**Syntax:**
```bash
curd [KEYWORD] [--config <file>] [--verbose]
```

**Options:**
- `KEYWORD` - Keyword to look up (default: "default")
- `--config <file>` - Use alternate config file
- `--verbose` - Show detailed output

**Examples:**
```bash
# Get default path
curd

# Get specific path
curd project

# Output: /home/user/projects/current
```

### help

Display help information.

**Syntax:**
```bash
curd help
curd --help
curd -h
```

### version

Display version information.

**Syntax:**
```bash
curd version
curd --version
curd -V
```

### completion (comp)

Helper command for shell completion. Used internally by bash completion scripts.

**Syntax:**
```bash
curd completion CMDLINE ...
curd comp CMDLINE ...  # Alias
```

This command is used by shell completion scripts and typically not called directly by users.

### completions

Generate a shell completion script for `curd` and `curr`.

**Syntax:**
```bash
curd completions [<shell>]
```

**Options:**
- `<shell>` - The shell to generate completion for: `bash`, `fish`, or `zsh`. If omitted, the shell is detected from the `SHELL` environment variable.

The script is written to standard output, so redirect or source it as appropriate for your shell.

**Examples:**
```bash
# Bash — source on shell startup
source <(curd completions bash)

# Zsh — install onto the fpath
curd completions zsh > ~/.zsh/completions/_curd

# Fish — install into fish's completions directory
curd completions fish > ~/.config/fish/completions/curd.fish

# Auto-detect the current shell from $SHELL
curd completions
```

## Global Options

These options can be used with most commands:

- `--config=<file>` - Use a custom configuration file instead of the default `~/.curdrc`
- `--verbose` or `-v` - Display extra information during command execution
- `--help` or `-h` - Show help information
- `--version` or `-V` - Show version information

## Navigation with curr

The `curr` function (set up during shell integration) is used for actual directory navigation:

```bash
# Navigate to default directory
curr

# Navigate to a saved directory
curr project

# curr is equivalent to:
# cd $(curd project)
```

## Configuration File

By default, CURD stores saved paths in `~/.curdrc` in YAML format:

```yaml
default: /home/user
docs: /home/user/Documents
project: /home/user/projects/current
downloads: /home/user/Downloads
```

The file uses restrictive permissions (0600) for security.

## Exit Codes

- `0` - Success
- `1` - Error (command failed, path not found, etc.)

## Examples Workflow

```bash
# Initial setup
cd ~/projects/backend
curd save be

cd ~/projects/frontend
curd save fe

cd ~/Documents
curd save docs

# Daily usage
curr be    # Jump to backend
curr fe    # Jump to frontend
curr docs  # Jump to documents

# Maintenance
curd list           # See all saved paths
curd clean          # Remove non-existent paths
curd remove old     # Remove specific path
```

---

[← Installation](installation.html) | [Home](index.html) | [FAQ →](faq.html)