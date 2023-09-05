package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Resp struct {
	Body string
}

func handleRequest(event events.SNSEvent) (events.APIGatewayProxyResponse, error) {

	for _, record := range event.Records {
		snsRecord := record.SNS
		log.Println("Message:", snsRecord.Message)
	}

	// Save the message to a database
	region := os.Getenv("REGION")
	sess, err := session.NewSession(&aws.Config{
		Region: &region,
	})
	if err != nil {
		log.Fatal(err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	//random int between 0-1000
	id := rand.Intn(1000)

	dbSession := dynamodb.New(sess)
	_, err = dbSession.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("FlowTable"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(strconv.Itoa(id)),
			},
			"message": {
				S: aws.String(event.Records[0].SNS.Message),
			},
		},
	})
	if err != nil {
		log.Fatal(err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "",
	}, nil
}

func main() {
	runtime.Start(handleRequest)
}
