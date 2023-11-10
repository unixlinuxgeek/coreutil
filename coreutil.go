// GNU Core Utilities helpers

package coreutil

import (
	"errors"
	"fmt"
	"log"
	"os"
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

// Execute the shell command
func Execute(cmd string, args ...string) bool {
	i := Installed(cmd)
	if i == false {
		_ = fmt.Errorf("%s is not installed ", cmd)
		return false
	}
	// Do something...
	c := exec.Command(cmd, args...)
	stdout, err := c.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(os.Stdout, stdout)
	//fmt.Println("command executed successfully")
	return true
}
