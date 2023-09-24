// GNU Core Utilities helpers

package coreutil

import (
	"errors"
	"fmt"
	"os/exec"
)

// Installed This utility-wrapper checks whether the unix-utility is installed.
func Installed(name string) bool {
	path, err := exec.LookPath(name)
	if errors.Is(err, exec.ErrDot) {
		err = nil
	}
	if err != nil {
		fmt.Printf("%s is not installed!!!\n", name)
		return false
	}
	fmt.Printf("%s is installed in \"%s\" \n", name, path)
	return true
}
