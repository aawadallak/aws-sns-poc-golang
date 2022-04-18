package main

import (
	"context"
	"log"

	"github.com/aawadallak/aws-sns/cfg"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

const TopicArn = "arn:aws:sns:us-east-1:000000000000:topico"

func main() {
	ctx := context.Background()
	c := cfg.GetClient(ctx)

	out, err := c.Publish(ctx, &sns.PublishInput{
		Message:  aws.String("test from a sns topic"),
		TopicArn: aws.String(TopicArn),
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(*out.MessageId)
}
