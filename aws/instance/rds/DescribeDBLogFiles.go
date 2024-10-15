package aws_rds

import (
	"context"
	"fmt"
	aws_client "github.com/AIntelligenceGame/cloud/aws/instance/client"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func DescribeDBLogFiles(ctx context.Context, _region string, _accessKeyId string, _accessKeySecret string, _dbInstanceIdentifier string, _slowquerylog string) (*rds.DescribeDBLogFilesOutput, error) {

	client, err := aws_client.ClientRds(ctx, _region, _accessKeyId, _accessKeySecret)
	if err != nil {
		return nil, err
	}
	// 获取慢查询日志文件名
	input := &rds.DescribeDBLogFilesInput{
		DBInstanceIdentifier: aws.String(_dbInstanceIdentifier),
		FilenameContains:     aws.String(_slowquerylog),
	}

	resp, err := client.DescribeDBLogFiles(ctx, input)
	if err != nil {
		fmt.Println("Error describing DB log files:", err)
		return nil, err
	}
	return resp, nil
}
