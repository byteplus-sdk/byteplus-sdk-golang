package cdn

import (
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/cdn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func SubmitBlockTask(t *testing.T) {
	resp, err := DefaultInstance.SubmitBlockTask(&cdn.SubmitBlockTaskRequest{
		Urls: exampleDomain,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Result.TaskID)
}
