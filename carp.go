package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

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
	requires []interface{}
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

func carp(args CarpArgs) error {
	carpfile, err := readCarpFile(args.fpath)

	if err != nil {
		return err
	}

	// resolve all dependencies

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
