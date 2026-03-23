package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

const QueueName = "Consumer-Sqs-Queue"

type SQSService struct {
	client *sqs.Client
}

func NewSQSService(cfg aws.Config) *SQSService {
	return &SQSService{
		client: sqs.NewFromConfig(cfg),
	}
}

func (s *SQSService) CreateQueue(ctx context.Context) (string, error) {
	output, err := s.client.CreateQueue(ctx, &sqs.CreateQueueInput{
		QueueName: aws.String(QueueName),
	})
	if err != nil {
		return "", fmt.Errorf("failed to create queue: %w", err)
	}
	fmt.Printf("SQS Queue Created: %s\n", *output.QueueUrl)
	return *output.QueueUrl, nil
}

func (s *SQSService) GetQueueARN(ctx context.Context, queueURL string) (string, error) {
	output, err := s.client.GetQueueAttributes(ctx, &sqs.GetQueueAttributesInput{
		QueueUrl:       aws.String(queueURL),
		AttributeNames: []types.QueueAttributeName{"QueueArn"},
	})
	if err != nil {
		return "", fmt.Errorf("failed to get queue arn: %w", err)
	}

	return output.Attributes["QueueArn"], nil
}

func (s *SQSService) ReceiveMessage(ctx context.Context, queueURL string) error {
	output, err := s.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: 10,
		WaitTimeSeconds:     1,
	})
	if err != nil {
		return fmt.Errorf("failed to receive message: %w", err)
	}

	if len(output.Messages) == 0 {
		fmt.Println("No messages received")
		return nil
	}

	for _, msg := range output.Messages {
		fmt.Printf("Received message: %s\n", *msg.Body)
	}

	return nil
}
