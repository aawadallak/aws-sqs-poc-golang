package main

import (
	"context"
	"log"

	"github.com/aawadallak/simple-cli-tool/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func main() {
	CreateSQSQueue()
}

func CreateSQSQueue() {
	cfg := config.Get()
	client := sqs.NewFromConfig(cfg)

	out, err := client.CreateQueue(context.TODO(), &sqs.CreateQueueInput{
		QueueName: aws.String("my-queue-test-2"),
	})
	if err != nil {
		log.Println(err)
	}

	log.Println(*out.QueueUrl)
}
