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

	out, err := c.ListSubscriptions(ctx, &sns.ListSubscriptionsInput{})

	if err != nil {
		log.Println(err)
		return
	}

	for _, sub := range out.Subscriptions {
		log.Println(*sub.SubscriptionArn)
		log.Println(*sub.Endpoint)
	}

	log.Println("finishing")
}
