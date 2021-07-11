package main

import (
	"os/exec"
	"strings"
)

// List apt packages on the system
func listAptPackages() ([]string, error) {
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

var cachedAptPackages []string = nil

// Checks that an apt dependency is present.
func TestAptDependency(tgt Dependency) (bool, []string) {
	var packages []string = nil

	// retrieve & cache apt packages
	if cachedAptPackages == nil {
		packages, err := listAptPackages()
		cachedAptPackages = packages

		if err != nil {
			return false, []string{"failed to list apt packages"}
		}
	}

	cachedPackages = packages

	for _, name := range packages {
		if name == tgt["name"] {
			return true, []string{"snap package \"" + tgt["name"] + "\" installed"}
		}
	}

	return false, []string{tgt["name"] + " not in listed apt-packages"}
}
