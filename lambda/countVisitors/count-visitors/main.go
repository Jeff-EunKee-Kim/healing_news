package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Item struct {
	Id       int
	Visitors int
	SiteName string
}

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func handler() (Response, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	tableName := "visitors"
	id := "1"

	// Read from dynamodb
	// result, err := svc.GetItem(&dynamodb.GetItemInput{
	// 	TableName: aws.String(tableName),
	// 	Key: map[string]*dynamodb.AttributeValue{
	// 		"id": {
	// 			N: aws.String(id),
	// 		},
	// 	},
	// })

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return Response{
	// 		Message:    "Error occured",
	// 		StatusCode: 400,
	// 	}, nil
	// }

	// item := Item{}
	// err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	// fmt.Println(item)

	// if item.Id == 0 {
	// 	return Response{
	// 		Message:    "No id found",
	// 		StatusCode: 400,
	// 	}, nil
	// }

	// Update from dynamodb

	increment := "1"

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":i": {
				N: aws.String(increment),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#V": aws.String("visitors"),
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set #V = #V + :i"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return Response{
			Message:    "Error occured",
			StatusCode: 400,
		}, nil

	}

	fmt.Println("Returning message")

	return Response{
		Message:    fmt.Sprintf("Incremented visitor count by 1"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
