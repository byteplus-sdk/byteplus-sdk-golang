package cdn

import (
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func AddResourceTags(t *testing.T) {
	resp, err := DefaultInstance.AddResourceTags(&cdn.AddResourceTagsRequest{
		Resources: []string{exampleDomain},
		ResourceTags: []cdn.ResourceTag{
			{Key: cdn.GetStrPtr("userKey"), Value: cdn.GetStrPtr("userValue")},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
