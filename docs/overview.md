---
layout: default
title: Overview - CURD
description: What CURD is, how it works, and the commands you'll use most.
---

# Overview

**CURD** (*Change to one of a User's Recurrent Directories*) is a small, fast
command-line tool for jumping to directories you visit often. Save a directory
under a keyword once, then return to it from anywhere with a short
`curr <keyword>` — no long paths, no fuzzy guessing.

[![GitHub Release](https://img.shields.io/github/v/release/dmcbane/curd)](https://github.com/dmcbane/curd/releases)
[![License](https://img.shields.io/github/license/dmcbane/curd)](https://github.com/dmcbane/curd/blob/main/LICENSE)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-blue)](https://go.dev/)

## How it works

A program can't change its parent shell's working directory — once `curd`
exits, any directory change it made is gone. CURD works around this by
**printing** the resolved path to standard output:

```bash
cd "$(curd)"        # change to the default directory
cd "$(curd go)"     # change to the directory saved as "go"
```

That is too much to type for everyday use, so the shell integration wraps it in
a `curr` function — plus tab completion for your saved keywords. After the
one-time setup, the two commands above become simply:

```bash
curr
curr go
```

See the [Installation Guide](installation.html) for the `curr` setup in Bash,
Zsh, Fish, PowerShell, and csh/tcsh.

## Commands at a glance

| Command | What it does |
| --- | --- |
| `curd save [keyword]` | Save the current directory (or `--dir <path>`) under a keyword |
| `curr [keyword]` | Jump to a saved directory (the shell wrapper) |
| `curd ls` / `curd ls -k` | List saved keyword→path pairs, or just the keywords |
| `curd rm <keyword>` | Remove a saved bookmark |
| `curd clean` | Drop bookmarks whose paths no longer exist |

The special `default` keyword is used whenever you omit one. Full details are in
the [Command Reference](commands.html).

## Useful options

```bash
# Generate command/keyword completion for your shell (bash, fish, zsh)
source <(curd completions bash)

# Use an alternate configuration file instead of ~/.curdrc
curd --config ~/work-curdrc ls

# Print extra detail about what curd is doing
curd save myproject --verbose
```

## Configuration & security

Bookmarks live in a YAML file (default `~/.curdrc`) as keyword→path pairs. The
file is created with user-only `0600` permissions, and CURD rejects paths
containing `..` to prevent directory-traversal surprises.

## Learn more

- [Installation Guide](installation.html) — per-shell setup with tab completion
- [Command Reference](commands.html) — every command and flag
- [FAQ](faq.html) — common questions
- [Changelog](https://github.com/dmcbane/curd/blob/main/CHANGELOG.md)

## Acknowledgments

CURD is inspired by [autojump](https://github.com/wting/autojump),
[fasd](https://github.com/clvv/fasd), and [z](https://github.com/rupa/z), but
focuses on simplicity and explicit control over your directory bookmarks.
