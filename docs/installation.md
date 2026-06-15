---
layout: default
title: Installation Guide - CURD
---

# Installation Guide

[← Back to Home](index.html)

## Prerequisites

CURD requires either:
- Go 1.18 or later (for building from source)
- Or download a pre-built binary

## Installation Methods

### Method 1: Using Go Install (Recommended)

If you have Go installed, this is the easiest method:

```bash
go install github.com/dmcbane/curd@latest
```

This will install the `curd` binary to `$GOPATH/bin` (usually `~/go/bin`). Make sure this directory is in your PATH.

### Method 2: Download Pre-built Binary

1. Visit the [releases page](https://github.com/dmcbane/curd/releases)
2. Download the appropriate binary for your operating system
3. Extract the binary to a directory in your PATH (e.g., `/usr/local/bin` or `~/bin`)
4. Make it executable (on Unix-like systems):
   ```bash
   chmod +x curd
   ```

### Method 3: Build from Source

Clone the repository and build:

```bash
git clone https://github.com/dmcbane/curd.git
cd curd
./build.sh  # or: go build
```

The binary will be created in the current directory. Move it to a location in your PATH.

## Shell Integration Setup

After installing the `curd` binary, you need to set up shell integration to use the `curr` command for directory navigation.

### Bash

Add to your `~/.bashrc`:

```bash
function curr() {
  D=$(curd "$@")
  cd "${D}"
}

# Optional: Enable bash completion
if [ -f ~/go/src/github.com/dmcbane/curd/curd_completion.bash ]; then
  source ~/go/src/github.com/dmcbane/curd/curd_completion.bash
fi
```

Then reload your configuration:
```bash
source ~/.bashrc
```

### Zsh

Add to your `~/.zshrc`:

```bash
function curr() {
  D=$(curd "$@")
  cd "${D}"
}

# Optional: Enable bash completion in zsh
if [ -f ~/go/src/github.com/dmcbane/curd/curd_completion.bash ]; then
  autoload bashcompinit
  bashcompinit
  source ~/go/src/github.com/dmcbane/curd/curd_completion.bash
fi
```

Then reload:
```bash
source ~/.zshrc
```

### Fish

Create `~/.config/fish/functions/curr.fish`:

```fish
function curr
    set -l D (curd $argv)
    cd "$D"
end
```

### PowerShell

Add to your PowerShell profile (run `$profile` to find its location):

```powershell
Function Get-Curd-Directory {
  [CmdletBinding()]
  Param($arg)
  $content = if ($arg) {curd $arg} Else {curd}
  Set-Location "$content"
}
Set-Alias curr Get-Curd-Directory -Description "Change to a curd directory"
```

Then reload:
```powershell
. $profile
```

### Windows Command Prompt

Create a batch file named `curr.bat` in a directory that's in your PATH:

```batch
@echo off
curd %* > %TEMP%\vv.tmp
set /p VV=<%TEMP%\vv.tmp
cd /D "%VV%"
```

## Verify Installation

After installation and shell setup:

1. **Check the curd binary:**
   ```bash
   curd --version
   ```
   Should output: `Curd 2.0.0` (or current version)

2. **Test saving a directory:**
   ```bash
   cd /tmp
   curd save temp
   ```

3. **Test navigation:**
   ```bash
   cd ~
   curr temp  # Should change to /tmp
   ```

## Troubleshooting

### "command not found: curd"

- Ensure the curd binary is in your PATH
- Check installation location: `which curd` or `where curd` (Windows)
- Add the installation directory to PATH if needed

### "command not found: curr"

- Ensure you've added the shell function to your shell configuration
- Reload your shell configuration or start a new terminal session

### Go Build Errors

If you encounter version mismatch errors when building:

```bash
# Use the provided build script
./build.sh

# Or manually unset GOROOT
unset GOROOT && go build
```

### Permission Denied

On Unix-like systems, ensure the binary is executable:
```bash
chmod +x /path/to/curd
```

## Upgrading from v1.x

If upgrading from CURD v1.x to v2.0.0:

1. Your config file permissions will be automatically updated to 0600
2. Any paths containing `..` will be rejected and need to be re-saved with absolute paths
3. Back up your config file before upgrading:
   ```bash
   cp ~/.curdrc ~/.curdrc.backup
   ```

## Uninstallation

To uninstall CURD:

1. Remove the binary:
   ```bash
   rm $(which curd)
   ```

2. Remove the shell function from your shell configuration file

3. Optionally, remove the config file:
   ```bash
   rm ~/.curdrc
   ```

---

[← Back to Home](index.html) | [Configuration →](configuration.html)