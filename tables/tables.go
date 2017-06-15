package tables

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/urfave/cli"
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

type ProvisionedThroughput struct {
	ReadCapUnits  int64 `yaml:"read_capacity_units"`
	WriteCapUnits int64 `yaml:"write_capacity_units"`
}

type GlobalSecIdx struct {
	Name                  string                 `yaml:"name"`
	KeySchema             map[string]string      `yaml:"key_schema"`
	Projection            map[string]interface{} `yaml:"projection"`
	ProvisionedThroughput ProvisionedThroughput  `yaml:"provisioned_throughput"`
}

type TblMetadata struct {
	Name                  string                `yaml:"name"`
	AttributeDefinitions  map[string]string     `yaml:"attribute_definitions"`
	KeySchema             map[string]string     `yaml:"key_schema"`
	GSIs                  []GlobalSecIdx        `yaml:"global_secondary_indexes"`
	ProvisionedThroughput ProvisionedThroughput `yaml:"provisioned_throughput"`
}

func Create(ddbSess *dynamodb.DynamoDB, ctx *cli.Context) {
	schema := Schema{}
	unmarshalSchemaTo(&schema, ctx)
	for _, table := range schema.Tables {
		createTableInput := constructCti(table)
		cto, err := ddbSess.CreateTable(createTableInput)
		if err != nil {
			stdout := fmt.Sprintf("ERROR: %s", err.Error())
			fmt.Println(stdout)
		}
		fmt.Println(cto)
	}
}

func unmarshalSchemaTo(schema *Schema, ctx *cli.Context) {
	fpath, err := filepath.Abs("./config/schema.yml")

	if err != nil {
		log.Fatal(err)
	}
	fileContent, err := ioutil.ReadFile(fpath)
	if err != nil {
		str := fmt.Sprintf("ioutil.ReadFile error: %s", err.Error())
		log.Print(str)
		fmt.Println(ctx.String("config"))
		fileContent, _ = ioutil.ReadFile(fmt.Sprintf("%s/schema.yml", ctx.String("config")))
	}
	err = yaml.Unmarshal(fileContent, &schema)
	if err != nil {
		log.Fatal(err)
	}
}

func constructCti(table TblMetadata) *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		TableName:              aws.String(table.Name),
		ProvisionedThroughput:  mapProvisionedThroughput(table.ProvisionedThroughput),
		AttributeDefinitions:   mapAttrDefinitions(table.AttributeDefinitions),
		KeySchema:              mapKeySchema(table.KeySchema),
		GlobalSecondaryIndexes: mapGsi(table.GSIs),
	}
}

func mapProvisionedThroughput(pt ProvisionedThroughput) *dynamodb.ProvisionedThroughput {
	return &dynamodb.ProvisionedThroughput{
		ReadCapacityUnits:  aws.Int64(pt.ReadCapUnits),
		WriteCapacityUnits: aws.Int64(pt.WriteCapUnits),
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

func mapGsi(indexDefs []GlobalSecIdx) []*dynamodb.GlobalSecondaryIndex {
	var gsis []*dynamodb.GlobalSecondaryIndex
	for _, gsIdx := range indexDefs {
		gsis = append(gsis, &dynamodb.GlobalSecondaryIndex{
			IndexName:             aws.String(gsIdx.Name),
			KeySchema:             mapKeySchema(gsIdx.KeySchema),
			Projection:            mapProjection(gsIdx.Projection),
			ProvisionedThroughput: mapProvisionedThroughput(gsIdx.ProvisionedThroughput),
		})
	}
	return gsis
}

func mapProjection(proj map[string]interface{}) *dynamodb.Projection {
	var nkAttrs []*string
	for _, attr := range proj["non_key_attributes"].([]interface{}) {
		nkAttrs = append(nkAttrs, aws.String(attr.(string)))
	}
	return &dynamodb.Projection{
		NonKeyAttributes: nkAttrs,
		ProjectionType:   aws.String(proj["projection_type"].(string)),
	}
}
