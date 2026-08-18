# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Fixed
- Fish completion for `curd` itself. The generated fish script now completes file paths after `--config` and directories after `--dir` instead of offering nothing, and splits `curd ls -k` output with `string split -n` so empty elements are discarded.

### Changed
- The generated fish script carries an install header explaining that fish loads a completion file only when completing the command matching its file name, so it must be installed as both `curd.fish` and `curr.fish`. README, installation guide, and command reference document the two-name install, the `$SHELL` detection caveat, and how to verify with `complete -C`.

### Added
- `TestGenerateCompletions_FishIntegration` runs the generated script under a real fish shell in a throwaway `XDG_CONFIG_HOME` and asserts on `complete -C` output. Skipped when fish is not installed.

## [2.2.0] - 2026-06-26

### Added
- Tab completion of saved keywords for the `curr` wrapper, built directly into the example integration scripts. Sourcing `curr.sh` (bash), `curr.fish`, or `curr.ps1` (PowerShell) now both defines `curr` and registers keyword completion using `curd ls -k` as the source.
- New example integration scripts `curr.zsh` (zsh `compdef` completion) and `curr.csh` (csh/tcsh alias with `complete` builtin), giving every supported shell a self-contained `curr` wrapper with keyword completion.

### Changed
- Documentation (README, installation guide, and overview) updated to show the self-contained `curr` snippets with completion for bash, zsh, fish, PowerShell, and csh/tcsh.

## [2.1.0] - 2026-06-14

### Added
- `curd completions [<shell>]` command that generates a shell completion script for `bash`, `fish`, or `zsh`. When the shell argument is omitted it is detected from the `SHELL` environment variable. The generated scripts wire up completion for both `curd` and `curr` using the existing `curd completion` dynamic helper.

### Removed
- The static `curd_completion.bash` file. Bash completion is now produced on demand with `curd completions bash`, which keeps a single source of truth and adds fish/zsh support.

## [2.0.1] - 2026-06-14

### Fixed
- Bumped reported version to 2.0.1 to match the published release tag
- Republished the module so the Go proxy serves the corrected `/v2` content (including the CHANGELOG `/v2` migration note)

## [2.0.0] - 2024-06-14

### Changed (Breaking)
- **BREAKING**: Module path changed to `github.com/dmcbane/curd/v2` as required by Go for v2+ modules. Install with `go install github.com/dmcbane/curd/v2@latest` and import internal packages from the `/v2` path
- **BREAKING**: Config file permissions changed from 0644 to 0600 for improved security (user-only access)
- **BREAKING**: Path validation now rejects directory traversal attempts (paths containing "..")

### Added
- Path traversal protection to prevent security vulnerabilities
- Warning message when HOME/USERPROFILE environment variables are not set
- Fallback to current directory when home directory is undefined
- Security-focused test suite
- Build and test shell scripts to handle Go environment conflicts
- Improved error messages with context

### Fixed
- Error handling for `filepath.Abs` that was previously ignored
- List command edge case when no keys are defined
- Inefficient error creation pattern (replaced `errors.New(fmt.Sprintf())` with `fmt.Errorf()`)
- Non-idiomatic error handling in `config.NewConfig`
- Go module version format (was 1.22.2, now 1.18)
- Code formatting issues identified by gofmt
- Naming convention (default_config to defaultConfig)
- Unnecessary parentheses in map access

### Improved
- Refactored BashCompletionHelper function for better maintainability
- Extracted helper functions to reduce code complexity
- Better separation of concerns in completion logic

## [1.2.4] - 2019-10-14
- Previous version (see git history for details)