package cdn

import (
	"fmt"
	"testing"

	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
)

func DeleteResourceTags(t *testing.T) {
	resp, err := DefaultInstance.DeleteResourceTags(&cdn.DeleteResourceTagsRequest{
		Resources: []string{exampleDomain},
		ResourceTags: []cdn.ResourceTag{
			{Key: cdn.GetStrPtr("userKey"), Value: cdn.GetStrPtr("userValue")},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", resp)
}
