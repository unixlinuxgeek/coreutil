// TODO: need to add a description

package coreutil

import (
	"fmt"
	"testing"
)

func execute(t *testing.T, cmd string, args ...string) {
	err, i := Execute(cmd, args...)
	if err != nil {
		t.Fatalf("%s is not installed !!!\n", cmd)
	}
	fmt.Printf("%s\n", i)
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
	execute(t, "ping", "-c3", "localhost")
}
