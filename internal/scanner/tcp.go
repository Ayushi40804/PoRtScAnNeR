package scanner

import (
	"fmt"
	"net"
	"time"

	"github.com/rivo/tview"
)

// TCPConnectScan performs a TCP Connect scan on the specified ports
func TCPConnectScan(hostname string, ports []int, resultsTextView *tview.TextView) {
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", hostname, port)
		conn, err := net.DialTimeout("tcp", address, time.Second) // Set a timeout for connection
		if err != nil {
			resultsTextView.SetText(fmt.Sprintf("Port %d is closed\n", port))
		} else {
			resultsTextView.SetText(fmt.Sprintf("Port %d is open\n", port))
			conn.Close() // Close the connection after checking
		}
	}
}
