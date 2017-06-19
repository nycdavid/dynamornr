package databases

import ()

type Database struct {
	configPath string
	Session    *dynamodb.DynamoDB
}

func NewDatabase(configPath string) *Database {
	dbYml := make(map[interface{}]interface{})
	return &Database{
		configPath: configPath,
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
