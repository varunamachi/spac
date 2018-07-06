package main

import (
	"os"

	"github.com/varunamachi/vaali/vapp"
	"github.com/varunamachi/vaali/vcmn"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := vapp.NewSimpleApp(
		"spac",
		vcmn.Version{
			Major: 0,
			Minor: 0,
			Patch: 1,
		},
		"0",
		[]cli.Author{
			cli.Author{
				Name: "varunamachi",
			},
		},
		false,
		"Sprw client",
	)
	app.Commands = GetCommands()
	app.Run(os.Args)
}
