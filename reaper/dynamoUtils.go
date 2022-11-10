package reaper

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/google/uuid"
)

type DeadLetter struct {
	UUID      string
	TIMESTAMP string
	Topic     string
	Partition int
	Offset    int
	Key       string
	String    string
	Headers   string
}

func makeDynamoSession() *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return sess
}

func createUUID() string {
	uuid := uuid.New()
	return uuid.String()
}

func makeDynamoClient(sess *session.Session) *dynamodb.DynamoDB {
	svc := dynamodb.New(sess)
	return svc
}

func writeDeadLetter(ev *kafka.Message, svc *dynamodb.DynamoDB) error {
	headersString := ""

	for _, header := range ev.Headers {
		headersString += header.Key + ": " + string(header.Value)
	}

	deadLetter := DeadLetter{
		UUID:      createUUID(),
		TIMESTAMP: ev.Timestamp.String(),
		Topic:     *ev.TopicPartition.Topic,
		Partition: int(ev.TopicPartition.Partition),
		Offset:    int(ev.TopicPartition.Offset),
		Key:       string(ev.Key),
		String:    string(ev.Value),
		Headers:   headersString,
	}

	av, err := dynamodbattribute.MarshalMap(deadLetter)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	tableName := "Dead_Letters" // should probably be topic_dead_letters; maybe an env var

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	fmt.Println("Successfully added '" + deadLetter.String + "' to table " + tableName)
	return err
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
