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
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var sess *dynamodb.DynamoDB

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:   "tables:list",
			Usage:  "List all tables.",
			Action: ListTables,
		},
		{
			Name:   "tables:create",
			Usage:  "Create a table.",
			Action: CreateTable,
		},
		{
			Name:   "items:list",
			Usage:  "List all items in a given table.",
			Action: ListItems,
		},
	}
	dbyml := make(map[interface{}]interface{})
	fpath, err := filepath.Abs("./config/database.yml")
	if err != nil {
		log.Fatal(err)
	}
	file, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}
	env := os.Getenv("ENV")
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
	tables.All(sess)
}

func CreateTable(ctx *cli.Context) {
	tables.Create(sess)
}

func ListItems(ctx *cli.Context) {
	items.List(sess, ctx)
}

func connectTo(url string) (*dynamodb.DynamoDB, error) {
	awsSession, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(url),
		Region:   aws.String("us-east-1"),
	})
	if err != nil {
		log.Fatal(err)
		return &dynamodb.DynamoDB{}, err
	}
	sess = dynamodb.New(awsSession)
	return sess, nil
}
