package cdn

import (
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeEdgeTopStatisticalData(t *testing.T) {
	metric := "flux"
	resp, err := DefaultInstance.DescribeEdgeTopStatisticalData(&cdn.DescribeEdgeTopStatisticalDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Metric:    metric,
		Domain:    exampleDomain,
		Item:      "url",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result.TopDataDetails)
}
