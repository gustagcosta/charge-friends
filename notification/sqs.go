package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
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

func (s *SqsClient) SendMessage(body string) (*string, error) {
	sMInput := &sqs.SendMessageInput{
		DelaySeconds: 0,
		MessageBody:  &body,
		QueueUrl:     s.queue,
	}

	resp, err := s.client.SendMessage(context.TODO(), sMInput)
	if err != nil {
		return nil, err
	}

	return resp.MessageId, nil
}

func (s *SqsClient) ReadQueue() (*string, error) {
	res, err := s.client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:            s.queue,
		MaxNumberOfMessages: 1,
		WaitTimeSeconds:     5,
	})

	if err != nil {
		return nil, err
	}

	if len(res.Messages) == 0 {
		return nil, errors.New("empty queue")
	}

	body := *res.Messages[0].Body

	s.client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
		QueueUrl:      s.queue,
		ReceiptHandle: res.Messages[0].ReceiptHandle,
	})

	return &body, nil
}
