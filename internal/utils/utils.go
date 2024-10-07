package utils

import (
	"strconv"
	"strings"
)

// ParsePorts converts a comma-separated list of port numbers (as a string) into a slice of integers.
func ParsePorts(ports string) ([]int, error) {
	var result []int
	for _, p := range strings.Split(ports, ",") {
		port, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return nil, err // Return an error if the conversion fails
		}
		result = append(result, port)
	}
	return result, nil
}

// IsValidPort checks if the given port number is valid (between 1 and 65535).
func IsValidPort(port int) bool {
	return port >= 1 && port <= 65535
}
