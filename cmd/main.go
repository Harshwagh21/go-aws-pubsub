package main

import (
	"context"
	"os"
	"time"

	awsclient "github.com/Harshwagh21/go-aws-pubsub/internal/aws"
	"github.com/Harshwagh21/go-aws-pubsub/internal/sns"
	"github.com/Harshwagh21/go-aws-pubsub/internal/sqs"
	"github.com/Harshwagh21/go-aws-pubsub/internal/ui"
)

func main() {
	ctx := context.Background()

	ui.Banner()

	ui.Bar("Loading AWS Configuration")
	cfg, err := awsclient.LocalStackConfig(ctx)
	if err != nil {
		ui.Fail("AWS Config", err.Error())
		os.Exit(1)
	}
	ui.Ok("Endpoint", "http://localhost:4566")
	ui.Ok("Region", "ap-south-1")
	ui.Ok("AWS Config", "Loaded Successfully")
	ui.Spacer()

	ui.Bar("Creating SNS Topic")
	snsService := sns.NewSNSService(cfg)
	topicArn, err := snsService.CreateTopic(ctx)
	if err != nil {
		ui.Fail("SNS Creation Failed", err.Error())
		os.Exit(1)
	}
	ui.Ok("Topic", sns.TopicName)
	ui.Ok("ARN", topicArn)
	ui.Spacer()

	ui.Bar("Creating SQS Queue")
	sqsService := sqs.NewSQSService(cfg)
	queueUrl, err := sqsService.CreateQueue(ctx)
	if err != nil {
		ui.Fail("SQS Creation Failed", err.Error())
		os.Exit(1)
	}
	ui.Ok("Queue", sqs.QueueName)
	ui.Ok("URL", queueUrl)
	ui.Spacer()

	ui.Bar("Fetching Queue ARN")
	queueArn, err := sqsService.GetQueueARN(ctx, queueUrl)
	if err != nil {
		ui.Fail("Queue ARN", err.Error())
		os.Exit(1)
	}
	ui.Ok("ARN", queueArn)
	ui.Spacer()

	ui.Bar("Subscribing SQS → SNS")
	err = snsService.Subscribe(ctx, topicArn, queueArn)
	if err != nil {
		ui.Fail("Subscription", err.Error())
		os.Exit(1)
	}
	ui.Ok("Status", "Consumer-Sqs-Queue subscribed to Publisher-Sns-Topic")
	ui.Spacer()

	var message string
	ui.Input("Enter message to publish: ", &message)
	ui.Spacer()

	ui.Bar("Publishing message to SNS")
	err = snsService.Publish(ctx, topicArn, message)
	if err != nil {
		ui.Fail("Publish", err.Error())
		os.Exit(1)
	}
	ui.MessageBox("Published", message)
	ui.Spacer()

	time.Sleep(500 * time.Millisecond)

	ui.Bar("Receiving message from SQS")
	received, err := sqsService.ReceiveMessage(ctx, queueUrl)
	if err != nil {
		ui.Fail("Receive Message", err.Error())
		os.Exit(1)
	}
	if received == "" {
		ui.MessageBox("Received", "No messages received")
	} else {
		ui.MessageBox("Received", received)
	}

	ui.Spacer()

	ui.Done()
}
