package client

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

// rds client

func ClientRds(ctx context.Context, _region string, _accessKeyId string, _accessKeySecret string) (*rds.Client, error) {
	// 设置 AK 和 SK
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(_accessKeyId, _accessKeySecret, "")),
		config.WithRegion(_region),
	)
	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		return nil, err
	}

	return rds.NewFromConfig(cfg), nil

}

// EC2 client

func ClientEC2(ctx context.Context, _region string, _accessKeyId string, _accessKeySecret string) (*ec2.Client, error) {
	// 设置 AK 和 SK
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(_accessKeyId, _accessKeySecret, "")),
		config.WithRegion(_region),
	)
	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		return nil, err
	}

	return ec2.NewFromConfig(cfg), nil
}
