package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/jsii-runtime-go"
)

type Resp struct {
	Body string
}

func handleRequest(event events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2"),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	svc := sns.New(sess)

	top := os.Getenv("TOPIC_ARN")

	resp := Resp{
		Body: fmt.Sprintf("Publishing to SNS topic %s", top),
	}

	msg, _ := json.Marshal(resp)

	log.Println("Publishing to SNS topic:", top)

	r, err := svc.Publish(&sns.PublishInput{
		Message:  jsii.String("test message"),
		TopicArn: &top,
	})

	if err != nil {
		log.Println("Unable to publish to SNS topic", err.Error())
		log.Fatal(err.Error())
	}

	log.Println("SNS Message published. Message ID:", *r.MessageId)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(msg),
	}, nil
}

func main() {
	runtime.Start(handleRequest)
}
