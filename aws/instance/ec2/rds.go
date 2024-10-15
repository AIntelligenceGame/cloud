package aws_ec2

import (
	"context"
	"fmt"

	aws_client "github.com/AIntelligenceGame/cloud/aws/instance/client"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func Regions(ctx context.Context, _region string, _accessKeyId string, _accessKeySecret string) (*ec2.DescribeRegionsOutput, error) {
	// 设置 AK 和 SK
	client, err := aws_client.ClientEC2(ctx, _region, _accessKeyId, _accessKeySecret)
	if err != nil {
		return nil, err
	}
	// 获取所有区域
	input := &ec2.DescribeRegionsInput{}

	resp, err := client.DescribeRegions(ctx, input)
	if err != nil {
		fmt.Println("Error describing regions:", err)
		return nil, err
	}

	for _, region := range resp.Regions {
		fmt.Println("_region", _region, "Region Name:", *region.RegionName)
		fmt.Println("----------------------------------")
	}
	return resp, nil
}
