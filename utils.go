package main

import "strings"

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
