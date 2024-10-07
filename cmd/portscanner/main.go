package main

import (
	"fmt"
	"os"

	"github.com/Ayushi40804/PoRtScAnNeR/internal/ui" // Replace with your actual module path
	"github.com/rivo/tview"
)

func main() {
	// Create the TUI (Terminal UI) application
	app := tview.NewApplication()

	// Show the custom UI we built in the internal/ui package
	if err := ui.ShowUI(app); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Run the application
	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}
}
