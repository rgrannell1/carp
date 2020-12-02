package main

import "os"

// TestFileDependency checks a file exists
func TestFileDependency(tgt Dependency) DependencyResult {
	if tgt["path"] == "" {
		return DependencyResult{Met: false}
	}

	info, err := os.Stat(tgt["path"])
	if os.IsNotExist(err) {
		return DependencyResult{Met: false}
	}
	if info.IsDir() {
		return DependencyResult{Met: false}

	}
	return DependencyResult{Met: true}
}

// TestEnvVarDependency checks an environmental variable exists
func TestEnvVarDependency(tgt Dependency) DependencyResult {
	if tgt["name"] == "" {
		return DependencyResult{Met: false}
	}

	val, present := os.LookupEnv(tgt["name"])

	if !present {
		return DependencyResult{Met: false}
	}

	if tgt["value"] != "" && val != tgt["value"] {
		return DependencyResult{Met: false}
	}

	return DependencyResult{Met: true}
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
