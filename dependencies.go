package main

import (
	"log"
	"os"
	"os/exec"
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

func indent(content []string) []string {
	indented := []string{}

	for _, key := range content {
		indented = append(indented, "  "+key)
	}

	return indented
}

// TestCarpGroupDependency checks
func TestCarpGroupDependency(carpfile map[string]Group, tgt Dependency) DependencyResult {
	if tgt["name"] == "" {
		return DependencyResult{
			Met:    false,
			Reason: []string{"group name not provided"},
		}
	}

	group := carpfile[tgt["name"]]

	depResult := TestGroup(carpfile, group.Requires)

	return DependencyResult{
		Met:    depResult.Met,
		Reason: indent(depResult.Reason),
	}
}

// TestAptDependency checks
func TestAptDependency(tgt Dependency) DependencyResult {
	return DependencyResult{Met: true}
}

// TestFolderDependency checks
func TestFolderDependency(tgt Dependency) DependencyResult {
	return DependencyResult{Met: true}
}

// TestSnapDependency checks
func TestSnapDependency(tgt Dependency) DependencyResult {
	return DependencyResult{Met: true}
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
