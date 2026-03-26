package sqs

import (
	"context"
	"encoding/json"
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

func (s *SQSService) ReceiveMessage(ctx context.Context, queueURL string) (string, error) {
	output, err := s.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: 1,
		WaitTimeSeconds:     2,
	})
	if err != nil {
		return "", fmt.Errorf("failed to receive message: %w", err)
	}

	if len(output.Messages) == 0 {
		return "", nil
	}

	msg := output.Messages[0]

	_, err = s.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: msg.ReceiptHandle,
	})
	if err != nil {
		return "", fmt.Errorf("failed to delete message: %w", err)
	}

	var envelope map[string]interface{}
	if err := json.Unmarshal([]byte(*msg.Body), &envelope); err == nil {
		pretty, err := json.MarshalIndent(envelope, "  ", "  ")
		if err == nil {
			return string(pretty), nil
		}
	}

	return *msg.Body, nil
}
