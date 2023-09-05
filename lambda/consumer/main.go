package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
)

type Resp struct {
	Body string
}

func handleRequest(event events.SNSEvent) (events.APIGatewayProxyResponse, error) {

	for _, record := range event.Records {
		snsRecord := record.SNS
		log.Println("Message:", snsRecord.Message)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "",
	}, nil
}

func main() {
	runtime.Start(handleRequest)
}
