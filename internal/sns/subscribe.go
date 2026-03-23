package sns

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func (s *SNSService) Subscribe(ctx context.Context, topicArn, queue string) error {
	_, err := s.client.Subscribe(ctx, &sns.SubscribeInput{
		TopicArn: aws.String(topicArn),
		Protocol: aws.String("sqs"),
		Endpoint: aws.String(queue),
	})
	if err != nil {
		return fmt.Errorf("subscribe failed: %w", err)
	}
	return nil
}
