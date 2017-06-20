package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/nycdavid/dynamornr/items"
	"github.com/nycdavid/dynamornr/tables"
	"github.com/nycdavid/dynamornr/databases"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
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
				}
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
				}
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

	file, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(file, &dbyml)
	envConfigs := dbyml[env].(map[interface{}]interface{})
	dbUrl := fmt.Sprintf("%s:%v", envConfigs["host"], envConfigs["port"])
	sess, err = connectTo(dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	cli.AppHelpTemplate = `INFO:
  Repo: https://www.github.com/nycdavid/dynamornr
  Author: David Ko
  `
	app.Run(os.Args)
}

func ListTables(ctx *cli.Context) {
	database := databases.NewDatabase(ctx.String("config"))
	tables.All(database.Session)
}

func CreateTable(ctx *cli.Context) {
	database := databases.NewDatabase(ctx.String("config"))
	tables.Create(database.Session, ctx)
}

func ListItems(ctx *cli.Context) {
	database := databases.NewDatabase(ctx.String("config"))
	items.List(database.Session, ctx)
}

func SeedItems(ctx *cli.Context) {
	database := databases.NewDatabase(ctx.String("config"))
	items.Seed(database.Session, ctx)
}
