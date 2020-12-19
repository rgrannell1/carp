package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// TestFileDependency checks a file exists
func TestFileDependency(tgt Dependency) DependencyResult {
	if tgt["path"] == "" {
		return DependencyResult{
			Met:    false,
			Reason: []string{"path not provided"},
		}
	}

	info, err := os.Stat(tgt["path"])
	if os.IsNotExist(err) {
		return DependencyResult{
			Met:    false,
			Reason: []string{tgt["path"] + " does not exist"},
		}
	}
	if info.IsDir() {
		return DependencyResult{
			Met:    false,
			Reason: []string{tgt["path"] + " is a folder, not a regular file"},
		}

	}
	return DependencyResult{
		Met:    true,
		Reason: []string{"file " + tgt["path"] + " exists"},
	}
}

// TestEnvVarDependency checks an environmental variable exists
func TestEnvVarDependency(tgt Dependency) DependencyResult {
	if tgt["name"] == "" {
		return DependencyResult{
			Met:    false,
			Reason: []string{"environmental variable name not provided"},
		}
	}

	val, present := os.LookupEnv(tgt["name"])

	if !present {
		return DependencyResult{
			Met:    false,
			Reason: []string{tgt["name"] + " does not exist"},
		}
	}

	if tgt["value"] != "" && val != tgt["value"] {
		return DependencyResult{
			Met:    false,
			Reason: []string{"environmental value " + tgt["name"] + " \"" + val + "\" does not match " + tgt["value"]},
		}
	}

	return DependencyResult{
		Met:    true,
		Reason: []string{"environmental variable " + tgt["name"] + " as expected."},
	}
}

// TestCarpGroupDependency checks
func TestCarpGroupDependency(carpfile map[string]Group, tgt Dependency) DependencyResult {
	if tgt["name"] == "" {
		return DependencyResult{
			Met:    false,
			Reason: []string{"group name not provided"},
		}
	}

	depResult := TestGroup(carpfile, tgt["name"])

	return DependencyResult{
		Met:    depResult.Met,
		Reason: IndentList(depResult.Reason, 2),
	}
}

// TestAptDependency checks
func TestAptDependency(tgt Dependency) DependencyResult {
	return DependencyResult{
		Met:    true,
		Reason: []string{"unimplemented."},
	}
}

// TestFolderDependency checks
func TestFolderDependency(tgt Dependency) DependencyResult {
	if tgt["path"] == "" {
		return DependencyResult{
			Met:    false,
			Reason: []string{"path not provided"},
		}
	}

	info, err := os.Stat(tgt["path"])
	if os.IsNotExist(err) {
		return DependencyResult{
			Met:    false,
			Reason: []string{tgt["path"] + " does not exist"},
		}
	}
	if !info.IsDir() {
		return DependencyResult{
			Met:    false,
			Reason: []string{tgt["path"] + " is not a folder"},
		}
	}

	return DependencyResult{
		Met:    true,
		Reason: []string{"folder " + tgt["path"] + " exists"},
	}
}

func listSnapPackages() ([]string, error) {
	cmd := exec.Command("snap", "list")
	stdout, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	output := strings.Split(string(stdout), "\n")
	installed := []string{}

	for ith, val := range output {
		if ith > 0 {
			parts := strings.Split(val, " ")
			installed = append(installed, parts[0])
		}
	}

	return installed, nil
}

// TestSnapDependency checks
func TestSnapDependency(tgt Dependency) DependencyResult {
	packages, err := listSnapPackages()

	if err != nil {
		return DependencyResult{
			Met:    false,
			Reason: []string{"failed to list snap packages."},
		}
	}

	for _, name := range packages {
		if name == tgt["name"] {
			return DependencyResult{
				Met:    true,
				Reason: []string{"snap package \"" + tgt["name"] + "\" installed"},
			}
		}
	}

	return DependencyResult{
		Met:    false,
		Reason: []string{tgt["name"] + " not in listed snap-packages"},
	}
}

// TestCommand checks
func TestCommand(tgt Dependency) DependencyResult {
	if tgt["name"] == "" {
		return DependencyResult{
			Met:    false,
			Reason: []string{"group name not provided"},
		}
	}

	cmd := exec.Command("which", tgt["name"])

	if err := cmd.Start(); err != nil {
		return DependencyResult{
			Met:    false,
			Reason: []string{"which error"},
		}
	}

	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				log.Printf("Exit Status: %d", status.ExitStatus())
			}
		} else {
			return DependencyResult{
				Met:    false,
				Reason: []string{"which error"},
			}
		}
	}

	return DependencyResult{Met: true}
}
