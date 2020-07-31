// Package matchingwriter contains an implementation of an "io.WriteCloser" that
// writes a chunk to a channel when a chunk is written to the writer that
// contains a given string.
//
// Example:
//
// 	package main
//
// 	import (
// 		"fmt"
// 		"github.com/alanshaw/matchingwriter"
// 	)
//
// 	func main() {
// 		w := matchingwriter.New("unicorn", 1) // 1 is channel buffer length
//
// 		w.Write([]byte("unicorns are great!")) // will match "unicorn"
// 		w.Close() // closes w.C
//
// 		for match := range w.C {
// 			fmt.Println(match)
// 		}
//
// 		// prints "unicorns are great!" and then exits
// 	}
package matchingwriter

import (
	"strings"
)

// MatchingWriter is a io.WriteCloser that matches on a substr
type MatchingWriter struct {
	str string
	C   chan string
}

// New creates a writer that writes the matched chunk to it's channel when a
// chunk is written that contains the passed substr.
func New(str string, bufferlen int) *MatchingWriter {
	matches := make(chan string, bufferlen)
	return &MatchingWriter{str, matches}
}

// Write calls `WriteString`.
func (mw *MatchingWriter) Write(p []byte) (n int, err error) {
	return mw.WriteString(string(p))
}

// WriteString checks to see if the written string contains the string to match
// and sends the written string to the channel if it does.
func (mw *MatchingWriter) WriteString(s string) (n int, err error) {
	if strings.Contains(s, mw.str) {
		select {
		case mw.C <- s:
		default:
		}
	}
	return len(s), nil
}

// Close closes the match channel.
func (mw *MatchingWriter) Close() error {
	close(mw.C)
	return nil
}
