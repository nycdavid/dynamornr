package main

import (
	"os"

	"github.com/nycdavid/dynamornr/databases"
	"github.com/nycdavid/dynamornr/items"
	"github.com/nycdavid/dynamornr/tables"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:   "tables:list",
			Usage:  "List all tables.",
			Action: ListTables,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "config, c",
				},
			},
		},
		{
			Name:   "tables:create",
			Usage:  "Create a table.",
			Action: CreateTable,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "config, c",
				},
			},
		},
		{
			Name:   "items:list",
			Usage:  "List all items in a given table.",
			Action: ListItems,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "config, c",
				},
			},
		},
		{
			Name:   "items:seed",
			Usage:  "Seeds items into DynamoDB.",
			Action: SeedItems,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "config, c",
				},
			},
		},
	}

	cli.AppHelpTemplate = `INFO:
  Repo: https://www.github.com/nycdavid/dynamornr
  Author: David Ko
  `
	app.Run(os.Args)
}

func ListTables(ctx *cli.Context) {
	database := databases.NewDatabase(ctx)
	tables.All(database.Session)
}

func CreateTable(ctx *cli.Context) {
	database := databases.NewDatabase(ctx)
	tables.Create(database.Session, ctx)
}

func ListItems(ctx *cli.Context) {
	database := databases.NewDatabase(ctx)
	items.List(database.Session, ctx)
}

func SeedItems(ctx *cli.Context) {
	database := databases.NewDatabase(ctx)
	items.Seed(database.Session, ctx)
}
