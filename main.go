package main

import (
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"os"
)

func main() {
	err := ui.Main(createUI)
	if err != nil {
		fmt.Printf("Failed to create ui: %s", err.Error())
		os.Exit(-1)
	}
}
