---
layout: default
title: CURD - Fast Directory Navigation
---

# CURD - Change to one of a User's Recurrent Directories

[← Back to Home](index.html)

[![GitHub Release](https://img.shields.io/github/v/release/dmcbane/curd)](https://github.com/dmcbane/curd/releases)
[![License](https://img.shields.io/github/license/dmcbane/curd)](https://github.com/dmcbane/curd/blob/main/LICENSE)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-blue)](https://go.dev/)

**CURD** is a lightning-fast command-line tool that lets you jump to frequently used directories with simple keywords. Stop typing long paths - save your directories once and access them instantly from anywhere.

<div style="text-align: center; margin: 2em 0;">
  <img src="https://asciinema.org/a/curd-demo.svg" alt="CURD Demo" style="max-width: 100%;">
  <p><em>Jump between project directories instantly</em></p>
</div>

## ✨ Features

- 🚀 **Instant Navigation** - Jump to any saved directory with a single command
- 🔖 **Keyword Bookmarks** - Save directories with memorable keywords
- 🌍 **Cross-Platform** - Works on Windows, macOS, and Linux
- 🐚 **Multi-Shell Support** - Compatible with Bash, Zsh, Fish, PowerShell, and Command Prompt
- 🔒 **Secure** - Config files are protected with user-only permissions
- 🎯 **Tab Completion** - Generate bash, fish, and zsh completion for commands and keywords
- 🧹 **Auto-Cleanup** - Remove non-existent paths with the clean command

## 🚀 Quick Start

### Installation

Using Go:
```bash
go install github.com/dmcbane/curd/v2@v2.1.0
```

Or download the latest binary from the [releases page](https://github.com/dmcbane/curd/releases).

### Shell Integration

Add the appropriate function to your shell configuration:

#### Bash/Zsh (`~/.bashrc` or `~/.zshrc`)
```bash
function curr() {
  D=$(curd "$@")
  cd "${D}"
}
```

#### Fish (`~/.config/fish/functions/curr.fish`)
```fish
function curr
    set -l D (curd $argv)
    cd "$D"
end
```

#### PowerShell (`$profile`)
```powershell
Function Get-Curd-Directory {
  [CmdletBinding()]
    Param($arg)
      $content = if ($arg) {curd $arg} Else {curd}
      Set-Location "$content"
};Set-Alias curr Get-Curd-Directory
```

#### Windows Command Prompt
Save as `curr.bat` in your PATH:
```batch
@echo off
curd %* > %TEMP%\vv.tmp
set /p VV=<%TEMP%\vv.tmp
cd /D "%VV%"
```

## 📖 Usage Examples

### Save Directories
```bash
# Save current directory with a keyword
cd ~/projects/myapp
curd save myapp

# Save a specific directory
curd save docs --dir=~/Documents

# Save current directory as default (no keyword needed)
cd ~/workspace
curd save
```

### Navigate to Saved Directories
```bash
# Jump to a saved directory
curr myapp

# Jump to default directory
curr

# Jump to docs
curr docs
```

### Manage Saved Paths
```bash
# List all saved paths
curd list

# List only keywords
curd list -k

# Remove a saved path
curd remove oldproject

# Clean up non-existent paths
curd clean
```

## 🎯 Common Use Cases

### Development Workflow
```bash
# Set up your project directories
cd ~/src/frontend && curd save fe
cd ~/src/backend && curd save be
cd ~/src/database && curd save db

# Quick navigation during development
curr fe  # Jump to frontend
curr be  # Jump to backend
curr db  # Jump to database
```

### Document Management
```bash
# Organize frequently accessed folders
curd save docs --dir=~/Documents
curd save downloads --dir=~/Downloads
curd save projects --dir=~/Projects
curd save config --dir=~/.config
```

## 🔧 Advanced Features

### Shell Completion

Generate a completion script for bash, fish, or zsh and enable tab completion for commands and keywords. Omit the shell to auto-detect it from `$SHELL`:

```bash
# bash — add to ~/.bashrc
source <(curd completions bash)

# zsh — add to ~/.zshrc
source <(curd completions zsh)

# fish
curd completions fish > ~/.config/fish/completions/curd.fish
```

### Custom Configuration File

Use a different configuration file:
```bash
curd --config ~/my-curd-config.yaml list
```

### Verbose Output

See detailed information about operations:
```bash
curd save myproject --verbose
```

## 🔒 Security

CURD v2.0.0 includes important security improvements:

- **Protected Config Files** - Configuration files use 0600 permissions (user-only access)
- **Path Traversal Protection** - Prevents directory traversal attacks
- **Environment Validation** - Safely handles missing HOME/USERPROFILE variables

## 📚 Documentation

- [Installation Guide](installation.html)
- [Configuration](configuration.html)
- [Command Reference](commands.html)
- [Shell Integration](shells.html)
- [FAQ](faq.html)

## 🤝 Contributing

Contributions are welcome! Please feel free to submit issues, feature requests, or pull requests.

- [Report Issues](https://github.com/dmcbane/curd/issues)
- [View Source](https://github.com/dmcbane/curd)
- [Changelog](https://github.com/dmcbane/curd/blob/main/CHANGELOG.md)

## 📄 License

CURD is released under the MIT License. See the [LICENSE](https://github.com/dmcbane/curd/blob/main/LICENSE) file for details.

## 🙏 Acknowledgments

CURD is inspired by tools like [autojump](https://github.com/wting/autojump), [fasd](https://github.com/clvv/fasd), and [z](https://github.com/rupa/z), but focuses on simplicity and explicit control over your directory bookmarks.

[← Back to Home](index.html)

---

<div style="text-align: center; margin-top: 3em;">
  <p>Made with ❤️ by <a href="https://github.com/dmcbane">H. Dale McBane</a></p>
  <p>
    <a href="https://github.com/dmcbane/curd">GitHub</a> •
    <a href="https://github.com/dmcbane/curd/releases">Releases</a> •
    <a href="https://github.com/dmcbane/curd/issues">Issues</a>
  </p>
</div>