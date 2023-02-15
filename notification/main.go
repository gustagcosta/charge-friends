package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/go-co-op/gocron"
)

func main() {
	env, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	queue, err := BootstrapSQS(env.QUEUE, env.AWS_ENDPOINT, env.AWS_REGION)
	if err != nil {
		panic(err)
	}

	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Every(5).Seconds().Do(func() {
		res, err := queue.client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
			QueueUrl:            queue.queue,
			MaxNumberOfMessages: 1,
			WaitTimeSeconds:     2,
		})
		if err != nil {
			panic(err)
		}

		if len(res.Messages) == 0 {
			fmt.Println("empty queue")
			return
		}

		body := *res.Messages[0].Body

		queue.client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
			QueueUrl:      queue.queue,
			ReceiptHandle: res.Messages[0].ReceiptHandle,
		})

		charge := &Charge{}

		err = json.Unmarshal([]byte(body), &charge)
		if err != nil {
			fmt.Println("error: ", err)
		}

		fmt.Println("charge notification:", charge)
	})

	scheduler.StartBlocking()
}
