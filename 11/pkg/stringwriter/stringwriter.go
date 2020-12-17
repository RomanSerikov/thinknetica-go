package stringwriter

import (
	"io"
)

// Print - prints only strings
func Print(w io.Writer, args ...interface{}) {
	for _, arg := range args {
		if val, ok := arg.(string); ok {
			w.Write([]byte(val))
		}
	}
}
