package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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

func readCarpFile(fpath string) (map[string]Group, error) {
	fileInfo, statErr := os.Stat(fpath)

	if statErr != nil {
		return nil, statErr
	}

	mode := fileInfo.Mode()

	if IsExecAny(mode) {
		cmd := exec.Command(fpath)
		byteValue, err := cmd.Output()
		fmt.Println(cmd.Stdout)

		if err != nil {
			return nil, err
		}

		var result map[string]Group
		json.Unmarshal([]byte(byteValue), &result)

		return result, nil
	} else {
		jsonFile, err := os.Open(fpath)
		if err != nil {
			return nil, err
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var result map[string]Group
		json.Unmarshal([]byte(byteValue), &result)

		return result, nil
	}
}

// DependencyResult s
type DependencyResult struct {
	Met    bool
	Reason string
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
	default:
		return DependencyResult{
			Met:    false,
			Reason: "invalid dependency.",
		}
	}
}

func epp(res DependencyResult) string {
	if res.Met {
		//return "[MET]"
		return ansi.Color("[MET]", "green")
	} else {
		return ansi.Color("[FAILED]", "red")
	}
}

func summariseResult(requiresStatus chan DependencyResult) (bool, string) {
	allMet := true
	message := ""

	for elem := range requiresStatus {
		allMet = allMet && elem.Met
		message = message + epp(elem) + " " + elem.Reason + "\n"
	}

	return allMet, message
}

// TestGroup checks a groups subdependencies in parallel
func TestGroup(carpfile map[string]Group, deps []Dependency) DependencyResult {
	requiresStatus := make(chan DependencyResult, len(deps))

	if len(deps) == 0 {
		return DependencyResult{
			Met:    true,
			Reason: "no dependencies provided.",
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

	// TODO read and summarise subdependencies.

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

	fmt.Println(groupResult.Reason)

	return nil
}
