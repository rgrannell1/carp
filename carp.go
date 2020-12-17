package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/mgutz/ansi"
)

// Dependency is a map specifying dependency data
type Dependency map[string]string

// Group defines a logically coherant group of dependencies
type Group struct {
	Requires []Dependency `json:"requires"`
}

// IsExecAny detects if the file is executable by the current user
func IsExecAny(mode os.FileMode) bool {
	return mode&0111 != 0
}

// readCarpFile reads (or executes) a carpfile
func readCarpFile(fpath string) (map[string]Group, error) {
	fileInfo, statErr := os.Stat(fpath)

	if statErr != nil {
		return nil, statErr
	}

	mode := fileInfo.Mode()

	var byteValue []byte

	if IsExecAny(mode) {
		cmd := exec.Command(fpath)
		stdout, err := cmd.Output()

		if err != nil {
			return nil, err
		}

		byteValue = stdout
	}

	var result map[string]Group
	err := json.Unmarshal([]byte(byteValue), &result)

	if err != nil {
		return nil, err
	}

	keys := make([]string, 0, len(result))
	for key := range result {
		keys = append(keys, key)
	}

	return result, nil
}

// DependencyResult s
type DependencyResult struct {
	Met    bool
	Reason []string
}

// TestDependency checks one dependency (or a carp-group) resolves as expected
func TestDependency(carpfile map[string]Group, tgt Dependency) DependencyResult {
	switch id := tgt["id"]; {
	case id == "core/file":
		return TestFileDependency(tgt)
	case id == "core/apt":
		return TestAptDependency(tgt)
	case id == "core/folder":
		return TestFolderDependency(tgt)
	case id == "core/envvar":
		return TestEnvVarDependency(tgt)
	case id == "core/carpgroup":
		return TestCarpGroupDependency(carpfile, tgt)
	case id == "core/snap":
		return TestSnapDependency(tgt)
	case id == "core/command":
		return TestCommand(tgt)
	default:
		return DependencyResult{
			Met:    false,
			Reason: []string{"invalid dependency."},
		}
	}
}

func statusLabel(res DependencyResult) string {
	if res.Met {
		return ansi.Color("[MET]", "green")
	}

	return ansi.Color("[FAILED]", "red")
}

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

func summariseResult(requiresStatus chan DependencyResult) (bool, []string) {
	allMet := true
	message := []string{}

	for elem := range requiresStatus {
		allMet = allMet && elem.Met

		dependencySummary := statusLabel(elem) + " " + strings.Join(elem.Reason, "\n")

		message = append(message, dependencySummary)
	}

	padded := Indent("\n"+strings.Join(message, "\n"), 2)

	return allMet, []string{padded}
}

// TestGroup checks a groups subdependencies in parallel
func TestGroup(carpfile map[string]Group, deps []Dependency) DependencyResult {
	requiresStatus := make(chan DependencyResult, len(deps))

	if len(deps) == 0 {
		return DependencyResult{
			Met:    true,
			Reason: []string{"no dependencies provided."},
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(deps))

	for _, val := range deps {
		go func(val Dependency) {
			defer wg.Done()
			requiresStatus <- TestDependency(carpfile, val)
		}(val)
	}

	wg.Wait()
	close(requiresStatus)

	allMet, summary := summariseResult(requiresStatus)

	return DependencyResult{
		Met:    allMet,
		Reason: summary,
	}
}

// Carp runs the core application
func Carp(args CarpArgs) error {
	carpfile, err := readCarpFile(args.fpath)

	if err != nil {
		return err
	}

	tgt := carpfile[args.group]

	groupResult := TestGroup(carpfile, tgt.Requires)

	fmt.Println(groupResult.Reason[0])

	if groupResult.Met == false {
		os.Exit(1)
	} else {
		os.Exit(0)
	}

	return nil
}
