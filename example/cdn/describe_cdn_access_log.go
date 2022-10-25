package cdn

import (
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeCdnAccessLog(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnAccessLog(&cdn.DescribeCdnAccessLogRequest{
		StartTime: testStartTime,
		EndTime:   testEndTime,
		Domain:    exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	if *resp.Result.TotalCount > 0 {
		assert.NotEmpty(t, resp.Result.DomainLogDetails)
	}
	assert.Greater(t, int(*resp.Result.PageNum), 0)
}
