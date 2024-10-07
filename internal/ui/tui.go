package ui

import (
	"log"
	"strings"

	"github.com/Ayushi40804/PoRtScAnNeR/internal/scanner" // Replace with your actual module path
	"github.com/Ayushi40804/PoRtScAnNeR/internal/utils"   // Import utils package
	"github.com/rivo/tview"
)

// ShowUI initializes and displays the terminal UI
func ShowUI(app *tview.Application) error {
	resultsTextView := tview.NewTextView().SetTextAlign(tview.AlignLeft).SetText("Results will appear here...\n")

	// Create input fields for hostname and ports
	hostnameInput := tview.NewInputField().
		SetLabel("Hostname: ").SetText("localhost")

	portsInput := tview.NewInputField().
		SetLabel("Ports: ").SetText("80,443,8080")

	// Create dropdown for scan type
	scanTypeDropdown := tview.NewDropDown().
		SetLabel("Scan Type: ").SetOptions([]string{"TCP", "UDP"}, nil)

	// Create buttons
	startButton := tview.NewButton("Start Scan").SetSelectedFunc(func() {
		hostname := hostnameInput.GetText()
		ports := portsInput.GetText()
		_, scanType := scanTypeDropdown.GetCurrentOption()

		startScan(hostname, scanType, ports, resultsTextView)
	})

	quitButton := tview.NewButton("Quit").SetSelectedFunc(func() {
		app.Stop()
	})

	// Layout the inputs and buttons manually
	formLayout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(hostnameInput, 1, 1, true).
		AddItem(portsInput, 1, 1, true).
		AddItem(scanTypeDropdown, 1, 1, true).
		AddItem(startButton, 1, 1, true).
		AddItem(quitButton, 1, 1, true)

	// Set layout with form and results view
	layout := tview.NewFlex().
		AddItem(formLayout, 0, 1, true).
		AddItem(resultsTextView, 0, 3, false)

	// Set the root UI element and run the application
	return app.SetRoot(layout, true).Run()
}

// startScan performs the actual scan based on user input
func startScan(hostname, scanType, portsInput string, resultsTextView *tview.TextView) {
	// Clear previous results
	resultsTextView.SetText("Results will appear here...\n")

	// Convert ports from string to slice of integers using utils.ParsePorts
	ports, err := utils.ParsePorts(portsInput)
	if err != nil {
		resultsTextView.SetText("Invalid port input: " + err.Error() + "\n")
		return
	}

	// Call scanner logic based on the selected scan type
	log.Printf("Scanning %s with %s scan on ports %v", hostname, scanType, ports)
	switch strings.ToUpper(scanType) {
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
