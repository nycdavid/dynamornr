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

type Schema struct {
	Tables []TblMetadata `yaml:"tables"`
}

type TblMetadata struct {
	Name                string `yaml:"name"`
	AttributeDefintions struct {
		Id string `yaml:"Id"`
	} `yaml:"attribute_definitions"`
	KeySchema struct {
		Id string `yaml:"Id"`
	} `yaml:"key_schema"`
	ProvisionedThroughput struct {
		ReadCapUnits  int `yaml:"read_capacity_units"`
		WriteCapUnits int `yaml:"write_capacity_units"`
	} `yaml:"provisioned_throughput"`
}

func Create(ddbSess *dynamodb.DynamoDB) {
	schema := Schema{}
	unmarshalSchemaTo(&schema)
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
	fmt.Println(schema)
	for k, v := range schema.Tables {
		// createTable(k, v)
	}

	cto, err := ddbSess.CreateTable(params)
	if err != nil {
		fmt.Println("The error is:")
		fmt.Println(err.Error())
	}
	fmt.Printf("Creating Table: %s \n\n", os.Getenv("TABLENAME"))
	fmt.Println(cto)
}

func unmarshalSchemaTo(schema *Schema) {
	fpath, err := filepath.Abs("./config/schema.yml")
	if err != nil {
		log.Fatal(err)
	}
	fileContent, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(fileContent, &schema)
	if err != nil {
		log.Fatal(err)
	}
}

func createTable(name interface{}, metadata interface{}) {
	params := &dynamodb.CreateTableInput{
		TableName: aws.String(name.(string)),
	}
	fmt.Println(params)
}
