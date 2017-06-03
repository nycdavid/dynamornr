package tables

import (
	"fmt"
	"io/ioutil"
	"log"
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
	Name                  string            `yaml:"name"`
	AttributeDefinitions  map[string]string `yaml:"attribute_definitions"`
	KeySchema             map[string]string `yaml:"key_schema"`
	ProvisionedThroughput struct {
		ReadCapUnits  int64 `yaml:"read_capacity_units"`
		WriteCapUnits int64 `yaml:"write_capacity_units"`
	} `yaml:"provisioned_throughput"`
}

func Create(ddbSess *dynamodb.DynamoDB) {
	schema := Schema{}
	unmarshalSchemaTo(&schema)
	for _, table := range schema.Tables {
		createTableInput := constructCti(table)
		cto, err := ddbSess.CreateTable(createTableInput)
		if err != nil {
			fmt.Println("The error is:")
			fmt.Println(err.Error())
		}
		fmt.Println(cto)
	}
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

func constructCti(table TblMetadata) *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		TableName: aws.String(table.Name),
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(table.ProvisionedThroughput.ReadCapUnits),
			WriteCapacityUnits: aws.Int64(table.ProvisionedThroughput.WriteCapUnits),
		},
		AttributeDefinitions: mapAttrDefinitions(table.AttributeDefinitions),
		KeySchema:            mapKeySchema(table.KeySchema),
	}
}

func mapAttrDefinitions(schemaDefs map[string]string) []*dynamodb.AttributeDefinition {
	var attrDefs []*dynamodb.AttributeDefinition
	for field, typ := range schemaDefs {
		attrDefs = append(attrDefs, &dynamodb.AttributeDefinition{
			AttributeName: aws.String(field),
			AttributeType: aws.String(typ),
		})
	}
	return attrDefs
}

func mapKeySchema(schemaDefs map[string]string) []*dynamodb.KeySchemaElement {
	var attrDefs []*dynamodb.KeySchemaElement
	for field, typ := range schemaDefs {
		attrDefs = append(attrDefs, &dynamodb.KeySchemaElement{
			AttributeName: aws.String(field),
			KeyType:       aws.String(typ),
		})
	}
	return attrDefs
}
