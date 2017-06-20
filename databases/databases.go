package databases

import (
	"path/filepath"

	"github.com/urfave/cli"
)

type Database struct {
	Session *dynamodb.DynamoDB
}

func NewDatabase(ctx *cli.Context) *Database {
	dbyml := newDatabaseYml(ctx)
	return &Database{
		Sessions: connectTo(dbyml.url()),
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
	sess = dynamodb.New(awsSession)
	return sess, nil
}
