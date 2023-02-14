package queue

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SqsClient struct {
	client *sqs.Client
	queue  *string
}

func BootstrapSQS(queue, awsEndpoint, awsRegion string) (*SqsClient, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		// returning EndpointNotFoundError will allow the service to fallback to its default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		return nil, fmt.Errorf("cannot load the AWS configs: %s", err)
	}

	client := sqs.NewFromConfig(cfg)

	gQInput := &sqs.GetQueueUrlInput{
		QueueName: &queue,
	}

	result, err := client.GetQueueUrl(context.TODO(), gQInput)
	if err != nil {
		return nil, fmt.Errorf("error at get queue url: %s", err)
	}

	return &SqsClient{
		client: client,
		queue:  result.QueueUrl,
	}, nil
}

func (s *SqsClient) SendMessage() (*string, error) {
	sMInput := &sqs.SendMessageInput{
		DelaySeconds: 10,
		MessageAttributes: map[string]types.MessageAttributeValue{
			"Title": {
				DataType:    aws.String("String"),
				StringValue: aws.String("The Whistler"),
			},
			"Author": {
				DataType:    aws.String("String"),
				StringValue: aws.String("John Grisham"),
			},
			"WeeksOn": {
				DataType:    aws.String("Number"),
				StringValue: aws.String("6"),
			},
		},
		MessageBody: aws.String("Information about the NY Times fiction bestseller for the week of 12/11/2016."),
		QueueUrl:    s.queue,
	}

	resp, err := s.client.SendMessage(context.TODO(), sMInput)
	if err != nil {
		return nil, err
	}

	return resp.MessageId, nil
}
