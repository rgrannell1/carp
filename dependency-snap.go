package main

import (
	"os/exec"
	"strings"
)

// List all installed snap packages
func ListSnapPackages() ([]string, error) {
	cmd := exec.Command("snap", "list")
	stdout, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	output := strings.Split(string(stdout), "\n")
	installed := []string{}

	for ith, val := range output {
		// remove headers and append
		if ith > 0 {
			parts := strings.Split(val, " ")
			installed = append(installed, parts[0])
		}
	}

	return installed, nil
}

// Test whether a snap package is installed on this system.
func TestSnapDependency(facts *SystemFacts, tgt Dependency) (bool, []string) {
	for _, name := range facts.SnapPackages {
		if name == tgt["name"] {
			return true, []string{"snap package \"" + tgt["name"] + "\" installed"}
		}
	}

	return false, []string{tgt["name"] + " not in listed snap-packages"}
}
