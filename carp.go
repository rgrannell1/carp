package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/docopt/docopt-go"
)

// CarpArgs specifies CLI arguments
type CarpArgs struct {
	fpath string
	group string
}

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

// TestFileDependency checks a file exists
func TestFileDependency(tgt Dependency) DependencyResult {
	if tgt["path"] == "" {
		return DependencyResult{Met: false}
	}

	info, err := os.Stat(tgt["path"])
	if os.IsNotExist(err) {
		return DependencyResult{Met: false}
	}
	if info.IsDir() {
		return DependencyResult{Met: false}

	}
	return DependencyResult{Met: true}
}

// TestEnvVarDependency checks an environmental variable exists
func TestEnvVarDependency(tgt Dependency) DependencyResult {
	if tgt["name"] == "" {
		return DependencyResult{Met: false}
	}

	val, present := os.LookupEnv(tgt["name"])

	if !present {
		return DependencyResult{Met: false}
	}

	if tgt["value"] != "" && val != tgt["value"] {
		return DependencyResult{Met: false}
	}

	return DependencyResult{Met: true}
}

// TestDependency checks one dependency (or a carp-group) resolves as expected
func TestDependency(tgt Dependency) DependencyResult {
	switch id := tgt["id"]; {
	case id == "core/file":
		return TestFileDependency(tgt)
	case id == "core/apt":
		return TestEnvVarDependency(tgt)
	case id == "core/folder":
		return TestEnvVarDependency(tgt)
	case id == "core/envvar":
		return TestEnvVarDependency(tgt)
	case id == "core/carpgroup":
		return DependencyResult{Met: false}
	case id == "core/snap":
		return DependencyResult{Met: false}
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

func main() {
	usage := `Carp
Usage:
	carp --file <path> --group <name>

Options:
	--group <name> the group to test [default: main]
`

	opts, _ := docopt.ParseDoc(usage)

	file, err := opts.String("<path>")

	if err != nil {
		log.Fatal(err)
	}

	group, err := opts.String("--group")

	if err != nil {
		log.Fatal(err)
	}

	carpErr := carp(CarpArgs{file, group})

	if carpErr != nil {
		log.Fatal(carpErr)
	}
}
