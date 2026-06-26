package execute_test

import (
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/dmcbane/curd/v2/args"
	"github.com/dmcbane/curd/v2/config"
	"github.com/dmcbane/curd/v2/execute"
	"gopkg.in/yaml.v3"
)

// captureStdout runs fn while capturing everything written to os.Stdout and
// returns it as a string.
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	fn()

	w.Close()
	out, _ := io.ReadAll(r)
	return string(out)
}

// Helper function to create a temporary config file for testing
func createTempConfigFile(t *testing.T, content map[string]string) string {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "curd_config_test_*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer tmpfile.Close()

	marshaledContent, err := yaml.Marshal(content)
	if err != nil {
		t.Fatalf("Failed to marshal content for temp file: %v", err)
	}

	if _, err := tmpfile.Write(marshaledContent); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	return tmpfile.Name()
}

func TestExecuteCommand_Clean(t *testing.T) {
	tempDir := t.TempDir()
	existingFile := filepath.Join(tempDir, "existing_file.txt")
	nonExistingFile := filepath.Join(tempDir, "non_existing_file.txt")

	if err := os.WriteFile(existingFile, []byte("content"), 0644); err != nil {
		t.Fatalf("Failed to create existing file: %v", err)
	}

	initialPaths := map[string]string{
		"exist":    existingFile,
		"nonexist": nonExistingFile,
		"default":  "/some/default/path",
	}
	configFilePath := createTempConfigFile(t, initialPaths)
	defer os.Remove(configFilePath)

	mockArgs := args.Args{Clean: true, ConfigFile: configFilePath}
	cfg, err := config.NewConfig(configFilePath)
	if err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}
	err = execute.ExecuteCommand(mockArgs, *cfg)
	if err != nil {
		t.Errorf("ExecuteCommand(Clean) failed: %v", err)
	}

	// Re-read config to verify changes
	updatedCfg, err := config.NewConfig(configFilePath)
	if err != nil {
		t.Fatalf("Failed to re-read config after clean: %v", err)
	}

	expectedPaths := map[string]string{
		"exist":   existingFile,
		"default": "/some/default/path",
	}

	if len(updatedCfg.Paths) != len(expectedPaths) {
		t.Errorf("Expected %d paths after clean, got %d", len(expectedPaths), len(updatedCfg.Paths))
	}
	for k, v := range expectedPaths {
		if val, ok := updatedCfg.Paths[k]; !ok || val != v {
			t.Errorf("Path mismatch for %q: expected %q, got %q", k, v, val)
		}
	}
	if _, ok := updatedCfg.Paths["nonexist"]; ok {
		t.Errorf("nonexist path was not removed")
	}
}

func TestExecuteCommand_Completion(t *testing.T) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	mockArgs := args.Args{Completion: true, Cmdline: []string{"curd", "ls", "-"}, ConfigFile: "/tmp/nonexistent"}
	paths := map[string]string{"foo": "/path/to/foo", "bar": "/path/to/bar"}
	mockConfig := config.Config{Paths: paths}

	err := execute.ExecuteCommand(mockArgs, mockConfig)
	if err != nil {
		t.Errorf("ExecuteCommand(Completion) failed: %v", err)
	}

	w.Close()
	os.Stdout = oldStdout // Restore stdout
	out, _ := io.ReadAll(r)

	expectedOutput := "-h --help -V --version --config -v --verbose -k --keywords-only\n"
	if string(out) != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, string(out))
	}
}

func TestExecuteCommand_List(t *testing.T) {
	paths := map[string]string{
		"default": "/default/path",
		"proj1":   "/path/to/proj1",
		"proj2":   "/path/to/proj2",
	}
	configFilePath := createTempConfigFile(t, paths)
	defer os.Remove(configFilePath)

	cfg, _ := config.NewConfig(configFilePath)

	testCases := []struct {
		name           string
		args           args.Args
		expectedOutput string
	}{
		{
			name:           "list all (default)",
			args:           args.Args{List: true, ConfigFile: configFilePath, KeywordsOnly: false},
			expectedOutput: "default - /default/path\nproj1 - /path/to/proj1\nproj2 - /path/to/proj2\n",
		},
		{
			name:           "list keywords only",
			args:           args.Args{List: true, ConfigFile: configFilePath, KeywordsOnly: true},
			expectedOutput: "proj1  proj2\n", // Default keyword is excluded
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := execute.ExecuteCommand(tc.args, *cfg)
			if err != nil {
				t.Errorf("ExecuteCommand(List) failed: %v", err)
			}

			w.Close()
			os.Stdout = oldStdout
			out, _ := io.ReadAll(r)

			// Due to map iteration order, sort the output lines for comparison
			actualLines := strings.Split(strings.TrimSpace(string(out)), "\n")
			sort.Strings(actualLines)
			expectedLines := strings.Split(strings.TrimSpace(tc.expectedOutput), "\n")
			sort.Strings(expectedLines)

			if len(actualLines) != len(expectedLines) {
				t.Errorf("Expected %d lines, got %d. Actual: %q, Expected: %q", len(expectedLines), len(actualLines), string(out), tc.expectedOutput)
			} else {
				for i := range actualLines {
					if actualLines[i] != expectedLines[i] {
						t.Errorf("Line %d mismatch: Expected %q, Got %q. Actual: %q, Expected: %q", i, expectedLines[i], actualLines[i], string(out), tc.expectedOutput)
					}
				}
			}
		})
	}
}

func TestExecuteCommand_Remove(t *testing.T) {
	initialPaths := map[string]string{
		"key1":    "/path/to/key1",
		"key2":    "/path/to/key2",
		"default": "/default/path",
	}
	configFilePath := createTempConfigFile(t, initialPaths)
	defer os.Remove(configFilePath)

	mockArgs := args.Args{Remove: true, Keyword: "key1", ConfigFile: configFilePath}
	cfg, err := config.NewConfig(configFilePath)
	if err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	err = execute.ExecuteCommand(mockArgs, *cfg)
	if err != nil {
		t.Errorf("ExecuteCommand(Remove) failed: %v", err)
	}

	updatedCfg, err := config.NewConfig(configFilePath)
	if err != nil {
		t.Fatalf("Failed to re-read config after remove: %v", err)
	}

	expectedPaths := map[string]string{
		"key2":    "/path/to/key2",
		"default": "/default/path",
	}

	if len(updatedCfg.Paths) != len(expectedPaths) {
		t.Errorf("Expected %d paths after remove, got %d", len(expectedPaths), len(updatedCfg.Paths))
	}
	for k, v := range expectedPaths {
		if val, ok := updatedCfg.Paths[k]; !ok || val != v {
			t.Errorf("Path mismatch for %q: expected %q, got %q", k, v, val)
		}
	}
	if _, ok := updatedCfg.Paths["key1"]; ok {
		t.Errorf("key1 path was not removed")
	}
}

func TestExecuteCommand_Save(t *testing.T) {
	tempDir := t.TempDir()
	configFilePath := createTempConfigFile(t, map[string]string{})
	defer os.Remove(configFilePath)

	// Test saving current directory
	t.Run("save current directory", func(t *testing.T) {
		oldPwd, _ := os.Getwd()
		os.Chdir(tempDir) // Change to temp directory for testing
		defer os.Chdir(oldPwd)

		// Derive the expected path the same way ExecuteCommand does, via
		// os.Getwd(). On some platforms (e.g. macOS, where /var is a symlink
		// to /private/var) Getwd resolves symlinks, so comparing against the
		// raw tempDir would spuriously fail.
		expectedPath, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get working directory: %v", err)
		}

		mockArgs := args.Args{Save: true, Keyword: "testkey", ConfigFile: configFilePath, Directory: ""}
		cfg, _ := config.NewConfig(configFilePath)

		if err := execute.ExecuteCommand(mockArgs, *cfg); err != nil {
			t.Errorf("ExecuteCommand(Save current) failed: %v", err)
		}

		updatedCfg, _ := config.NewConfig(configFilePath)
		if val, ok := updatedCfg.Paths["testkey"]; !ok || val != expectedPath {
			t.Errorf("Expected %q, got %q for saved path", expectedPath, val)
		}
	})

	// Test saving a specified directory
	t.Run("save specified directory", func(t *testing.T) {
		testDir := filepath.Join(t.TempDir(), "sub_dir")
		if err := os.Mkdir(testDir, 0755); err != nil {
			t.Fatalf("Failed to create test directory: %v", err)
		}

		mockArgs := args.Args{Save: true, Keyword: "anotherkey", ConfigFile: configFilePath, Directory: testDir}
		cfg, _ := config.NewConfig(configFilePath)

		err := execute.ExecuteCommand(mockArgs, *cfg)
		if err != nil {
			t.Errorf("ExecuteCommand(Save specified) failed: %v", err)
		}

		updatedCfg, _ := config.NewConfig(configFilePath)
		absTestDir, _ := filepath.Abs(filepath.Clean(testDir))
		if val, ok := updatedCfg.Paths["anotherkey"]; !ok || val != absTestDir {
			t.Errorf("Expected %q, got %q for saved path", absTestDir, val)
		}
	})

	t.Run("save non-existent specified directory", func(t *testing.T) {
		nonExistentDir := filepath.Join(t.TempDir(), "nonexistent")
		mockArgs := args.Args{Save: true, Keyword: "badkey", ConfigFile: configFilePath, Directory: nonExistentDir}
		cfg, _ := config.NewConfig(configFilePath)

		err := execute.ExecuteCommand(mockArgs, *cfg)
		if err == nil {
			t.Error("Expected an error for saving non-existent directory, but got none")
		}
	})
}

func TestExecuteCommand_Read(t *testing.T) {
	paths := map[string]string{
		"default": "/default/path",
		"proj1":   "/path/to/proj1",
	}
	configFilePath := createTempConfigFile(t, paths)
	defer os.Remove(configFilePath)

	cfg, _ := config.NewConfig(configFilePath)

	testCases := []struct {
		name           string
		args           args.Args
		expectedOutput string
		expectError    bool
	}{
		{
			name:           "read existing keyword",
			args:           args.Args{Read: true, Keyword: "proj1", ConfigFile: configFilePath},
			expectedOutput: "/path/to/proj1\n",
			expectError:    false,
		},
		{
			name:           "read non-existent keyword",
			args:           args.Args{Read: true, Keyword: "nonexistent", ConfigFile: configFilePath},
			expectedOutput: "",
			expectError:    true,
		},
		{
			name:           "read default keyword",
			args:           args.Args{Read: true, Keyword: "default", ConfigFile: configFilePath},
			expectedOutput: "/default/path\n",
			expectError:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := execute.ExecuteCommand(tc.args, *cfg)

			w.Close()
			os.Stdout = oldStdout
			out, _ := io.ReadAll(r)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if string(out) != tc.expectedOutput {
					t.Errorf("Expected output %q, got %q", tc.expectedOutput, string(out))
				}
			}
		})
	}
}

func TestGenerateCompletions(t *testing.T) {
	testCases := []struct {
		shell    string
		contains []string
	}{
		{
			shell: "bash",
			contains: []string{
				"complete -F _completions_curd curd",
				"complete -F _completions_curr curr",
				"curd completion --",
				"curd ls -k",
			},
		},
		{
			shell: "zsh",
			contains: []string{
				"#compdef curd curr",
				"compdef _curd_complete curd",
				"compdef _curr_complete curr",
				"curd completion --",
				"curd ls -k",
			},
		},
		{
			shell: "fish",
			contains: []string{
				"complete -c curd",
				"complete -c curr",
				"curd completion --",
				"curd ls -k",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.shell, func(t *testing.T) {
			var genErr error
			out := captureStdout(t, func() {
				genErr = execute.GenerateCompletions(tc.shell)
			})
			if genErr != nil {
				t.Fatalf("GenerateCompletions(%q) returned error: %v", tc.shell, genErr)
			}
			for _, want := range tc.contains {
				if !strings.Contains(out, want) {
					t.Errorf("GenerateCompletions(%q) output missing %q.\nGot:\n%s", tc.shell, want, out)
				}
			}
		})
	}
}

func TestGenerateCompletions_Unsupported(t *testing.T) {
	out := captureStdout(t, func() {
		err := execute.GenerateCompletions("powershell")
		if err == nil {
			t.Error("Expected an error for unsupported shell, but got none")
		}
	})
	if out != "" {
		t.Errorf("Expected no output for unsupported shell, got %q", out)
	}
}

func TestGenerateCompletions_DetectFromEnv(t *testing.T) {
	t.Setenv("SHELL", "/usr/bin/fish")
	var genErr error
	out := captureStdout(t, func() {
		genErr = execute.GenerateCompletions("")
	})
	if genErr != nil {
		t.Fatalf("GenerateCompletions(\"\") returned error: %v", genErr)
	}
	if !strings.Contains(out, "complete -c curd") {
		t.Errorf("Expected fish completion when SHELL=/usr/bin/fish, got:\n%s", out)
	}
}

func TestGenerateCompletions_DetectUnset(t *testing.T) {
	t.Setenv("SHELL", "")
	out := captureStdout(t, func() {
		err := execute.GenerateCompletions("")
		if err == nil {
			t.Error("Expected an error when shell cannot be detected, but got none")
		}
	})
	if out != "" {
		t.Errorf("Expected no output when shell cannot be detected, got %q", out)
	}
}

func TestExecuteCommand_GenerateCompletions(t *testing.T) {
	var execErr error
	out := captureStdout(t, func() {
		execErr = execute.ExecuteCommand(
			args.Args{GenerateCompletions: true, Shell: "bash"},
			config.Config{},
		)
	})
	if execErr != nil {
		t.Fatalf("ExecuteCommand(GenerateCompletions) failed: %v", execErr)
	}
	if !strings.Contains(out, "complete -F _completions_curd curd") {
		t.Errorf("Expected bash completion output, got:\n%s", out)
	}
}

func TestBashCompletionHelper(t *testing.T) {
	paths := map[string]string{"foo": "/path/to/foo", "bar": "/path/to/bar", "default": "/default/path"}

	testCases := []struct {
		name           string
		cmdline        []string
		expectedOutput string
	}{
		{
			name:           "completion for -- option",
			cmdline:        []string{"curd", "--"},
			expectedOutput: "--help --version --config --verbose\n",
		},
		{
			name:           "completion for - option",
			cmdline:        []string{"curd", "-"},
			expectedOutput: "-h --help -V --version --config -v --verbose\n",
		},
		{
			name:           "completion for command",
			cmdline:        []string{"curd", ""},
			expectedOutput: "-h --help -V --version --config -v --verbose clean ls list save rm remove bar foo\n",
		},
		{
			name:           "completion for ls",
			cmdline:        []string{"curd", "ls", ""},
			expectedOutput: "-h --help -V --version --config -v --verbose -k --keywords-only\n",
		},
		{
			name:           "completion for rm",
			cmdline:        []string{"curd", "rm", ""},
			expectedOutput: "-h --help -V --version --config -v --verbose bar foo\n",
		},
		{
			name:           "completion for save",
			cmdline:        []string{"curd", "save", ""},
			expectedOutput: "-h --help -V --version --config -v --verbose --dir\n",
		},
		{
			name:           "completion for save --dir",
			cmdline:        []string{"curd", "save", "--dir", ""},
			expectedOutput: "-h --help -V --version --config -v --verbose\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			execute.BashCompletionHelper(tc.cmdline, paths)

			w.Close()
			os.Stdout = oldStdout
			out, _ := io.ReadAll(r)

			if string(out) != tc.expectedOutput {
				t.Errorf("Expected output %q, got %q", tc.expectedOutput, string(out))
			}
		})
	}
}
