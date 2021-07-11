package main

import "os"

// Check that a file exists in the expected state. Returns true when
// a file exists according to stat, otherwise returns false and an explanation.
func TestFileDependency(tgt Dependency) (bool, []string) {
	// return false if the file path is not provided.

	if tgt["path"] == "" {
		return false, []string{"path not provided"}
	}

	fpath := tgt["path"]
	info, err := os.Stat(fpath)

	if os.IsNotExist(err) {
		return false, []string{fpath + " does not exist"}
	}

	if info.IsDir() {
		return false, []string{fpath + " is a folder"}
	}

	// lazy, but just don't point this at fifos or other weird things!
	return true, []string{fpath + " exists"}
}
