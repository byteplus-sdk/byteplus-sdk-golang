package cdn

import (
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DescribeIPInfo(t *testing.T) {
	resp, err := DefaultInstance.DescribeIPInfo(&cdn.DescribeIPInfoRequest{
		IP: "1.1.1.1",
	})
	assert.NoError(t, err)
	assert.Equal(t, "1.1.1.1", resp.Result.IP)
}
