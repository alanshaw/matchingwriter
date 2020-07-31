package matchingwriter

import (
	"fmt"
	"math/rand"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func randomStrings(including []string) []string {
	var out []string
	count := rand.Intn(5)
	for i := 0; i < count; i++ {
		out = append(out, randomString(rand.Intn(10)+1))
	}
	for _, s := range including {
		out = append(out, s)
		count := rand.Intn(5)
		for i := 0; i < count; i++ {
			out = append(out, randomString(rand.Intn(10)+1))
		}
	}
	return out
}

func TestMatchingWriter(t *testing.T) {
	substr := "match"
	testmatches := []string{
		fmt.Sprintf("%s", substr),
		fmt.Sprintf("in the middle %s of", substr),
		fmt.Sprintf("containedwithin%sanother", substr),
		fmt.Sprintf("ends with %s", substr),
		fmt.Sprintf("%s starts with", substr),
	}

	input := randomStrings(testmatches)
	mw := New(substr, len(input))

	go func() {
		for _, s := range input {
			mw.Write([]byte(s))
		}
		mw.Close()
	}()

	var matches []string
	for match := range mw.C {
		matches = append(matches, match)
	}

	if len(matches) != len(testmatches) {
		t.Fatal(fmt.Sprintf("received an unexpected number of matches: %d vs %d", len(matches), len(testmatches)))
	}

	for i, m := range testmatches {
		if matches[i] != m {
			t.Fatal("expected match \"" + matches[i] + "\" to match \"" + m + "\"")
		}
	}
}
