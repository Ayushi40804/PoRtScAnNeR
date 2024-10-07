package scanner

import (
	"fmt"
	"net"
	"time"

	"github.com/rivo/tview"
)

// UDPScan performs a UDP scan on the specified ports
func UDPScan(hostname string, ports []int, resultsTextView *tview.TextView) {
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", hostname, port)
		conn, err := net.DialTimeout("udp", address, time.Second) // Set a timeout for connection
		if err != nil {
			resultsTextView.SetText(fmt.Sprintf("Port %d is closed or filtered (UDP)\n", port))
			continue
		}
		defer conn.Close()

		// Send an empty UDP packet
		_, err = conn.Write([]byte{})
		if err != nil {
			resultsTextView.SetText(fmt.Sprintf("Failed to send UDP packet to port %d\n", port))
			continue
		}
		resultsTextView.SetText(fmt.Sprintf("Port %d is open or filtered (UDP)\n", port))
	}
}
