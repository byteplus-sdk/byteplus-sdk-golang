package cdn

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TagResources(t *testing.T) {
	resp, err := DefaultInstance.TagResources(&cdn.TagResourcesRequest{
		ResourceType: "Domain",
		ResourceIds:  []string{operateDomain},
		Tags: []cdn.ResourceTag{
			{
				Key:   cdn.GetStrPtr("k1"),
				Value: cdn.GetStrPtr("v1"),
			},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
