package main

import (
	"fmt"

	"github.com/varunamachi/vaali/vcmn"
	"github.com/varunamachi/vaali/vlog"
	"github.com/varunamachi/vaali/vnet"
	"gopkg.in/urfave/cli.v1"
)

func withDefaultFlags(needsAuth bool, flg ...cli.Flag) []cli.Flag {
	if len(flg) == 0 {
		flg = make([]cli.Flag, 0, 3)
	}
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
	return []cli.Command{
		authPingCommand(),
		noAuthPingCommand(),
		entityPingCommand(),
	}
}

func authPingCommand() cli.Command {
	return cli.Command{
		Name:  "ping-u",
		Usage: "Pings sprw server as a user",
		Flags: withDefaultFlags(true),
		Action: func(ctx *cli.Context) (err error) {
			ag := vcmn.NewArgGetter(ctx)
			userID := ag.GetRequiredString("userID")
			password := ag.GetOptionalString("password")
			server := ag.GetRequiredString("server")
			if len(password) == 0 {
				password = vcmn.AskPassword("Password")
			}
			if err = ag.Err; err == nil {
				c := NewClient(server, "sprw", "v0")
				err = c.Login(userID, password)
				if err == nil {
					fmt.Println("Login successful. User: ")
					vcmn.DumpJSON(c.User)
					var session *vnet.Session
					session, err = c.Ping()
					vcmn.DumpJSON(session)
				}
			}
			return vlog.LogError("Client", err)
		},
	}
}

func entityPingCommand() cli.Command {
	return cli.Command{
		Name:  "ping-e",
		Usage: "Pings sprw server as an entity",
		Flags: withDefaultFlags(false,
			cli.StringFlag{
				Name:  "entityID",
				Value: "",
				Usage: "Unique ID of the entity",
			},
			cli.StringFlag{
				Name:  "ownerID",
				Value: "",
				Usage: "User ID of the owner of the entity",
			},
			cli.StringFlag{
				Name:  "secret",
				Value: "",
				Usage: "Entity secret for login",
			},
		),
		Action: func(ctx *cli.Context) (err error) {
			ag := vcmn.NewArgGetter(ctx)
			entityID := ag.GetRequiredString("entityID")
			userID := ag.GetRequiredString("ownerID")
			server := ag.GetRequiredString("server")
			password := ag.GetOptionalString("secret")
			if len(password) == 0 {
				password = vcmn.AskPassword("Secret")
			}
			if err = ag.Err; err == nil {
				c := NewClient(server, "sprw", "v0")
				err = c.EntityLogin(EntityCreds{
					EntityID: entityID,
					Owner:    userID,
					Secret:   password,
				})
				if err == nil {
					fmt.Println("Login successful. User: ")
					vcmn.DumpJSON(c.User)
					var session *vnet.Session
					session, err = c.Ping()
					vcmn.DumpJSON(session)
				}
			}
			return vlog.LogError("Client", err)
		},
	}
}

func noAuthPingCommand() cli.Command {
	return cli.Command{
		Name:  "ping-na",
		Usage: "Pings sprw server without authentication",
		Flags: withDefaultFlags(false),
		Action: func(ctx *cli.Context) (err error) {
			ag := vcmn.NewArgGetter(ctx)
			server := ag.GetRequiredString("server")
			if err = ag.Err; err == nil {
				c := NewClient(server, "sprw", "v0")
				var session *vnet.Session
				session, err = c.Ping()
				if err == nil {
					fmt.Println("Pinged server, Info:")
					vcmn.DumpJSON(session)
				}
			}
			return vlog.LogError("Client", err)
		},
	}
}
