package main

import (
	"context"
	"log"

	"github.com/aawadallak/aws-sns/cfg"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func main() {
	ctx := context.Background()
	c := cfg.GetClient(ctx)

	out, err := c.Subscribe(ctx, &sns.SubscribeInput{
		Endpoint:              aws.String("http://localhost:9090/subscribe"),
		Protocol:              aws.String("http"),
		TopicArn:              aws.String("arn:aws:sns:us-east-1:000000000000:topico"),
		ReturnSubscriptionArn: true,
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(*out.SubscriptionArn)
}
