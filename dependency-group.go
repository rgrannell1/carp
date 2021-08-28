package main

import (
	"strings"
	"sync"

	"github.com/mgutz/ansi"
)

// Indent a string by a fixed amount of whitespace.
func indent(content string, count int) string {
	lines := strings.Split(content, "\n")
	padded := make([]string, len(lines))

	for ith, line := range lines {
		padded[ith] = strings.Repeat(" ", count) + line
	}

	return strings.Join(padded, "\n")
}

// Indent a list of strings by a fixed amount of whitespace.
func indentList(lines []string, count int) []string {
	padded := make([]string, len(lines))

	for ith, line := range lines {
		padded[ith] = strings.Repeat(" ", count) + line
	}

	return padded
}

// Create a string label for a requirement status
func statusLabel(status bool) string {
	if status {
		return ansi.Color("[MET]", "green")
	}

	return ansi.Color("[FAILED]", "red")
}

// Ingest dependency results from a channel; determine if all match and
//   summarise the message.
func aggregateStatus(group string, requirements chan DependencyResult) (bool, []string) {
	allMet := true
	message := []string{"groups." + group}

	for req := range requirements {
		allMet = allMet && req.Met

		status := statusLabel(req.Met)
		summary := status + " " + strings.Join(req.Reason, "\n")
		message = append(message, summary)
	}

	padded := indent("\n"+strings.Join(message, "\n"), 2)

	return allMet, []string{padded}
}

// Check whether all dependencies for a group hold.
func testGroup(facts *SystemFacts, carpfile CarpFile, group string) (bool, []string) {
	deps := carpfile.entries[group].Requires
	requirements := make(chan DependencyResult, len(deps))

	if len(deps) == 0 {
		return true, []string{"no dependencies provided."}
	}

	var wg sync.WaitGroup
	wg.Add(len(deps))

	for _, val := range deps {
		go func(val Dependency) {
			defer wg.Done()
			met, reason := TestDependency(facts, carpfile, val)

			requirements <- DependencyResult{
				Met:    met,
				Reason: reason,
			}
		}(val)
	}

	wg.Wait()
	close(requirements)

	return aggregateStatus(group, requirements)
}

// Test that all of the dependencies
func TestCarpGroupDependency(facts *SystemFacts, carpfile CarpFile, tgt Dependency) (bool, []string) {
	if tgt["name"] == "" {
		return false, []string{"group name not provided"}
	}

	result, reasons := testGroup(facts, carpfile, tgt["name"])

	return result, indentList(reasons, 2)
}
