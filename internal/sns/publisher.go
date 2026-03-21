package sns

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

const TopicName = "Publisher-Sns-Topic"

type SNSService struct {
	client *sns.Client
}

func NewSNSService(cfg aws.Config) *SNSService {
	return &SNSService{
		client: sns.NewFromConfig(cfg),
	}
}

func (s *SNSService) CreateTopic(ctx context.Context) (string, error) {
	output, err := s.client.CreateTopic(ctx, &sns.CreateTopicInput{
		Name: aws.String(TopicName),
	})
	if err != nil {
		return "", fmt.Errorf("failed to create topic: %w", err)
	}
	fmt.Printf("SNS Topic Created: %s\n", *output.TopicArn)
	return *output.TopicArn, nil
}

func (s *SNSService) Publish(ctx context.Context, TopicArn, message string) error {
	output, err := s.client.Publish(ctx, &sns.PublishInput{
		TopicArn: aws.String(TopicArn),
		Message:  aws.String(message),
	})
	if err != nil {
		return fmt.Errorf("Failed to Publish: %w", err)
	}
	fmt.Printf("Publish Successful | Message ID: %s\n", *output.MessageId)
	return nil
}
