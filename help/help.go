package help

import (
	"fmt"
	"os"
)

func Help() {
	args := os.Args
	for _, a := range args {
		if a == "help" || a == "-h" || a == "--help" {
			fmt.Println("Help is on the way")
			os.Exit(0)
		}
	}
}
