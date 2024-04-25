package infrastructure

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewAWSConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = os.Getenv("AWS_REGION")
		return nil

	})
	if err != nil {
		panic(err)
	}

	return cfg
}

func NewDynamoDBClient(sdkConfig aws.Config) *dynamodb.Client {
	return dynamodb.NewFromConfig(sdkConfig)
}
