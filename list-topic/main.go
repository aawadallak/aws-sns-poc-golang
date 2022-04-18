package main

import (
	"context"
	"log"

	"github.com/aawadallak/aws-sns/cfg"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func main() {
	ctx := context.Background()
	c := cfg.GetClient(ctx)

	out, err := c.ListTopics(ctx, &sns.ListTopicsInput{})
	if err != nil {
		log.Fatalln(err)
	}

	for _, topic := range out.Topics {
		log.Println(*topic.TopicArn)
	}
}
