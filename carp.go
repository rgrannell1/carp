package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type CarpArgs struct {
	fpath string
	group string
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

	return result, nil
}

// Carp runs the core application
func Carp(args CarpArgs) error {
	carpfile, err := readCarpFile(args.fpath)

	if err != nil {
		return err
	}

	met, summary := testGroup(carpfile, args.group)

	fmt.Println(summary)

	if met {
		os.Exit(0)
	} else {
		os.Exit(1)
	}

	return nil
}
