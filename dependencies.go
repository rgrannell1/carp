package main

import (
	"os"
	"strings"
)

// TestFileDependency checks a file exists
func TestFileDependency(tgt Dependency) DependencyResult {
	if tgt["path"] == "" {
		return DependencyResult{
			Met:    false,
			Reason: "path not provided",
		}
	}

	info, err := os.Stat(tgt["path"])
	if os.IsNotExist(err) {
		return DependencyResult{
			Met:    false,
			Reason: tgt["path"] + " does not exist",
		}
	}
	if info.IsDir() {
		return DependencyResult{
			Met:    false,
			Reason: tgt["path"] + " is a folder, not a regular file",
		}

	}
	return DependencyResult{
		Met:    true,
		Reason: "file " + tgt["path"] + " exists",
	}
}

// TestEnvVarDependency checks an environmental variable exists
func TestEnvVarDependency(tgt Dependency) DependencyResult {
	if tgt["name"] == "" {
		return DependencyResult{
			Met:    false,
			Reason: "environmental variable name not provided",
		}
	}

	val, present := os.LookupEnv(tgt["name"])

	if !present {
		return DependencyResult{
			Met:    false,
			Reason: tgt["name"] + " does not exist",
		}
	}

	if tgt["value"] != "" && val != tgt["value"] {
		return DependencyResult{
			Met:    false,
			Reason: "environmental value " + tgt["name"] + " \"" + val + "\" does not match " + tgt["value"],
		}
	}

	return DependencyResult{
		Met:    true,
		Reason: "environmental variable " + tgt["name"] + " as expected.",
	}
}

func indent(content string) string {
	message := ""
	for _, key := range strings.Split(content, "\n") {
		message = message + "  " + key + "\n"
	}

	return message
}

// TestCarpGroupDependency checks
func TestCarpGroupDependency(carpfile map[string]Group, tgt Dependency) DependencyResult {
	if tgt["name"] == "" {
		return DependencyResult{
			Met:    false,
			Reason: "group name not provided",
		}
	}

	group := carpfile[tgt["name"]]

	depResult := TestGroup(carpfile, group.Requires)

	return DependencyResult{
		Met:    depResult.Met,
		Reason: "\n" + indent(depResult.Reason),
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
