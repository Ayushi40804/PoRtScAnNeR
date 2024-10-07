package ui

import (
	"log"
	"strconv"
	"strings"

	"github.com/rivo/tview"
	"github.com/yourname/portscanner/internal/scanner" // Import scanner package to perform scans
)

// ShowUI initializes and displays the terminal UI
func ShowUI(app *tview.Application) error {
	resultsTextView := tview.NewTextView().SetTextAlign(tview.AlignLeft).SetText("Results will appear here...\n")

	// Create a new form
	form := tview.NewForm().
		AddInputField("Hostname", "localhost", 20, nil, nil).
		AddDropDown("Scan Type", []string{"TCP", "UDP"}, 0, nil).
		AddInputField("Ports", "80,443,8080", 20, nil, nil).
		AddButton("Start Scan", func() {
			// Clear previous results
			resultsTextView.SetText("Results will appear here...\n")

			// Get values from the form
			hostname := form.GetFormItemByLabel("Hostname").(*tview.InputField).GetText()
			_, scanType := form.GetFormItemByLabel("Scan Type").(*tview.DropDown).GetCurrentOption()
			portsInput := form.GetFormItemByLabel("Ports").(*tview.InputField).GetText()

			// Convert ports from string to slice of integers
			ports := parsePorts(portsInput)

			// Call scanner logic based on the selected scan type
			log.Printf("Scanning %s with %s scan on ports %v", hostname, scanType, ports)
			switch scanType {
			case "TCP":
				resultsTextView.SetText("Starting TCP Scan...\n")
				scanner.TCPConnectScan(hostname, ports, resultsTextView)
			case "UDP":
				resultsTextView.SetText("Starting UDP Scan...\n")
				scanner.UDPScan(hostname, ports, resultsTextView)
			default:
				log.Println("Invalid scan type selected")
			}
		}).
		AddButton("Quit", func() {
			app.Stop() // Exit the application
		})

	// Set form layout and title
	form.SetBorder(true).SetTitle("Port Scanner").SetTitleAlign(tview.AlignLeft)

	// Set the layout of the application
	layout := tview.NewFlex().
		AddItem(form, 0, 1, true).
		AddItem(resultsTextView, 0, 3, false)

	// Set the layout as the root UI element and run the application
	return app.SetRoot(layout, true).Run()
}

// parsePorts converts a comma-separated list of port numbers (as a string) into a slice of integers.
func parsePorts(ports string) []int {
	var result []int
	for _, p := range strings.Split(ports, ",") {
		port, err := strconv.Atoi(strings.TrimSpace(p))
		if err == nil {
			result = append(result, port)
		}
	}
	return result
}
