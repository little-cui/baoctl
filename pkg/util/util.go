package util

import (
	"fmt"
	"os"
)

func PrintError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
}
