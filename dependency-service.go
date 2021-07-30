package main

import (
	"fmt"
	"os/exec"
)

// Check that a file exists in the expected state. Returns true when
// a file exists according to stat, otherwise returns false and an explanation.
func TestServiceDependency(tgt Dependency) (bool, []string) {
	cmd := exec.Command("systemctl", "list-units", "--type=service", "--no-pager", "--output json")
	runErr := cmd.Run()

	if runErr != nil {
		return false, []string{"systemctl"}
	}

	stdout, err := cmd.Output()

	fmt.Println(stdout)
	fmt.Println(stdout)
	fmt.Println(stdout)
	fmt.Println(stdout)
	fmt.Println(stdout)

	if err != nil {
		return false, []string{"systemctl"}
	}

	return true, []string{}
}
