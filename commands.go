package main

import (
	"fmt"

	"github.com/varunamachi/vaali/vcmn"
	"github.com/varunamachi/vaali/vlog"
	"github.com/varunamachi/vaali/vnet"
	"gopkg.in/urfave/cli.v1"
)

func withDefaultFlags(flg []cli.Flag) []cli.Flag {
	return append(flg, []cli.Flag{
		cli.StringFlag{
			Name:  "server",
			Value: "http://localhost:8000",
			Usage: "Server url of for <protocol>://<host>:<port>",
		},
		cli.StringFlag{
			Name:  "userID",
			Value: "",
			Usage: "Sprw user ID",
		},
		cli.StringFlag{
			Name:   "password",
			Value:  "",
			Usage:  "Sprw password",
			Hidden: true,
		},
	}...)
}

//GetCommands - get commands for spc
func GetCommands() []cli.Command {
	return []cli.Command{}
}

func authPingCommand() {
	return &cli.Command{
		Name:  "ping",
		Usage: "Pings sprw server",
		Flags: withDefaultFlags([]cli.Flag{}),
		Action: func(ctx *cli.Context) (err error) {
			ag := vcmn.NewArgGetter(ctx)
			userID := ag.GetRequiredString("userID")
			password := ag.GetOptionalString("password")
			if len(password) == 0 {
				password = vcmn.AskPassword("Password")
			}
			if err = ag.Err; err == nil {
				c := vnet.NewClient("http://localhost:8000", "vaali", "v0")
				err = c.Login(userID, password)
				if err == nil {
					fmt.Println("Login successful. User: ")
					vcmn.DumpJSON(c.User)
				}
			}
			return vlog.LogError("Client", err)
		},
	}
}
