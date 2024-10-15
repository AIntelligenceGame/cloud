package aws_rds

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	aws_client "github.com/AIntelligenceGame/cloud/aws/instance/client"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

func DownloadDBLogFilePortion(ctx context.Context, _region string, _accessKeyId string, _accessKeySecret string, _dbInstanceIdentifier string, logFile types.DescribeDBLogFilesDetails) (error, int64) {
	// 设置 AK 和 SK
	client, err := aws_client.ClientRds(ctx, _region, _accessKeyId, _accessKeySecret)
	if err != nil {
		return err, 0
	}
	// 下载日志文件内容
	downloadInput := &rds.DownloadDBLogFilePortionInput{
		DBInstanceIdentifier: aws.String(_dbInstanceIdentifier),
		LogFileName:          logFile.LogFileName,
		Marker:               nil,
		NumberOfLines:        aws.Int32(1000), // 可根据需要调整获取行数
	}

	downloadResp, err := client.DownloadDBLogFilePortion(context.TODO(), downloadInput)
	if err != nil {
		fmt.Println("Error downloading log file portion:", err)
		return err, 0
	}

	fmt.Println("Log File Content:")
	fmt.Println(*downloadResp.LogFileData)
	// 指定保存路径和文件名
	filepath := "./" + *logFile.LogFileName

	// 创建保存路径（如果不存在）
	os.MkdirAll(filepath[:strings.LastIndex(filepath, "/")], 0755)

	// 将日志文件内容保存到指定路径的文件中
	ioutil.WriteFile(filepath, []byte(*downloadResp.LogFileData), 0644)

	// 打开文件
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err, 0
	}
	defer file.Close() // 确保在函数结束时关闭文件

	// 获取文件信息
	info, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return err, 0
	}

	// 获取文件大小（以字节为单位）
	size := info.Size()

	fmt.Printf("Downloaded log file content saved to %s\n", filepath)
	fmt.Printf("Downloaded log file content size to %d\n", size)

	fmt.Println("----------------------------------")
	return nil, size
}
