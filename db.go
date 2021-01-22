package main

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)
const AWS_REGION = "us-east-2"
const TABLE_NAME = "go-transmitter"

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(AWS_REGION))

// GetUsers retrieves all the users from the DB
func GetSatellite() ([]SatelliteData, error) {

	input := &dynamodb.ScanInput{
	  TableName: aws.String(TABLE_NAME),
	}
	result, err := db.Scan(input)
	if err != nil {
	  return []SatelliteData{}, err
	}
	if len(result.Items) == 0 {
	  return []SatelliteData{}, nil
	}

	var satellites[]SatelliteData
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &satellites)
	if err != nil {
	  return []SatelliteData{}, err
	}

	return satellites, nil
}


func CreateSatellite(satellite SatelliteData) error {

	item, err := dynamodbattribute.MarshalMap(satellite)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(TABLE_NAME),
	}

	_, err = db.PutItem(input)
	return err
  }