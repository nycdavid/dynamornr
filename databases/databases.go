package databases

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/urfave/cli"
)

type Database struct {
	Session *dynamodb.DynamoDB
}

func NewDatabase(ctx *cli.Context) *Database {
	dbyml := newDatabaseYml(ctx)
	sess, err := connectTo(dbyml.url())
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Database{
		Session: sess,
	}
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
	return dynamodb.New(awsSession), nil
}
