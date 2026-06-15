---
layout: default
title: FAQ - CURD
---

# Frequently Asked Questions

[← Back to Home](index.html)

## General Questions

### What is CURD?

CURD (Change to one of a User's Recurrent Directories) is a command-line tool that helps you navigate to frequently-used directories quickly using keywords. Instead of typing long paths or using multiple `cd` commands, you can jump directly to any saved directory with a simple keyword.

### How is CURD different from cd?

While `cd` requires you to type the full or relative path each time, CURD lets you save directories with short keywords and jump to them from anywhere in your filesystem. Think of it as bookmarks for your terminal.

### How is CURD different from autojump/z/fasd?

Those tools automatically track your directory usage and use frecency (frequency + recency) algorithms to guess where you want to go. CURD takes a different approach:
- **Explicit control** - You decide which directories to save
- **Predictable** - Keywords always go to the same place
- **Simple** - No learning algorithms or databases
- **Lightweight** - Just a simple YAML config file

### Is CURD compatible with my shell?

CURD works with:
- Bash
- Zsh
- Fish
- PowerShell
- Windows Command Prompt

Any shell that can run external commands and change directories can work with CURD.

## Installation Issues

### Why do I get "command not found: curd"?

The `curd` binary is not in your PATH. Check where it was installed:
- If installed with `go install`: Check `~/go/bin`
- Add the directory to your PATH in your shell configuration file

### Why do I get "command not found: curr"?

You haven't set up the shell integration. The `curr` function needs to be added to your shell configuration. See the [installation guide](installation.html) for your specific shell.

### Why do I get Go version errors when building?

This is usually due to conflicting GOROOT settings. Use the provided build script:
```bash
./build.sh
```

Or manually unset GOROOT:
```bash
unset GOROOT && go build
```

## Usage Questions

### Can I have multiple keywords for the same directory?

Yes! You can save the same directory with different keywords:
```bash
cd ~/projects/important
curd save important
curd save imp
curd save i
```

All three keywords will point to the same directory.

### Can I use CURD with relative paths?

When you save a directory, CURD automatically converts it to an absolute path. This ensures the keyword always works regardless of your current location.

### What happens if I save a keyword that already exists?

The new path will overwrite the old one. CURD will not warn you about overwrites.

### Can I use spaces in keywords?

No, keywords cannot contain spaces. Use underscores or hyphens instead:
```bash
curd save my_project
curd save my-project
```

### How do I update a saved path?

Simply save it again with the same keyword:
```bash
cd ~/new/location
curd save myproject  # Updates existing 'myproject' keyword
```

### What is the "default" keyword?

When you don't specify a keyword, CURD uses "default". This is useful for your most frequently accessed directory:
```bash
cd ~/workspace
curd save          # Saves as "default"
curr               # Goes to default directory
```

## Security Questions

### Why does v2.0.0 change file permissions?

CURD v2.0.0 restricts config file permissions to 0600 (user read/write only) to prevent other users on the system from seeing your saved directories. This is a security best practice.

### Why are paths with ".." rejected?

Directory traversal sequences like `..` can be security risks. CURD v2.0.0 validates paths to prevent potential security issues. Use absolute paths or paths without `..` sequences.

### Is my configuration file encrypted?

No, the configuration file is plain text YAML. However, it's protected with user-only permissions (0600) so other users cannot read it.

## Troubleshooting

### CURD says my keyword doesn't exist

Check if the keyword is saved:
```bash
curd list
```

Keywords are case-sensitive. "Project" and "project" are different keywords.

### The clean command removed paths I wanted to keep

The `clean` command removes paths that don't exist in the filesystem. This might happen if:
- Network drives are temporarily unavailable
- External drives are not mounted
- Directories were temporarily moved

Always review what will be cleaned with `curd list` before running `curd clean`.

### My config file seems corrupted

The config file is simple YAML. You can edit it manually:
```bash
cat ~/.curdrc
# Edit if needed
vim ~/.curdrc
```

Or start fresh:
```bash
mv ~/.curdrc ~/.curdrc.backup
# Re-save your directories
```

## Advanced Usage

### Can I use multiple config files?

Yes, use the `--config` option:
```bash
# Work directories
curd save project --config ~/.curdrc-work

# Personal directories
curd save photos --config ~/.curdrc-personal

# Use them
curr project  # Uses default config
curd project --config ~/.curdrc-work  # Uses work config
```

### Can I share config files between machines?

Yes, the config file is portable. You can:
- Sync it with Dropbox/Google Drive
- Store it in a git repository
- Copy it between machines

Just remember that paths must exist on each machine.

### Can I use CURD in scripts?

Yes! CURD works well in scripts:
```bash
#!/bin/bash
PROJECT_DIR=$(curd myproject)
cd "$PROJECT_DIR"
# Do something in the project directory
```

### Can I use environment variables in paths?

No, CURD stores absolute paths. Environment variables are expanded when you save:
```bash
curd save home --dir=$HOME  # Saves as /home/username
```

## Contributing

### How can I contribute to CURD?

- Report bugs and request features on [GitHub Issues](https://github.com/dmcbane/curd/issues)
- Submit pull requests with improvements
- Update documentation
- Share CURD with others who might find it useful

### Where can I get help?

- Check this FAQ
- Read the [documentation](index.html)
- Open an issue on [GitHub](https://github.com/dmcbane/curd/issues)

---

[← Commands](commands.html) | [Home](index.html)