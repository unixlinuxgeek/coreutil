// TODO: need to add a description

package coreutil

import (
	"fmt"
	"testing"
)

func execute(t *testing.T, cmd string) {
	i := Execute(cmd)
	if i == false {
		t.Fatalf("%s is not installed !!!\n", cmd)
	}
	fmt.Printf("%t\n", i)
}
func executeWithArgs(t *testing.T, cmd string, args ...string) {
	i := Execute(cmd, args...)
	if i == false {
		t.Fatalf("%s is not installed !!!\n", cmd)
	}
	fmt.Printf("%t\n", i)
}
func TestExecuteLS(t *testing.T) {
	execute(t, "ls")
}

func TestExecuteUname(t *testing.T) {
	execute(t, "uname")
}

func TestExecuteLSCPU(t *testing.T) {
	execute(t, "lscpu")
}

func TestExecuteWithArgs(t *testing.T) {
	executeWithArgs(t, "ping", "-c3", "192.168.1.1")
}
