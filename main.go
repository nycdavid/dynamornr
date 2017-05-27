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
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

func main() {
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
	_, err = connectTo(dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "Fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Printf("Connected to database on %s:%d...\n", envConfigs["host"], envConfigs["port"])
		return nil
	}
	app.Run(os.Args)
}

func connectTo(url string) (*dynamodb.DynamoDB, error) {
	awsSession, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(url),
		Region:   aws.String("us-east-1"),
	})
	if err != nil {
		log.Fatal(err)
	}
	sess := dynamodb.New(awsSession)
	params := &dynamodb.ListTablesInput{
		ExclusiveStartTableName: aws.String("users"),
		Limit: aws.Int64(100),
	}
	_, err = sess.ListTables(params)
	if err != nil {
		return &dynamodb.DynamoDB{}, err
	}
	return sess, nil
}
