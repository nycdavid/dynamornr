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
	json.NewDecoder(fileContent).Decode(&sf.Tables)
	input := constructBatchWriteInput(sf.Tables)
	bwo, err := ddbSess.BatchWriteItem(input)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(bwo)
}

func constructBatchWriteInput(tables interface{}) *dynamodb.BatchWriteItemInput {
	reqItems := make(map[string][]*dynamodb.WriteRequest)
	for tableName, items := range tables.(map[string]interface{}) {
		reqItems[tableName] = mapItemsToWriteRequests(items)
	}
	return &dynamodb.BatchWriteItemInput{
		RequestItems: reqItems,
	}
}

func mapItemsToWriteRequests(items interface{}) []*dynamodb.WriteRequest {
	var writeReqs []*dynamodb.WriteRequest
	for _, item := range items.([]interface{}) {
		writeReqs = append(writeReqs, &dynamodb.WriteRequest{
			PutRequest: &dynamodb.PutRequest{
				Item: mapItemToAttributeValues(item.(map[string]interface{})),
			},
		})
	}
	return writeReqs
}

func mapItemToAttributeValues(item map[string]interface{}) map[string]*dynamodb.AttributeValue {
	attrValues := make(map[string]*dynamodb.AttributeValue)
	for k, v := range item {
		attrValues[k] = &dynamodb.AttributeValue{S: aws.String(v.(string))}
	}
	return attrValues
}
