package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

const (
	awsEndpoint = "http://localhost:4566"
	awsRegion   = "us-east-1"
)

func Get() aws.Config {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("us-east-1"),
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
			)))

	if err != nil {
		panic("configuration error, " + err.Error())
	}
	return cfg
}
