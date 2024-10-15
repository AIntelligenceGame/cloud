package aws_rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func Regions(ctx context.Context, _region string, _accessKeyId string, _accessKeySecret string) (*rds.DescribeSourceRegionsOutput, error) {
	// 设置 AK 和 SK
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(_accessKeyId, _accessKeySecret, "")),
		config.WithRegion(_region),
	)
	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		return nil, err
	}

	client := rds.NewFromConfig(cfg)

	// 获取所有区域
	input := &rds.DescribeSourceRegionsInput{}

	resp, err := client.DescribeSourceRegions(ctx, input)
	if err != nil {
		fmt.Println("Error describing regions:", err)
		return nil, err
	}
	//for _, region := range resp.SourceRegions {
	//	fmt.Println("_region", _region, "Region Name:", *region.RegionName)
	//	fmt.Println("----------------------------------")
	//}

	return resp, err
}
