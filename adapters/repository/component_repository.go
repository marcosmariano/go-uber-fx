package repository

import (
	"context"
	"healthchecker/adapters/logger"
	"healthchecker/domain/model"
	appconfig "healthchecker/ports/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
)

type ComponentRepository struct {
	logger     logger.Logger
	config     appconfig.Config
	ddbfactory DynamoDbFactory
	tableName  string
}

func NewComponentRepository(
	l logger.Logger,
	c appconfig.Config,
	f DynamoDbFactory,
) ComponentRepository {
	return ComponentRepository{
		l,
		c,
		f,
		"component",
	}
}

func (r ComponentRepository) Save(m *model.Component) error {
	client, err := r.ddbfactory.GetClient()
	if err != nil {
		r.logger.Error("Error setting up DynamoDbClient from DdbFactory (%s) - %v", r.tableName, err)
		return errors.Wrap(err, "Error setting up DynamoDbClient from DdbFactory")
	}

	item := map[string]types.AttributeValue{
		"componentId": &types.AttributeValueMemberS{
			Value: "1234",
		},
		"name": &types.AttributeValueMemberS{
			Value: "EKS",
		},
		"url": &types.AttributeValueMemberS{
			Value: "https://eks.test.com.br/health",
		},
		"retry": &types.AttributeValueMemberS{
			Value: "2",
		},
	}

	input := &dynamodb.PutItemInput{
		Item:                   item,
		TableName:              aws.String(r.tableName),
		ReturnConsumedCapacity: "",
	}

	_, err2 := client.PutItem(context.TODO(), input)
	if err2 != nil {
		r.logger.Errorf("Error putting item in table (%s) - %v", r.tableName, err2)
		return errors.Wrap(err2, "Error loading AWS SDK config")
	}

	r.logger.Debug("Component saved")
	return nil
}
