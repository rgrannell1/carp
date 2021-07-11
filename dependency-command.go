package main

import (
	"os/exec"
)

// Checks that a command is present on a system
func TestCommand(tgt Dependency) (bool, []string) {
	if tgt["name"] == "" {
		return false, []string{"group name not provided"}
	}

	cmd := exec.Command("which", tgt["name"])
	runErr := cmd.Run()

	if runErr != nil {
		exitError := runErr.(*exec.ExitError)

		if exitError != nil {
			code := exitError.ExitCode()

			if code != 0 {
				return false, []string{tgt["name"] + " not present"}
			}
		}
	}

	return true, []string{tgt["name"] + " present"}
}
