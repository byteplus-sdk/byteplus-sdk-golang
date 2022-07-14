package cdn

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func DescribeAccountingData(t *testing.T) {
	domain := exampleDomain
	resp, err := DefaultInstance.DescribeAccountingData(&cdn.DescribeAccountingDataRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Domain:    &domain,
		Metric:    "flux",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Result)
}
