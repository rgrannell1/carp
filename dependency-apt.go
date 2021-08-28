package main

import (
	"os/exec"
	"strings"
)

// List apt packages on the system
func ListAptPackages() ([]string, error) {
	cmd := exec.Command("apt", "list")
	stdout, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	output := strings.Split(string(stdout), "\n")
	installed := []string{}

	for ith, val := range output {
		if ith > 0 {
			parts := strings.Split(val, "/")
			installed = append(installed, parts[0])
		}
	}

	return installed, nil
}

// -- broken cache

// Checks that an apt dependency is present.
func TestAptDependency(facts *SystemFacts, tgt Dependency) (bool, []string) {
	for _, name := range facts.AptPackages {
		if name == tgt["name"] {
			return true, []string{"snap package \"" + tgt["name"] + "\" installed"}
		}
	}

	return false, []string{tgt["name"] + " not in listed apt-packages"}
}
