package repository

import (
	"context"
	"healthchecker/adapters/logger"
	appconfig "healthchecker/ports/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/pkg/errors"
)

type DynamoDbFactory struct {
	logger logger.Logger
	config appconfig.Config
}

func NewDynamoDbFactory(l logger.Logger, c appconfig.Config) DynamoDbFactory {
	return DynamoDbFactory{
		logger: l,
		config: c,
	}
}

func (f DynamoDbFactory) GetClient() (*dynamodb.Client, error) {
	if f.config.Viper.GetString("mode") == "local" {
		return f.createLocalClient()
	}

	return f.createClient()
}

func (f DynamoDbFactory) createLocalClient() (*dynamodb.Client, error) {
	f.logger.Debug("DynamoDB Factory :: createLocalClient started")

	var endpointUrl = f.config.Viper.GetString("aws.dynamodb.endpoint")
	var region = f.config.Viper.GetString("aws.region")

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: endpointUrl, SigningRegion: region}, nil
			})),
	)

	if err != nil {
		f.logger.Errorf("Error loading AWS SDK config %v", err)
		return nil, errors.Wrap(err, "Error loading AWS SDK config")
	}

	f.logger.Debug("DynamoDB Factory :: AWS local config loaded")

	client := dynamodb.NewFromConfig(
		cfg,
		func(o *dynamodb.Options) {
			o.Credentials = credentials.NewStaticCredentialsProvider("dummy", "dummy", "")
		},
	)

	f.logger.Debug("DynamoDB Factory :: AWS local client ready")
	return client, nil
}

func (f DynamoDbFactory) createClient() (*dynamodb.Client, error) {
	f.logger.Debug("DynamoDB Factory :: createClient start")

	var region = f.config.Viper.GetString("aws.region")
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region))

	if err != nil {
		f.logger.Errorf("Error loading AWS SDK config, %v", err)
		return nil, errors.Wrap(err, "DynamoDB Factory :: Error loading AWS SDK config")
	}

	f.logger.Debug("DynamoDB Factory :: AWS config loaded")

	client := dynamodb.NewFromConfig(cfg)

	f.logger.Debug("DynamoDB Factory :: AWS client ready")

	return client, nil
}
