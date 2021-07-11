package main

import (
	"log"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `Carp
Usage:
	carp <path> [--group <name>]

Description:
	Carp is a simple dependency-checker that checks a host matches the expected configuration. Dependencies are specified
	as a JSON file with groups of subdependencies. A minimal carpfile is shown below:

	{
		"folders": {
			"requires": [
				{
					"id": "core/folder",
					"path": "/home/myname/"
				}
			]
		},
		"vars": {
			"requires": [
				{
					"id": "core/envvar",
					"name": "SHELL",
					"value": "/usr/bin/zsh"
				}
			]
		},
		"main": {
			"requires": [
				{
					"id": "core/carpgroup",
					"name": "vars"
				},
				{
					"id": "core/carpgroup",
					"name": "folders"
				}
			]
		}
	}

	this example has three groups: main, vars, and folders. "vars" checks SHELL points to ZSH, and "folders" checks that
	a home folder is setup for "myname". "main" in turn checks both the "vars" and "folder" groups meet their expected
	state.

	Writing JSON by hand can be repetitious, so carp can also read stdout if it's provided an executable file that prints JSON.

Options:
	--group <name>    the group to test [default: main]
	--file <path>     the path of the carpfile
`

	opts, _ := docopt.ParseDoc(usage)
	file, err := opts.String("<path>")

	if err != nil {
		log.Fatal(err)
	}

	group := ""
	group, err = opts.String("--group")

	if err != nil {
		group = "main"
	}

	carpErr := Carp(CarpArgs{file, group})

	if carpErr != nil {
		log.Fatal(carpErr)
	}
}
