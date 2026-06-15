package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dmcbane/curd/v2/config"
	"gopkg.in/yaml.v3"
)

// Helper function to create a temporary config file for testing
func createTempConfigFile(t *testing.T, content string) string {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "curd_config_test_*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer tmpfile.Close()

	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	return tmpfile.Name()
}

func TestNewConfig_ReadConfig(t *testing.T) {
	testCases := []struct {
		name          string
		configFile    string
		fileContent   string
		expectedPaths map[string]string
		expectError   bool
	}{
		{
			name:        "valid config file",
			fileContent: "key1: value1\nkey2: value2\n",
			expectedPaths: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			expectError: false,
		},
		{
			name:          "empty config file",
			fileContent:   "",
			expectedPaths: map[string]string{},
			expectError:   false,
		},
		{
			name:          "non-existent config file",
			configFile:    filepath.Join(os.TempDir(), "non_existent.yaml"),
			expectedPaths: map[string]string{}, // Should return empty map, no error
			expectError:   false,
		},
		{
			name:        "invalid yaml",
			fileContent: "key1: value1\nkey2: value2\n  bad-indent: value\n", // Malformed YAML (indentation error)
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var configFilePath string
			if tc.configFile != "" {
				configFilePath = tc.configFile
			} else {
				configFilePath = createTempConfigFile(t, tc.fileContent)
				defer os.Remove(configFilePath)
			}

			cfg, err := config.NewConfig(configFilePath)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if cfg == nil {
				t.Fatalf("Config object is nil")
			}

			if cfg.ConfigFile != configFilePath {
				t.Errorf("Expected ConfigFile %q, got %q", configFilePath, cfg.ConfigFile)
			}

			if len(cfg.Paths) != len(tc.expectedPaths) {
				t.Errorf("Expected %d paths, got %d", len(tc.expectedPaths), len(cfg.Paths))
			}

			for k, v := range tc.expectedPaths {
				if val, ok := cfg.Paths[k]; !ok || val != v {
					t.Errorf("Expected path %q:%q, got %q:%q", k, v, k, val)
				}
			}
		})
	}
}

func TestConfig_WriteConfig(t *testing.T) {
	testCases := []struct {
		name            string
		initialPaths    map[string]string
		expectedContent string
		expectError     bool
	}{
		{
			name: "write multiple paths",
			initialPaths: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			expectedContent: "key1: value1\nkey2: value2\n", // Order might vary, need to handle
			expectError:     false,
		},
		{
			name:            "write empty paths",
			initialPaths:    map[string]string{},
			expectedContent: "null\n", // YAML marshals empty map to null
			expectError:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tmpfile, err := os.CreateTemp("", "curd_write_config_test_*.yaml")
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			configFilePath := tmpfile.Name()
			tmpfile.Close() // Close to allow WriteFile to work
			defer os.Remove(configFilePath)

			cfg := &config.Config{
				ConfigFile: configFilePath,
				Paths:      tc.initialPaths,
			}

			err = cfg.WriteConfig()

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			content, err := os.ReadFile(configFilePath)
			if err != nil {
				t.Fatalf("Failed to read written config file: %v", err)
			}

			// For map, order is not guaranteed. We should unmarshal and compare maps.
			var actualPaths map[string]string
			err = yaml.Unmarshal(content, &actualPaths)
			if err != nil {
				t.Fatalf("Failed to unmarshal written content: %v", err)
			}

			if len(actualPaths) != len(tc.initialPaths) {
				t.Errorf("Expected %d paths after writing, got %d", len(tc.initialPaths), len(actualPaths))
			}

			for k, v := range tc.initialPaths {
				if val, ok := actualPaths[k]; !ok || val != v {
					t.Errorf("Expected path %q:%q after writing, got %q:%q", k, v, k, val)
				}
			}
		})
	}
}

func TestUnmarshalNonExistentPath(t *testing.T) {
	content := []byte(`
nonexistent_key: /path/that/does/not/exist
`)
	var paths map[string]string
	err := yaml.Unmarshal(content, &paths)

	if err != nil {
		t.Fatalf("yaml.Unmarshal returned an unexpected error for non-existent path: %v", err)
	}

	expectedPath := "/path/that/does/not/exist"
	if val, ok := paths["nonexistent_key"]; !ok || val != expectedPath {
		t.Errorf("Expected 'nonexistent_key' to be unmarshaled with value %q, got %q (ok: %t)", expectedPath, val, ok)
	}
}
