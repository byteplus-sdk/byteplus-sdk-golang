package cdn

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func DescribeContentBlockTasks(t *testing.T) {
	resp, err := DefaultInstance.DescribeContentBlockTasks(&cdn.DescribeContentBlockTasksRequest{
		TaskType:  "block_url",
		StartTime: testStartTime,
		EndTime:   testEndTime,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result.Data)
}
