package main

import (
	"log"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `Carp
Usage:
	carp <path> [--group <name>]
	carp (-h|--help)

Description:
	Carp is a dependency-checker that checks a host matches the expected configuration. Dependencies are specified
	as a JSON file with groups of subdependencies. A carpfile is shown below:

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

  It's recommended to use a group for each device you would like to configure; e.g laptop, raspberry pi, VM, etc.

	Writing JSON by hand can be repetitious, so carp can also read stdout if it's provided an executable file that prints JSON.

Dependencies:
  Carp can check several types of dependency; the required parameters are:

  id-property   |  properties
  --------------------------------
  core/service     { name: <str> }
	core/file        { path: <str> }
	core/apt         { name: <str> }
	core/folder      { path: <str> }
	core/envvar      { name: <str>, value: <str> }
	core/carpgroup   { <groupname>: requires[ <dependency> ] }
	core/snap        { name: <str> }
	core/command     { name: <str> }

core/service
  Do not use, yet

core/file
  Does the specified file exist?

core/apt
  Is a specified apt-package installed, according to 'apt list' for the current user?

core/folder
  Does the specified folder exist?

core/envvar
  Does the environmental variable exist? If value is specified, check the value matches the expected value

core/carpgroup
  Are all subdependencies in this group met?

core/snap
  Is a specified snap package installed, according to 'snap list' for the current user?

core/command
  Does the command exist on PATH for the current user, according to 'which'?

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
