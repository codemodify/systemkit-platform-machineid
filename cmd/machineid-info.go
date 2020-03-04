package main

import (
	"fmt"

	machineid "github.com/codemodify/systemkit-platform-machineid"
)

func main() {
	fmt.Println(machineid.MachineID())
}
