package cdn

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func DescribeIPInfo(t *testing.T) {
	resp, err := DefaultInstance.DescribeIPInfo(&cdn.DescribeIPInfoRequest{
		IP: exampleDomain,
	})
	assert.NoError(t, err)
	assert.Equal(t, exampleDomain, resp.Result.IP)
}
