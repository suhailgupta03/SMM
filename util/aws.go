package util

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"log"
	"time"
)

type AWS struct {
	cfg aws.Config
}

func AWSInit() *AWS {
	// Load the Shared AWS Configuration (~/.aws/config)
	// If the file is not present, will try to read the standard ENV variables
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Failed to read the AWS configuration %v", err)
	}
	return &AWS{
		cfg,
	}
}

func (aws *AWS) DoesLogGroupExist(logGroupName string) bool {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	client := cloudwatchlogs.NewFromConfig(aws.cfg)
	params := &cloudwatchlogs.DescribeLogGroupsInput{
		LogGroupNamePrefix: &logGroupName,
	}
	logGroups, err := client.DescribeLogGroups(ctx, params)
	if err != nil {
		log.Printf("Failed to read the existing log groups %v", err)
		return false
	}

	found := false
	if logGroups != nil {
		for _, lg := range logGroups.LogGroups {
			if *lg.LogGroupName == logGroupName {
				found = true
				break
			}
		}
	}

	return found
}

func (aws *AWS) DoesLogStreamExist(logGroupName, logStreamName string) bool {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	client := cloudwatchlogs.NewFromConfig(aws.cfg)
	params := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        &logGroupName,
		LogStreamNamePrefix: &logStreamName,
	}
	streams, err := client.DescribeLogStreams(ctx, params)
	if err != nil {
		log.Printf("Failed to read the existing log streams for logGroup %s - %v", logGroupName, err)
	}

	found := false
	if streams != nil {
		for _, stream := range streams.LogStreams {
			if *stream.LogStreamName == logStreamName {
				found = true
				break
			}
		}
	}

	return found
}

func (aws *AWS) GetLogs(logGroupName, logStreamName string) ([]*string, error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	var (
		maxLogEventsToFetch int32 = 1000
		startFromHead             = false
		previousWeek              = time.Now().Add(time.Hour * 24 * 7 * -1).UnixMilli()
	)

	client := cloudwatchlogs.NewFromConfig(aws.cfg)
	events, err := client.GetLogEvents(ctx, &cloudwatchlogs.GetLogEventsInput{
		LogStreamName: &logStreamName,
		Limit:         &maxLogEventsToFetch,
		LogGroupName:  &logGroupName,
		StartFromHead: &startFromHead,
		StartTime:     &previousWeek,
	})

	if err != nil {
		log.Printf("Failed to read logs for log-group %s and log-stream %s %v", logGroupName, logStreamName, err)
		return nil, err
	}

	logs := make([]*string, 0)

	if events != nil {
		for _, event := range events.Events {
			logs = append(logs, event.Message)
		}
	}

	return logs, nil
}
