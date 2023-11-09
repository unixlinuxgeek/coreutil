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
func TestExecuteLS(t *testing.T) {
	execute(t, "ls")
}

func TestExecuteUname(t *testing.T) {
	execute(t, "uname")
}

func TestExecuteLSCPU(t *testing.T) {
	execute(t, "lscpu")
}
