package tables

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"gopkg.in/yaml.v2"
)

func All(ddbSess *dynamodb.DynamoDB) {
	params := &dynamodb.ListTablesInput{
		Limit: aws.Int64(100),
	}
	lto, err := ddbSess.ListTables(params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table Listing:")
	fmt.Println(lto)
}

func Create(ddbSess *dynamodb.DynamoDB) {
	schemaYaml()
	params := &dynamodb.CreateTableInput{
		TableName: aws.String(os.Getenv("TABLENAME")),
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       aws.String("HASH"),
			},
		},
	}
	cto, err := ddbSess.CreateTable(params)
	if err != nil {
		fmt.Println("The error is:")
		fmt.Println(err.Error())
	}
	fmt.Printf("Creating Table: %s \n\n", os.Getenv("TABLENAME"))
	fmt.Println(cto)
}

func schemaYaml() {
	fpath, err := filepath.Abs("./config/schema.yml")
	if err != nil {
		log.Fatal(err)
	}
	schema := make(map[interface{}]interface{})
	fileContent, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(fileContent, &schema)
	fmt.Println(schema)
}
