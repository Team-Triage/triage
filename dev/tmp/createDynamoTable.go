package tmp

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func makeDynamoClient(sess *session.Session) *dynamodb.DynamoDB {
	svc := dynamodb.New(sess)
	return svc
}

func makeDynamoSession() *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return sess
}

func createDynamoTable() {
	svc := makeDynamoClient(makeDynamoSession())
	tableName := "Dead_Letters"

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("UUID"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("TIMESTAMP"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("UUID"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("TIMESTAMP"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		log.Fatalf("Got error calling CreateTable: %s", err)
	}

	fmt.Println("Created the table", tableName)

}
