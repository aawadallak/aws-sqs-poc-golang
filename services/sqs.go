package services

import (
	"context"
	"log"

	"github.com/aawadallak/simple-cli-tool/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

const (
	single = "http://localhost:4566/000000000000/my-queue-test"
	batch  = "http://localhost:4566/000000000000/my-queue-test-2"
)

func publishSingleMessage(ctx context.Context) error {
	client := sqs.NewFromConfig(config.Get())

	out, err := client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(single),
		MessageBody: aws.String("Testando alguma coisa aqui"),
	})
	if err != nil {
		return err
	}
	log.Println(*out.MessageId)
	return nil
}

func publishBatchMessage(ctx context.Context) error {
	client := sqs.NewFromConfig(config.Get())
	out, err := client.SendMessageBatch(ctx, &sqs.SendMessageBatchInput{
		Entries: []types.SendMessageBatchRequestEntry{
			{Id: aws.String("1"), MessageBody: aws.String("a")},
			{Id: aws.String("2"), MessageBody: aws.String("b")},
			{Id: aws.String("3"), MessageBody: aws.String("c")},
			{Id: aws.String("4"), MessageBody: aws.String("d")},
			{Id: aws.String("5"), MessageBody: aws.String("e")},
			{Id: aws.String("6"), MessageBody: aws.String("f")},
			{Id: aws.String("7"), MessageBody: aws.String("g")},
		},
		QueueUrl: aws.String(batch),
	})
	if err != nil {
		return err
	}
	log.Println(len(out.Successful))
	return nil
}

func consumeSingleMessage(ctx context.Context) error {
	log.Println("starting single consumer")
	client := sqs.NewFromConfig(config.Get())
	for {
		out, err := client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl: aws.String(single),
			AttributeNames: []types.QueueAttributeName{
				"ApproximateReceiveCount",
				"SentTimestamp"},
		})

		if err != nil {
			return err
		}

		for _, msg := range out.Messages {
			log.Println(*msg.MessageId)
			log.Println(*msg.Body)
			log.Println(msg.Attributes)

			client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(single),
				ReceiptHandle: msg.ReceiptHandle,
			})
		}

	}
}

func consumeBatchMesasge(ctx context.Context) error {
	client := sqs.NewFromConfig(config.Get())
	log.Println("starting batch consumer")
	for {
		out, err := client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(batch),
			MaxNumberOfMessages: 10,
			AttributeNames: []types.QueueAttributeName{
				"ApproximateReceiveCount",
				"SentTimestamp"},
		})

		if err != nil {
			return err
		}

		for _, msg := range out.Messages {
			log.Println(*msg.MessageId)
			log.Println(*msg.Body)
			log.Println(msg.Attributes)

			client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(batch),
				ReceiptHandle: msg.ReceiptHandle,
			})
		}

	}
}
