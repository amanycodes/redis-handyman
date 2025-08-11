package main

import (
	"fmt"
	"os"

	"github.com/amanycodes/redis-handyman/cmd/rhm"
)

func main() {
	if err := rhm.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
