package items

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/urfave/cli"
)

func List(ddbSess *dynamodb.DynamoDB, ctx *cli.Context) {
	tableName := ctx.Args().Get(0)
	scanOutput := ddbSess.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})
	fmt.Println(scanOutput)
}
