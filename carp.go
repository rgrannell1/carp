package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

// Dependency is a map specifying dependency data
type Dependency map[string]string

// Group defines a logically coherant group of dependencies
type Group struct {
	Requires []Dependency `json:"requires"`
}

func readCarpFile(fpath string) (map[string]Group, error) {
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

// DependencyResult s
type DependencyResult struct {
	Met bool
}

// TestDependency checks one dependency (or a carp-group) resolves as expected
func TestDependency(tgt Dependency) DependencyResult {
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
		return DependencyResult{Met: false}
	case id == "core/snap":
		return TestSnapDependency(tgt)
	default:
		return DependencyResult{Met: false}
	}
}

// TestGroup checks a groups subdependencies in parallel
func TestGroup(tgt Group) chan DependencyResult {
	requiresMet := make(chan DependencyResult)

	if len(tgt.Requires) == 0 {
		return requiresMet
	}

	var wg sync.WaitGroup
	wg.Add(len(tgt.Requires))

	for _, val := range tgt.Requires {
		go func(val Dependency) {
			defer wg.Done()
			requiresMet <- TestDependency(val)
		}(val)
	}

	return requiresMet
}

func carp(args CarpArgs) error {
	carpfile, err := readCarpFile(args.fpath)

	if err != nil {
		return err
	}

	tgt := carpfile[args.group]

	TestGroup(tgt)

	return nil
}
