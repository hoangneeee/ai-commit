package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func getGitDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--staged")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to get git diff: %w", err)
	}

	if out.Len() == 0 {
		return "", fmt.Errorf("no staged changes found. Please stage your changes with 'git add'")
	}

	return out.String(), nil
}

func getGitRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("not a git repository: %w", err)
	}

	return out.String(), nil
}
