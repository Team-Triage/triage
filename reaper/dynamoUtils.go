package reaper

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/google/uuid"
)

func makeDynamoSession() *session.Session {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigDisable,
	}))
	return sess
}

func makeDynamoClient(sess *session.Session) *dynamodb.DynamoDB {
	svc := dynamodb.New(sess)
	return svc
}

func createUUID() string {
	uuid := uuid.New()
	return uuid.String()
}
