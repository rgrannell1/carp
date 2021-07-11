package main

import "os"

// Checks that a folder exists
func TestFolderDependency(tgt Dependency) (bool, []string) {
	if tgt["path"] == "" {
		return false, []string{"path not provided"}
	}

	info, err := os.Stat(tgt["path"])
	if os.IsNotExist(err) {
		return false, []string{tgt["path"] + " does not exist"}
	}
	if !info.IsDir() {
		return false, []string{tgt["path"] + " is not a folder"}
	}

	return true, []string{"folder " + tgt["path"] + " exists"}
}
