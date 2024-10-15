package aws_rds

import (
	"context"
	"fmt"

	aws_client "github.com/AIntelligenceGame/cloud/aws/instance/client"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

// 获取 rds Instance list

func Instance(ctx context.Context, _region string, _accessKeyId string, _accessKeySecret string) (*rds.DescribeDBInstancesOutput, error) {
	// 设置 AK 和 SK
	client, err := aws_client.ClientRds(ctx, _region, _accessKeyId, _accessKeySecret)
	if err != nil {
		return nil, err
	}
	// 获取实例列表
	instancesInput := &rds.DescribeDBInstancesInput{}

	resp, err := client.DescribeDBInstances(context.TODO(), instancesInput)
	if err != nil {
		fmt.Println("Error describing DB instances:", err)
		return nil, err
	}

	// for _, dbInstance := range respInstances.DBInstances {
	// 	fmt.Println("_region", _region, "DB Instance ID:", *dbInstance.DBInstanceIdentifier)
	// 	fmt.Println("----------------------------------")
	// }
	return resp, nil
}
