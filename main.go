package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/go-git/go-git/v5"
)

// TODO: add option to open expanded dif by appending ?`expand=1` at the end of the URL
func main() {
	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalf("Failed to open git repository: %v", err)
	}

	remoteName := "origin"
	remote, err := repo.Remote(remoteName)
	if err != nil {
		log.Fatalf("Failed to get remote '%s': %v", remoteName, err)
	}

	remoteURL := remote.Config().URLs[0]
	githubURL, err := parseGithubURL(remoteURL)
	if err != nil {
		log.Fatalf("Failed to parse GitHub URL: %v", err)
	}

	headRef, err := repo.Head()
	if err != nil {
		log.Fatalf("Failed to get HEAD reference: %v", err)
	}
	currentBranch := headRef.Name().Short()

	newPRURL := fmt.Sprintf("%s/compare/%s", githubURL, url.PathEscape(currentBranch))

	fmt.Printf("Opening: %s\n", newPRURL)
	openBrowser(newPRURL)
}

func parseGithubURL(remoteURL string) (string, error) {
	if strings.HasPrefix(remoteURL, "git@") {
		// SSH URL format: git@github.com:owner/repo.git
		parts := strings.SplitN(remoteURL, ":", 2)
		if len(parts) < 2 {
			return "", fmt.Errorf("invalid SSH remote URL format")
		}
		repoPath := strings.TrimSuffix(parts[1], ".git")
		return fmt.Sprintf("https://github.com/%s", repoPath), nil
	} else if strings.HasPrefix(remoteURL, "http://") || strings.HasPrefix(remoteURL, "https://") {
		// HTTP/HTTPS URL format: https://github.com/owner/repo.git
		trimmedURL := strings.TrimSuffix(remoteURL, ".git")
		return trimmedURL, nil
	} else {
		return "", fmt.Errorf("unsupported remote URL format")
	}
}

var execCommand = exec.Command

func openBrowser(url string) error {
	var cmd *exec.Cmd

	switch os := os.Getenv("OS"); os {
	case "windows":
		cmd = execCommand("cmd", "/c", "start", url)
	case "darwin":
		cmd = execCommand("open", url)
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = execCommand("xdg-open", url)
	}

	return cmd.Run()
}
