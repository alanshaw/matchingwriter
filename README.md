# matchingwriter

[![Build Status](https://travis-ci.org/alanshaw/matchingwriter.svg?branch=master)](https://travis-ci.org/alanshaw/matchingwriter)
[![Coverage](https://codecov.io/gh/alanshaw/matchingwriter/branch/master/graph/badge.svg)](https://codecov.io/gh/alanshaw/matchingwriter)
[![Standard README](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg)](https://github.com/RichardLitt/standard-readme)
[![pkg.go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/alanshaw/matchingwriter)
[![golang version](https://img.shields.io/badge/golang-%3E%3D1.14.0-orange.svg)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/alanshaw/matchingwriter)](https://goreportcard.com/report/github.com/alanshaw/matchingwriter)

An implementation of an `io.WriteCloser` that writes a chunk to a channel when a chunk is written to the writer that contains a given string.

## Install

```sh
go get github.com/alanshaw/matchingwriter
```

## Usage

Example:

```go
package main

import (
	"fmt"
	"github.com/alanshaw/matchingwriter"
)

func main() {
	w := matchingwriter.New("unicorn", 1) // 1 is channel buffer length

	w.Write([]byte("unicorns are great!")) // will match "unicorn"
	w.Close() // closes w.C

	for match := range w.C {
		fmt.Println(match)
	}

	// prints "unicorns are great!" and then exits
}
```

## API

[pkg.go.dev Reference](https://pkg.go.dev/github.com/alanshaw/matchingwriter)

## Contribute

Feel free to dive in! [Open an issue](https://github.com/alanshaw/matchingwriter/issues/new) or submit PRs.

## License

[MIT](LICENSE) © Alan Shaw
