package main

import (
	"os/exec"
	"strings"
)

func Indent(content string, count int) string {
	lines := strings.Split(content, "\n")
	padded := make([]string, len(lines))

	for ith, line := range lines {
		padded[ith] = strings.Repeat(" ", count) + line
	}

	return strings.Join(padded, "\n")
}

func IndentList(lines []string, count int) []string {
	padded := make([]string, len(lines))

	for ith, line := range lines {
		padded[ith] = strings.Repeat(" ", count) + line
	}

	return padded
}

func ListSnapPackages() ([]string, error) {
	cmd := exec.Command("snap", "list")
	stdout, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	output := strings.Split(string(stdout), "\n")
	installed := []string{}

	for ith, val := range output {
		if ith > 0 {
			parts := strings.Split(val, " ")
			installed = append(installed, parts[0])
		}
	}

	return installed, nil
}
