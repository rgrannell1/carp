package main

import (
	"log"

	"github.com/docopt/docopt-go"
)

// CarpArgs specifies CLI arguments
type CarpArgs struct {
	fpath string
	group string
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

	carpErr := Carp(CarpArgs{file, group})

	if carpErr != nil {
		log.Fatal(carpErr)
	}
}
