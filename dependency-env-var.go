package main

import "os"

// Check that an environmental variable is present, and if a value is provided
// check it matches the expected value.
func TestEnvVarDependency(tgt Dependency) (bool, []string) {
	if tgt["name"] == "" {
		return false, []string{"environmental variable not provided"}
	}

	name := tgt["name"]
	val, present := os.LookupEnv(tgt["name"])

	if !present {
		return false, []string{name + " does not exist"}
	}

	// value mismatches
	if tgt["value"] != "" && val != tgt["value"] {
		return false, []string{"environmental variable " + name + "\"" + val + "\" does not match " + tgt["value"]}
	}

	return true, []string{"environmental variable " + name + " as expected"}
}
