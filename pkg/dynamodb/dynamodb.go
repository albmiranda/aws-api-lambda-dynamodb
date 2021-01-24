package dynamodb

import (
	"go-meli/internal/satellite"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const awsRegion = "us-east-2"
const tableName = "go-transmitter"

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(awsRegion))

// Scan TODO: adicionar comentario
func Scan() (*dynamodb.ScanOutput, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	result, err := db.Scan(input)

	return result, err
}

// GetItemSatellite TODO: adicionar comentario
func GetItemSatellite(data *dynamodb.ScanOutput) ([]satellite.Data, error) {
	var satellites []satellite.Data

	err := dynamodbattribute.UnmarshalListOfMaps(data.Items, &satellites)
	if err != nil {
		return []satellite.Data{}, err
	}

	return satellites, nil
}

// NewItem TODO: adicionar comentario
func NewItem(in interface{}) (*dynamodb.PutItemInput, error) {
	item, err := dynamodbattribute.MarshalMap(in)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	return input, nil
}

// PutItem TODO: adicionar comentario
func PutItem(in *dynamodb.PutItemInput) error {
	_, err := db.PutItem(in)
	return err
}