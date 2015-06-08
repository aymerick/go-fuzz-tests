package raymond

import (
	"github.com/aymerick/raymond"
)

// Fuzz is the entry point for the go-fuzz tool.
func Fuzz(data []byte) int {
	_, err := raymond.Parse(string(data))
	if err == nil {
		return 1
	}

	return 0
}
