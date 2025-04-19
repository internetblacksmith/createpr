package main

import (
	"os/exec"
	"reflect"
	"testing"
)

func TestParseGithubURL(t *testing.T) {
	testCases := []struct {
		name           string
		inputURL       string
		expectedURL    string
		expectError    bool
		expectedErrMsg string
	}{
		{
			name:           "Valid HTTPS URL",
			inputURL:       "https://github.com/owner/repo.git",
			expectedURL:    "https://github.com/owner/repo",
			expectError:    false,
			expectedErrMsg: "",
		},
		{
			name:           "Valid SSH URL",
			inputURL:       "git@github.com:owner/repo.git",
			expectedURL:    "https://github.com/owner/repo",
			expectError:    false,
			expectedErrMsg: "",
		},
		{
			name:           "Invalid URL",
			inputURL:       "invalid-url",
			expectedURL:    "",
			expectError:    true,
			expectedErrMsg: "unsupported remote URL format",
		},
		{
			name:           "HTTPS URL without .git",
			inputURL:       "https://github.com/owner/repo",
			expectedURL:    "https://github.com/owner/repo",
			expectError:    false,
			expectedErrMsg: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualURL, err := parseGithubURL(tc.inputURL)
			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error for %s, but got none", tc.inputURL)
				}
				if err.Error() != tc.expectedErrMsg {
					t.Errorf("Expected error message %q, but got %q", tc.expectedErrMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for %s: %v", tc.inputURL, err)
				}
				if actualURL != tc.expectedURL {
					t.Errorf("For %s, expected URL %q, but got %q", tc.inputURL, tc.expectedURL, actualURL)
				}
			}
		})
	}
}

// Mock command execution
type MockCommand struct {
	cmd  string
	args []string
}

func TestOpenBrowser(t *testing.T) {
	// Save original function and restore after test
	originalExecCommand := execCommand
	originalGetOS := getOS
	defer func() {
		execCommand = originalExecCommand
		getOS = originalGetOS
	}()

	testCases := []struct {
		name         string
		goos         string
		url          string
		expectedCmd  string
		expectedArgs []string
	}{
		{
			name:         "Windows",
			goos:         "windows",
			url:          "https://github.com/owner/repo/compare/branch",
			expectedCmd:  "cmd",
			expectedArgs: []string{"/c", "start", "https://github.com/owner/repo/compare/branch"},
		},
		{
			name:         "macOS",
			goos:         "darwin",
			url:          "https://github.com/owner/repo/compare/branch",
			expectedCmd:  "open",
			expectedArgs: []string{"https://github.com/owner/repo/compare/branch"},
		},
		{
			name:         "Linux",
			goos:         "linux",
			url:          "https://github.com/owner/repo/compare/branch",
			expectedCmd:  "xdg-open",
			expectedArgs: []string{"https://github.com/owner/repo/compare/branch"},
		},
		{
			name:         "Default OS",
			goos:         "something-else",
			url:          "https://github.com/owner/repo/compare/branch",
			expectedCmd:  "xdg-open",
			expectedArgs: []string{"https://github.com/owner/repo/compare/branch"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Mock out command execution
			var mock MockCommand
			execCommand = func(command string, args ...string) *exec.Cmd {
				mock.cmd = command
				mock.args = args
				// Return a dummy command that succeeds
				return exec.Command("echo", "success")
			}

			// Set OS environment
			getOS = func() string {
				return tc.goos
			}

			// Call the function we're testing
			err := openBrowser(tc.url)
			if err != nil {
				t.Errorf("openBrowser returned unexpected error: %v", err)
			}

			// Verify the correct command was used
			if mock.cmd != tc.expectedCmd {
				t.Errorf("Expected command %q, got %q", tc.expectedCmd, mock.cmd)
			}

			// Verify the arguments
			if !reflect.DeepEqual(mock.args, tc.expectedArgs) {
				t.Errorf("Expected args %v, got %v", tc.expectedArgs, mock.args)
			}
		})
	}
}
