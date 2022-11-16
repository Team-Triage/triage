package reaper

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/team-triage/triage/types"
)

func writeDeadLetter(ev *kafka.Message, dynamoClient *dynamodb.DynamoDB, deadLetterTableName string) error {
	headersString := ""

	for _, header := range ev.Headers {
		headersString += header.Key + ": " + string(header.Value)
	}

	deadLetter := types.DeadLetter{
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
		fmt.Printf("Got error marshalling new deadLetter item: %s\n", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(deadLetterTableName),
	}

	_, err = dynamoClient.PutItem(input)
	if err != nil {
		fmt.Printf("Got error calling PutItem: %s\n", err)
	}

	fmt.Println("Successfully added '" + deadLetter.String + "' to table " + deadLetterTableName)
	return err
}
