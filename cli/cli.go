package cli

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// SpreadCli is the spread command line client.
type SpreadCli struct {
	// input stream (ie. stdin)
	in io.ReadCloser

	// output stream (ie. stdout)
	out io.Writer

	// error stream (ie. stderr)
	err io.Writer

	version string
}

// NewSpreadCli returns a spread command line interface (CLI) client.NewSpreadCli.
// All functionality accessible from the command line is attached to this struct.
func NewSpreadCli(in io.ReadCloser, out, err io.Writer, version string) *SpreadCli {
	return &SpreadCli{
		in:      in,
		out:     out,
		err:     err,
		version: version,
	}
}

func (c SpreadCli) printf(message string, data ...interface{}) {
	// add newline if doesn't have one
	if !strings.HasSuffix(message, "\n") {
		message = message + "\n"
	}
	fmt.Fprintf(c.out, message, data...)
}

func (c SpreadCli) fatal(message string, data ...interface{}) {
	c.printf(message, data...)
	os.Exit(1)
}
