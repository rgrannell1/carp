package main

import (
	"encoding/json"
	"fmt"
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

// FileDependency specifies CLI arguments
type FileDependency struct {
	id   string
	path string
}

// AptDependency specifies CLI arguments
type AptDependency struct {
	id   string
	name string
}

// FolderDependency specifies CLI argument
type FolderDependency struct {
	id   string
	path string
}

// EnvVarDependency specifies CLI arguments
type EnvVarDependency struct {
	id    string
	name  string
	value string
}

// CarpGroupDependency specifies CLI arguments
type CarpGroupDependency struct {
	id   string
	name string
}

// SnapDependency specifies CLI arguments
type SnapDependency struct {
	id   string
	name string
}

// Group defines a logically coherant group of dependencies
type Group struct {
	Requires []interface{} `json:"requires"`
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

type GroupResult struct {
	Met bool
}

// DependencyResult s
type DependencyResult struct {
	Met bool
}

type Identified interface {
	Id() string
}

func Id(interface{}) {
	return ""
}

func TestDependency(tgt Identified) DependencyResult {

	switch os := Id(tgt); os {
	case "core/carpgroup":
		fmt.Println("OS X.")
	default:
		fmt.Printf("%s.\n", os)
	}

	return DependencyResult{
		Met: true,
	}
}

func TestGroup(tgt Group) chan DependencyResult {
	requiresMet := make(chan DependencyResult)

	if len(tgt.Requires) == 0 {
		return requiresMet
	}

	var wg sync.WaitGroup
	wg.Add(len(tgt.Requires))

	for _, val := range tgt.Requires {
		go func(val interface{}) {
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

	// resolve all dependencies

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
