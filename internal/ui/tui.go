package ui

import (
	"log"

	"github.com/Ayushi40804/PoRtScAnNeR/internal/scanner" // Replace with your actual module path
	"github.com/Ayushi40804/PoRtScAnNeR/internal/utils"   // Import utils package
	"github.com/rivo/tview"
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
			startScan(form, resultsTextView)
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

// startScan performs the actual scan based on user input
func startScan(form *tview.Form, resultsTextView *tview.TextView) {
	// Clear previous results
	resultsTextView.SetText("Results will appear here...\n")

	// Get values from the form
	hostname := form.GetFormItemByLabel("Hostname").(*tview.InputField).GetText()
	_, scanType := form.GetFormItemByLabel("Scan Type").(*tview.DropDown).GetCurrentOption()
	portsInput := form.GetFormItemByLabel("Ports").(*tview.InputField).GetText()

	// Convert ports from string to slice of integers using utils.ParsePorts
	ports, err := utils.ParsePorts(portsInput)
	if err != nil {
		resultsTextView.SetText("Invalid port input: " + err.Error() + "\n")
		return
	}

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
}
