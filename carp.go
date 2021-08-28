package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// Carp CLI arguments
type CarpArgs struct {
	// Carpfile
	fpath string
	// Carp group-name
	group string
}

// IsExecAny detects if the file is executable by the current user
func IsExecOwner(mode os.FileMode) bool {
	return mode&0100 != 0
}

// Represents the content of a carpfile
type CarpFile struct {
	// The map representation of a carpfile
	entries map[string]Group
}

// ReadCarpFile reads (or executes) a carpfile
func ReadCarpFile(fpath string) (CarpFile, error) {
	fileInfo, statErr := os.Stat(fpath)

	if statErr != nil {
		return CarpFile{}, statErr
	}

	var byteValue []byte

	if IsExecOwner(fileInfo.Mode()) {
		cmd := exec.Command(fpath)
		stdout, err := cmd.Output()

		if err != nil {
			return CarpFile{}, err
		}

		byteValue = stdout
	}

	// read into a carpfile variable
	var result CarpFile
	err := json.Unmarshal([]byte(byteValue), &result.entries)

	if err != nil {
		return CarpFile{}, err
	}

	return result, nil
}

type SystemFacts struct {
	AptPackages  []string
	SnapPackages []string
}

func (carpfile *CarpFile) RetrieveFacts() (SystemFacts, error) {
	// lets be lazy, and load everything once rather than inspecting which dependencies are used.

	aptPackages, err := ListAptPackages()

	if err != nil {
		return SystemFacts{}, err
	}

	snapPackages, err := ListSnapPackages()

	if err != nil {
		return SystemFacts{}, err
	}

	return SystemFacts{
		aptPackages,
		snapPackages,
	}, nil
}

// Carp runs the core application
func Carp(args CarpArgs) error {
	carpfile, err := ReadCarpFile(args.fpath)

	if err != nil {
		fmt.Printf("CARP: failed to read carpfile. %v\n", err)
		os.Exit(1)
	}

	facts, err := carpfile.RetrieveFacts()

	if err != nil {
		fmt.Printf("CARP: failed to read system facts. %v\n", err)
		os.Exit(1)
	}

	// Test some group, it's up to the user to wire everything into this group
	met, summary := testGroup(&facts, carpfile, args.group)

	if len(summary) > 0 {
		fmt.Println(summary[0])
	}

	if met {
		os.Exit(0)
	} else {
		os.Exit(1)
	}

	return nil
}
