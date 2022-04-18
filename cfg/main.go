package cfg

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

const (
	awsEndpoint = "http://localhost:4566"
	awsRegion   = "us-east-1"
)

func GetClient(ctx context.Context) *sns.Client {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(awsRegion),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					if awsEndpoint != "" {
						return aws.Endpoint{
							PartitionID:   "aws",
							URL:           awsEndpoint,
							SigningRegion: awsRegion,
						}, nil
					}
					return aws.Endpoint{}, &aws.EndpointNotFoundError{}
				},
			),
		))

	if err != nil {
		panic("configuration error, " + err.Error())
	}

	return sns.NewFromConfig(cfg)
}
