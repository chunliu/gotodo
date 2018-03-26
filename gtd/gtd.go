package main

import (
	"os"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:        "gtd",
		Usage:       "A cli to interact with gotodo web api.",
		Version:     "0.0.1",
		Description: "A cli to interact with gotodo web api.",
		Commands: []*cli.Command{
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "Get the todo list from the server.",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "id",
						Usage: "The id of the existing todo item.",
					},
				},
				Action: get,
			},
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create a new todo item. ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "The name of the new todo item.",
					},
				},
				Action: create,
			},
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "Update an existing todo item.",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "id",
						Usage: "The id of the existing todo item.",
					},
				},
				Action: update,
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "Delete an existing todo item",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "id",
						Usage: "The id of the existing todo item.",
					},
				},
				Action: delete,
			},
		},
	}

	app.Run(os.Args)
}
