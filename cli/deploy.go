package cli

import (
	"fmt"
	"os"

	"rsprd.com/spread/pkg/input/dir"

	"github.com/codegangsta/cli"
)

// Version returns the current spread version
func (spread SpreadCli) Dir() *cli.Command {
	return &cli.Command{
		Name:  "file",
		Usage: "File info",
		Action: func(c *cli.Context) {
			wd, _ := os.Getwd()
			fmt.Fprintf(spread.out, "Current directory is %s\n", wd)

			input, err := dir.NewFileInput(wd)
			if err != nil {
				panic(err)
			}

			e, err := input.Build()
			if err != nil {
				panic(err)
			}

			dep, err := e.Deployment()
			if err != nil {
				panic("Could not deploy: " + err.Error())
			}

			fmt.Printf("Found %d objects in `%s`.", dep.Len(), wd)
			fmt.Println(dep.String())
		},
	}
}
