package items

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/urfave/cli"
)

type SeedFile struct {
	Tables interface{}
}

type Table struct {
	Items []map[string]interface{}
}

func List(ddbSess *dynamodb.DynamoDB, ctx *cli.Context) {
	tableName := ctx.Args().Get(0)
	scanOutput, err := ddbSess.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(scanOutput)
}

func Seed(ddbSess *dynamodb.DynamoDB, ctx *cli.Context) {
	fpath, err := filepath.Abs("./config/seeds.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	fileContent, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err.Error())
	}
	var sf SeedFile
	json.NewDecoder(fileContent).Decode(&sf)
	fmt.Println(sf.Tables)
}
