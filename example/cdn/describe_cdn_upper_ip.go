package cdn

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func DescribeCdnUpperIp(t *testing.T) {
	resp, err := DefaultInstance.DescribeCdnUpperIp(&cdn.DescribeCdnUpperIpRequest{Domain: exampleDomain})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
