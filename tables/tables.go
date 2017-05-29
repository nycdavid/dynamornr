package tables

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func All(ddbSess *dynamodb.DynamoDB) {
	params := &dynamodb.ListTablesInput{
		ExclusiveStartTableName: aws.String("users"),
		Limit: aws.Int64(100),
	}
	lto, err := ddbSess.ListTables(params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lto)
}
