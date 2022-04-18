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

	out, err := c.CreateTopic(ctx, &sns.CreateTopicInput{
		Name:       aws.String("topico"),
		Attributes: map[string]string{},
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(*out.TopicArn)
}
