package main

import (
	"fmt"

	"github.com/mitchellh/go-ps"
)

func main() {
	procs, err := ps.Processes()
	if err != nil {
		panic(err)
	}
	for _, p := range procs {
		fmt.Printf("PID: %d\tProcess: %s\n", p.Pid(), p.Executable())
	}
}
