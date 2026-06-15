package execute_test

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dmcbane/curd/v2/args"
	"github.com/dmcbane/curd/v2/config"
	"github.com/dmcbane/curd/v2/execute"
)

func TestPathTraversalPrevention(t *testing.T) {
	configFilePath := createTempConfigFile(t, map[string]string{})
	defer os.Remove(configFilePath)

	testCases := []struct {
		name          string
		directory     string
		expectError   bool
		errorContains string
	}{
		{
			name:          "reject path with ..",
			directory:     "../etc/passwd",
			expectError:   true,
			errorContains: "path traversal not allowed",
		},
		{
			name:          "reject complex traversal",
			directory:     "/home/user/../../etc/passwd",
			expectError:   true,
			errorContains: "path traversal not allowed",
		},
		{
			name:          "allow normal relative path",
			directory:     "./subdir",
			expectError:   true, // Will fail as directory doesn't exist
			errorContains: "does not exist",
		},
		{
			name:        "allow absolute path",
			directory:   t.TempDir(),
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockArgs := args.Args{
				Save:       true,
				Keyword:    "testkey",
				ConfigFile: configFilePath,
				Directory:  tc.directory,
			}
			cfg, _ := config.NewConfig(configFilePath)

			err := execute.ExecuteCommand(mockArgs, *cfg)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else if !strings.Contains(err.Error(), tc.errorContains) {
					t.Errorf("Expected error containing %q, got %q", tc.errorContains, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestConfigFilePermissions(t *testing.T) {
	// Create a temporary config file
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "test.curdrc")

	cfg := &config.Config{
		ConfigFile: configPath,
		Paths: map[string]string{
			"test": "/test/path",
		},
	}

	// Write the config
	err := cfg.WriteConfig()
	if err != nil {
		t.Fatalf("Failed to write config: %v", err)
	}

	// Check file permissions
	info, err := os.Stat(configPath)
	if err != nil {
		t.Fatalf("Failed to stat config file: %v", err)
	}

	mode := info.Mode().Perm()
	expectedMode := os.FileMode(0600)

	if mode != expectedMode {
		t.Errorf("Config file has wrong permissions. Expected %o, got %o", expectedMode, mode)
	}
}

func TestEmptyKeysListCommand(t *testing.T) {
	// Test with empty paths
	configFilePath := createTempConfigFile(t, map[string]string{})
	defer os.Remove(configFilePath)

	mockArgs := args.Args{
		List:         true,
		KeywordsOnly: true,
		ConfigFile:   configFilePath,
	}
	cfg, _ := config.NewConfig(configFilePath)

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := execute.ExecuteCommand(mockArgs, *cfg)
	if err != nil {
		t.Errorf("ExecuteCommand failed: %v", err)
	}

	w.Close()
	os.Stdout = oldStdout
	out, _ := io.ReadAll(r)

	// Should output empty line, not crash
	if string(out) != "\n" {
		t.Errorf("Expected empty output with newline, got %q", string(out))
	}
}

func TestHomeDirectoryHandling(t *testing.T) {
	// This test would need to mock environment variables
	// which is tricky in Go, so we'll just document the behavior
	t.Skip("Environment variable mocking not implemented")
}